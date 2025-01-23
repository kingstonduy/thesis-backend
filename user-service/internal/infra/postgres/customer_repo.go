package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/user-service/internal/bootstrap"
	"github.com/kingstonduy/user-service/internal/domain"
	gensql "github.com/kingstonduy/user-service/internal/pkg/gen_sql"
)

type customerRepoImpl struct {
	db *configuration.PostgresCon
}

func NewCustomerRepo(db *configuration.PostgresCon) domain.IUserRepo {
	return &customerRepoImpl{
		db: db,
	}
}

// GetUserByUserID implements domain.IUserRepo.
func (repo *customerRepoImpl) GetUserByUserID(ctx context.Context, userID string) (domain.UserEntity, error) {
	logger.Info(ctx, "GetUserByUserID start")
	defer logger.Info(ctx, "GetUserByUserID end")

	sqlQuery := `
        SELECT 
        "USER_ID", "USER_NAME", "USER_PASSWORD", "USER_IMAGE", 
        "FIRST_NAME", "LAST_NAME", "GENDER", "EMAIL", 
        "STREET", "CITY", "CITY_CODE", "DISTRICT", 
        "DISTRICT_CODE", "WARD", "WARD_CODE", "CREATED_AT", "UPDATED_AT"
        FROM public."CUSTOMER" WHERE "USER_ID"=$1;
    `
	var user domain.UserEntity
	err := repo.db.DB.Get(ctx, &user, sqlQuery, userID)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUserByUserName implements domain.IUserRepo.
func (repo *customerRepoImpl) GetUserByUserName(ctx context.Context, userName string) (domain.UserEntity, error) {
	logger.Info(ctx, "GetUserByUserName start")
	defer logger.Info(ctx, "GetUserByUserName end")

	sqlQuery := `
        SELECT 
        "USER_ID", "USER_NAME", "USER_PASSWORD", "USER_IMAGE", 
        "FIRST_NAME", "LAST_NAME", "GENDER", "EMAIL", 
        "STREET", "CITY", "CITY_CODE", "DISTRICT", 
        "DISTRICT_CODE", "WARD", "WARD_CODE", "CREATED_AT", "UPDATED_AT"
        FROM public."CUSTOMER" WHERE "USER_NAME"=$1;
    `
	var user domain.UserEntity
	err := repo.db.DB.Get(ctx, &user, sqlQuery, userName)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Insert implements domain.IUserRepo.
func (repo *customerRepoImpl) Insert(ctx context.Context, entity domain.UserEntity) error {
	logger.Info(ctx, "Insert CUSTOMER starts")
	defer logger.Info(ctx, "Insert CUSTOMER ends")

	sqlQuery, err := gensql.GenInsertSql("CUSTOMER", entity)
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

// Update implements domain.IUserRepo.
func (repo *customerRepoImpl) Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error {
	logger.Info(ctx, "Update CUSTOMER starts")
	defer logger.Info(ctx, "Update CUSTOMER ends")

	sqlQuery, err := gensql.GenUpdateSql("CUSTOMER", cols, conditions)
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
