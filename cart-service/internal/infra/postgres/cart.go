package postgres

import (
	"context"
	"fmt"

	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
	gensql "github.com/kingstonduy/cart-service/internal/pkg/gen_sql"
	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
)

type cartRepoImpl struct {
	db *configuration.PostgresCon
}

func NewCartRepo(db *configuration.PostgresCon) domain.ICartRepo {
	return &cartRepoImpl{
		db: db,
	}
}

// GetCart implements domain.ICartRepo.
func (repo *cartRepoImpl) GetCart(ctx context.Context, paramsIn domain.GetCartParamsIn) (paramsOut domain.GetCartParamsOut, err error) {
	logger.Info(ctx, "GetCartItemsByUserID start")
	defer logger.Info(ctx, "GetCartItemsByUserID end")

	sqlQuery := `
        SELECT 
            ci."CART_ITEM_ID" AS cartItemID,
            ci."PRODUCT_ID" AS productID,
            p."PRODUCT_NAME" AS productName,
            p."PRODUCT_IMAGE" AS productImage,
            p."PRODUCT_CATEGORY" AS productCategory,
            ci."CART_ITEM_QUANTITY" AS cartItemQuantity,
            p."PRODUCT_PRICE" AS productPrice
        FROM 
            public."CART_ITEM" ci
        INNER JOIN 
            public."PRODUCT" p
        ON 
            ci."PRODUCT_ID" = p."PRODUCT_ID"
        WHERE 
            ci."USER_ID" = $1;
    `

	rows, err := repo.db.DB.Query(ctx, sqlQuery, paramsIn.UserID)
	if err != nil {
		return paramsOut, err
	}
	defer rows.Close()

	var cartItems []domain.GetCartItemDetail
	for rows.Next() {
		var item domain.GetCartItemDetail
		err := rows.Scan(
			&item.CartItemID,
			&item.ProductID,
			&item.ProductName,
			&item.ProductImage,
			&item.ProductCatergory,
			&item.CartItemQuantity,
			&item.ProductPrice,
		)
		if err != nil {
			return paramsOut, err
		}
		cartItems = append(cartItems, item)
	}

	if err = rows.Err(); err != nil {
		return paramsOut, err
	}

	paramsOut.CartItems = cartItems

	return paramsOut, nil
}

// DeleteCartItemsByID implements domain.ICartRepo.
func (repo *cartRepoImpl) DeleteById(ctx context.Context, id string) error {
	logger.Info(ctx, "DeleteCartItem start")
	defer logger.Info(ctx, "DeleteCartItem end")

	sqlQuery := `
	        DELETE FROM public."CART_ITEM"
	        WHERE "CART_ITEM_ID" = $1;
	    `

	res, err := repo.db.DB.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf(errorx.ErrorMessageNoRowAffected)
	}

	return nil
}

// Insert implements domain.ICartRepo.
func (repo *cartRepoImpl) Insert(ctx context.Context, entity domain.CartItem) error {
	logger.Info(ctx, "Insert CART_ITEM starts")
	defer logger.Info(ctx, "Insert CART_ITEM ends")

	sqlQuery, err := gensql.GenInsertSql("CART_ITEM", entity)
	if err != nil {
		return err
	}

	logger.Info(ctx, sqlQuery)

	res, err := repo.db.DB.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return fmt.Errorf(errorx.ErrorMessageNoRowAffected)
	}

	return nil
}

// Update implements domain.ICartRepo.
func (repo *cartRepoImpl) Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error {
	logger.Info(ctx, "Update CART_ITEM starts")
	defer logger.Info(ctx, "Update CART_ITEM ends")

	sqlQuery, err := gensql.GenUpdateSql("CART_ITEM", cols, conditions)
	if err != nil {
		return err
	}

	logger.Info(ctx, sqlQuery)

	res, err := repo.db.DB.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return fmt.Errorf(errorx.ErrorMessageNoRowAffected)
	}

	return nil
}
