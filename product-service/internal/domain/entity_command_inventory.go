package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type WriteInventoryInventory struct {
	ID                uuid.UUID `json:"id" db:"ID"`                                 // Primary key
	ProductID         uuid.UUID `json:"product_id" db:"PRODUCT_ID"`                 // Foreign key referencing PRODUCT table
	InventoryQuantity int       `json:"inventory_quantity" db:"INVENTORY_QUANTITY"` // Nullable integer for inventory quantity
	CreatedAt         time.Time `json:"created_at" db:"CREATED_AT"`                 // Timestamp for record creation
	UpdatedAt         time.Time `json:"updated_at" db:"UPDATED_AT"`                 // Timestamp for last record update
}

type IWriteInventoryRepo interface {
	Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error
	Insert(ctx context.Context, entity interface{}) error
	SelectByProductID(ctx context.Context, id string) (WriteInventoryInventory, error)
}
