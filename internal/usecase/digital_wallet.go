package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/dto"
	"github.com/Hidayathamir/gocheck/internal/entity"
	"github.com/Hidayathamir/gocheck/internal/repo"
	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/gocheckerror"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/Hidayathamir/txmanager"
)

// IDigitalWallet -.
type IDigitalWallet interface {
	Transfer(ctx context.Context, req dto.ReqTransfer) (dto.ResTransfer, error)
}

// DigitalWallet implements IDigitalWallet.
type DigitalWallet struct {
	cfg               *config.Config
	txManager         txmanager.ITransactionManager
	repoDigitalWallet repo.IDigitalWallet
}

var _ IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet -.
func NewDigitalWallet(cfg *config.Config, txManager txmanager.ITransactionManager, repoDigitalWallet repo.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:               cfg,
		txManager:         txManager,
		repoDigitalWallet: repoDigitalWallet,
	}
}

// Transfer implements IDigitalWallet.
func (d *DigitalWallet) Transfer(ctx context.Context, req dto.ReqTransfer) (dto.ResTransfer, error) {
	err := d.validateReqTransfer(ctx, req)
	if err != nil {
		err := fmt.Errorf("%w: %w", gocheckerror.ErrInvalidRequest, err)
		return dto.ResTransfer{}, trace.Wrap(err)
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
		return dto.ResTransfer{}, trace.Wrap(err)
	}

	res := dto.ResTransfer{ID: transactionID}

	return res, nil
}

func (d *DigitalWallet) validateReqTransfer(_ context.Context, req dto.ReqTransfer) error {
	if req.SenderID == 0 {
		err := errors.New("sender id can not be empty")
		return trace.Wrap(err)
	}

	if req.RecipientID == 0 {
		err := errors.New("recipient id can not be empty")
		return trace.Wrap(err)
	}

	if req.SenderID == req.RecipientID {
		err := errors.New("can not transfer to yourself")
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
