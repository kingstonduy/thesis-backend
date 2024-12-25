package postgres

import (
	"context"
	"fmt"

	"github.com/kingstonduy/go-core/logger"
	configuration "github.com/kingstonduy/user-service/internal/bootstrap"
	"github.com/kingstonduy/user-service/internal/domain"
)

type productRepoImlp struct {
	db *configuration.PostgresCon
}

func NewProductRepoImpl(db *configuration.PostgresCon) domain.IUserRepo {
	return &productRepoImlp{
		db: db,
	}
}

// GetUserByUserID implements domain.IUserRepo.
func (repo *productRepoImlp) GetUserByUserID(ctx context.Context, userID string) (domain.UserEntity, error) {
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
		logger.Errorf(ctx, "Error fetching user by ID: %v", err)
		return user, err
	}
	return user, nil
}

// GetUserByUserName implements domain.IUserRepo.
func (repo *productRepoImlp) GetUserByUserName(ctx context.Context, userName string) (domain.UserEntity, error) {
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
		logger.Errorf(ctx, "Error fetching user by username: %v", err)
		return user, err
	}
	return user, nil
}

// Insert implements domain.IUserRepo.
func (repo *productRepoImlp) Insert(ctx context.Context, user domain.UserEntity) error {
	logger.Info(ctx, "Insert start")
	defer logger.Info(ctx, "Insert end")

	sqlQuery := `
        INSERT INTO public."CUSTOMER"
        ("USER_ID", "USER_NAME", "USER_PASSWORD", "USER_IMAGE", 
        "FIRST_NAME", "LAST_NAME", "DATE_OF_BIRTH", "GENDER", 
        "EMAIL", "PHONE_NUMBER", "STREET", "CITY", 
        "CITY_CODE", "DISTRICT", "DISTRICT_CODE", "WARD", 
        "WARD_CODE", "CREATED_AT", "UPDATED_AT")
        VALUES (
            gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, 
            $8, $9, $10, $11, $12, $13, $14, $15, $16, 
            CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
        );
    `
	_, err := repo.db.DB.Exec(ctx, sqlQuery,
		user.UserName.String, user.UserPassword.String, user.UserImage.String,
		user.FirstName.String, user.LastName.String, user.DateOfBirth, user.Gender.String,
		user.Email.String, user.PhoneNumber, user.Street.String, user.City.String,
		user.CityCode.String, user.District.String, user.DistrictCode.String,
		user.Ward.String, user.WardCode.String,
	)
	if err != nil {
		logger.Errorf(ctx, "Error inserting user: %v", err)
		return err
	}
	return nil
}

// Update implements domain.IUserRepo.
func (repo *productRepoImlp) Update(ctx context.Context, user domain.UserEntity) error {
	logger.Info(ctx, "Update start")
	defer logger.Info(ctx, "Update end")

	sqlQuery := `
        UPDATE public."CUSTOMER"
        SET 
            "USER_NAME" = COALESCE($2, "USER_NAME"),
            "USER_PASSWORD" = COALESCE($3, "USER_PASSWORD"),
            "USER_IMAGE" = COALESCE($4, "USER_IMAGE"),
            "FIRST_NAME" = COALESCE($5, "FIRST_NAME"),
            "LAST_NAME" = COALESCE($6, "LAST_NAME"),
            "GENDER" = COALESCE($7, "GENDER"),
            "EMAIL" = COALESCE($8, "EMAIL"),
            "STREET" = COALESCE($9, "STREET"),
            "CITY" = COALESCE($10, "CITY"),
            "CITY_CODE" = COALESCE($11, "CITY_CODE"),
            "DISTRICT" = COALESCE($12, "DISTRICT"),
            "DISTRICT_CODE" = COALESCE($13, "DISTRICT_CODE"),
            "WARD" = COALESCE($14, "WARD"),
            "WARD_CODE" = COALESCE($15, "WARD_CODE"),
            "UPDATED_AT" = CURRENT_TIMESTAMP
        WHERE "USER_ID" = $1;
    `
	res, err := repo.db.DB.Exec(ctx, sqlQuery,
		user.UserID, user.UserName, user.UserPassword, user.UserImage,
		user.FirstName, user.LastName, user.Gender, user.Email,
		user.Street, user.City, user.CityCode, user.District,
		user.DistrictCode, user.Ward, user.WardCode,
	)
	if err != nil {
		logger.Errorf(ctx, "Error updating user: %v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Errorf(ctx, "Error getting rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}
