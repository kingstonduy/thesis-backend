package domain

import (
	"context"
	"time"
)

const (
	INIT_STATUS     = "INIT"
	COMPLETE_STATUS = "COMPLETE"
	REVERT_STATUS   = "REVERT"
	FAILED_STATUS   = "FAIL"
)

type TransactionEntity struct {
	TransactionID string    `db:"TRANSACTION_ID" json:"transaction_id"`
	Status        string    `db:"STATUS" json:"status"` // INIT COMPLETE REVERT
	Processing    int       `db:"PROCESSING" json:"processing"`
	CreatedAt     time.Time `db:"CREATED_AT" json:"created_at"`
	UpdatedAt     time.Time `db:"UPDATED_AT" json:"updated_at"`
}

func (tr TransactionEntity) GetInitStatus() string {
	return INIT_STATUS
}

func (tr TransactionEntity) GetCompleteStatus() string {
	return COMPLETE_STATUS
}

func (tr TransactionEntity) GetRevertStatus() string {
	return REVERT_STATUS
}

func (tr TransactionEntity) GetFailedStatus() string {
	return FAILED_STATUS
}

type ITransactionRepo interface {
	Insert(ctx context.Context, tr TransactionEntity) error
	Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error
}
