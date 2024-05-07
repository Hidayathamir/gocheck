package repologgermw

import (
	"context"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/entity"
	"github.com/Hidayathamir/gocheck/internal/repo"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
)

// DigitalWallet represents the implementation of the DigitalWallet logger middleware.
type DigitalWallet struct {
	cfg  *config.Config
	next repo.IDigitalWallet
}

var _ repo.IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet creates a new instane of DigitalWallet logger middleware.
func NewDigitalWallet(cfg *config.Config, next repo.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:  cfg,
		next: next,
	}
}

// GetUserByID implements repo.IDigitalWallet.
func (d *DigitalWallet) GetUserByID(ctx context.Context, id uint) (entity.User, error) {
	user, err := d.next.GetUserByID(ctx, id)

	level := logrus.InfoLevel
	if err != nil {
		level = logrus.ErrorLevel
	}

	logrus.WithFields(logrus.Fields{
		"funcName": trace.FuncName(),
		"in": logrus.Fields{
			"id": id,
		},
		"out": logrus.Fields{
			"user": user,
			"err":  err,
		},
	}).Log(level, level)

	return user, err
}

// CreateTransaction implements repo.IDigitalWallet.
func (d *DigitalWallet) CreateTransaction(ctx context.Context, transaction entity.Transaction) (uint, error) {
	transactionID, err := d.next.CreateTransaction(ctx, transaction)

	level := logrus.InfoLevel
	if err != nil {
		level = logrus.ErrorLevel
	}

	logrus.WithFields(logrus.Fields{
		"funcName": trace.FuncName(),
		"in": logrus.Fields{
			"transaction": transaction,
		},
		"out": logrus.Fields{
			"transactionID": transactionID,
			"err":           err,
		},
	}).Log(level, level)

	return transactionID, err
}

// UpdateUserBalance implements repo.IDigitalWallet.
func (d *DigitalWallet) UpdateUserBalance(ctx context.Context, userID uint, balance int) error {
	err := d.next.UpdateUserBalance(ctx, userID, balance)

	level := logrus.InfoLevel
	if err != nil {
		level = logrus.ErrorLevel
	}

	logrus.WithFields(logrus.Fields{
		"funcName": trace.FuncName(),
		"in": logrus.Fields{
			"userID":  userID,
			"balance": balance,
		},
		"out": logrus.Fields{
			"err": err,
		},
	}).Log(level, level)

	return err
}
