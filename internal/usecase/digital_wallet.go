package usecase

import (
	"context"
	"fmt"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/dto"
	"github.com/Hidayathamir/gocheck/internal/entity"
	"github.com/Hidayathamir/gocheck/internal/repo"
	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/gocheckerror"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/Hidayathamir/txmanager"
	"github.com/go-playground/validator/v10"
)

// IDigitalWallet defines the interface for the DigitalWallet usecase.
type IDigitalWallet interface {
	// Transfer transfers money from sender to recipient.
	Transfer(ctx context.Context, req dto.ReqDigitalWalletTransfer) (dto.ResDigitalWalletTransfer, error)
}

// DigitalWallet represents the implementation of the DigitalWallet usecase.
type DigitalWallet struct {
	cfg               *config.Config
	validator         *validator.Validate
	txManager         txmanager.ITransactionManager
	repoDigitalWallet repo.IDigitalWallet
}

var _ IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet creates a new instance of the DigitalWallet usecase.
func NewDigitalWallet(cfg *config.Config, txManager txmanager.ITransactionManager, repoDigitalWallet repo.IDigitalWallet) *DigitalWallet {
	validator := validator.New(validator.WithRequiredStructEnabled())
	return &DigitalWallet{
		cfg:               cfg,
		validator:         validator,
		txManager:         txManager,
		repoDigitalWallet: repoDigitalWallet,
	}
}

// Transfer implements IDigitalWallet.
func (d *DigitalWallet) Transfer(ctx context.Context, req dto.ReqDigitalWalletTransfer) (dto.ResDigitalWalletTransfer, error) {
	err := d.validateReqTransfer(ctx, req)
	if err != nil {
		err := fmt.Errorf("%w: %w", gocheckerror.ErrInvalidRequest, err)
		return dto.ResDigitalWalletTransfer{}, trace.Wrap(err)
	}

	var transactionID uint
	err = d.txManager.SQLTransaction(ctx, func(ctx context.Context) error {
		sender, err := d.repoDigitalWallet.GetUserByID(ctx, req.SenderID)
		if err != nil {
			return trace.Wrap(err)
		}

		if req.Amount > sender.Balance {
			err = gocheckerror.ErrInsufficientFunds
			return trace.Wrap(err)
		}

		recipient, err := d.repoDigitalWallet.GetUserByID(ctx, req.RecipientID)
		if err != nil {
			return trace.Wrap(err)
		}

		transactionID, err = d.repoDigitalWallet.CreateTransaction(ctx, entity.Transaction{
			SenderID:    sender.ID,
			RecipientID: recipient.ID,
			Amount:      req.Amount,
		})
		if err != nil {
			return trace.Wrap(err)
		}

		err = d.updateSenderAndRecipientBalance(ctx, sender, recipient, req.Amount)
		if err != nil {
			return trace.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return dto.ResDigitalWalletTransfer{}, trace.Wrap(err)
	}

	res := dto.ResDigitalWalletTransfer{ID: transactionID}

	return res, nil
}

func (d *DigitalWallet) validateReqTransfer(_ context.Context, req dto.ReqDigitalWalletTransfer) error {
	err := d.validator.Struct(req)
	if err != nil {
		return trace.Wrap(err)
	}

	if req.Amount < gocheck.MinimumTransferAmount {
		err := fmt.Errorf("amount can not be less than %d", gocheck.MinimumTransferAmount)
		return trace.Wrap(err)
	}

	return nil
}

func (d *DigitalWallet) updateSenderAndRecipientBalance(ctx context.Context, sender, recipient entity.User, amount int) error {
	err := d.txManager.SQLTransaction(ctx, func(ctx context.Context) error {
		err := d.repoDigitalWallet.UpdateUserBalance(ctx, sender.ID, sender.Balance-amount)
		if err != nil {
			return trace.Wrap(err)
		}

		err = d.repoDigitalWallet.UpdateUserBalance(ctx, recipient.ID, recipient.Balance+amount)
		if err != nil {
			return trace.Wrap(err)
		}

		return nil
	})
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}
