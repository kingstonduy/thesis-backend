package postgres

import (
	"context"
	"fmt"

	configuration "github.com/kingstonduy/cart-service/internal/bootstrap"
	"github.com/kingstonduy/cart-service/internal/domain"
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

// AddCartItem implements domain.ICartRepo.
func (repo *cartRepoImpl) AddCartItem(ctx context.Context, params domain.AddCartItemParams) error {
	logger.Info(ctx, "AddCartItem start")
	defer logger.Info(ctx, "AddCartItem end")

	sqlQuery := `
        INSERT INTO public."CART_ITEM" (
            "USER_ID", 
            "PRODUCT_ID", 
            "CART_ITEM_QUANTITY", 
            "CREATED_AT", 
            "UPDATED_AT"
        )
        VALUES (
            $1, 
            $2, 
            $3, 
            CURRENT_TIMESTAMP, 
            CURRENT_TIMESTAMP
        );
    `

	return repo.db.DB.WithinTransaction(ctx, func(ctx context.Context) error {
		for _, item := range params.CartItems {
			res, err := repo.db.DB.Exec(ctx, sqlQuery, item.UserID, item.ProductID, item.CartItemQuantity)
			if err != nil {
				logger.Errorf(ctx, err.Error())
				return err
			}

			rowsAffected, err := res.RowsAffected()
			if err != nil {
				return err
			}
			if rowsAffected == 0 {
				err = fmt.Errorf("no rows affected")
				logger.Errorf(ctx, err.Error())
				return err
			}
		}
		return nil
	})
}

// DeleteUserCart implements domain.ICartRepo.
func (repo *cartRepoImpl) DeleteUserCart(ctx context.Context, params domain.DeleteUserCartParams) error {
	logger.Info(ctx, "DeleteUserCart start")
	defer logger.Info(ctx, "DeleteUserCart end")

	sqlQuery := `
        DELETE FROM public."CART_ITEM"
        WHERE "USER_ID" = $1;
    `

	res, err := repo.db.DB.Exec(ctx, sqlQuery, params.UserID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no cart items found for the given user ID")
	}

	return nil
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

// DeleteCartItem implements domain.ICartRepo.
func (repo *cartRepoImpl) DeleteCartItem(ctx context.Context, params domain.DeleteCartItemParams) error {
	logger.Info(ctx, "DeleteCartItem start")
	defer logger.Info(ctx, "DeleteCartItem end")

	sqlQuery := `
        DELETE FROM public."CART_ITEM"
        WHERE "CART_ITEM_ID" = $1;
    `

	res, err := repo.db.DB.Exec(ctx, sqlQuery, params.CartItemID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no cart item found with the given CartItemID")
	}

	return nil
}

// UpdateCartItem implements domain.ICartRepo.
func (repo *cartRepoImpl) UpdateCartItem(ctx context.Context, params domain.UpdateCartItemParams) error {
	logger.Info(ctx, "UpdateCartItem start")
	defer logger.Info(ctx, "UpdateCartItem end")

	sqlQuery := `
        UPDATE public."CART_ITEM"
        SET 
            "CART_ITEM_QUANTITY" = $1,
            "UPDATED_AT" = CURRENT_TIMESTAMP
        WHERE 
            "CART_ITEM_ID" = $2;
    `

	res, err := repo.db.DB.Exec(ctx, sqlQuery, params.CartItemQuantity, params.CartItemID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no cart item found with the given CartItemID")
	}

	return nil
}
