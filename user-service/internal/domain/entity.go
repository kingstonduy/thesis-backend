package domain

import (
	"context"
	"database/sql"
)

type UserEntity struct {
	UserID       sql.NullString `json:"user_id" db:"USER_ID"`
	UserName     sql.NullString `json:"user_name" db:"USER_NAME"`
	UserPassword sql.NullString `json:"user_password" db:"USER_PASSWORD"`
	UserImage    sql.NullString `json:"user_image" db:"USER_IMAGE"`
	FirstName    sql.NullString `json:"first_name" db:"FIRST_NAME"`
	LastName     sql.NullString `json:"last_name" db:"LAST_NAME"`
	DateOfBirth  sql.NullString `json:"date_of_birth" db:"DATE_OF_BIRTH"`
	Gender       sql.NullString `json:"gender" db:"GENDER"`
	Email        sql.NullString `json:"email" db:"EMAIL"`
	PhoneNumber  sql.NullString `json:"phone_number" db:"PHONE_NUMBER"`
	Street       sql.NullString `json:"street" db:"STREET"`
	City         sql.NullString `json:"city" db:"CITY"`
	CityCode     sql.NullString `json:"city_code" db:"CITY_CODE"`
	District     sql.NullString `json:"district" db:"DISTRICT"`
	DistrictCode sql.NullString `json:"district_code" db:"DISTRICT_CODE"`
	Ward         sql.NullString `json:"ward" db:"WARD"`
	WardCode     sql.NullString `json:"ward_code" db:"WARD_CODE"`
	CreatedAt    sql.NullTime   `json:"created_at" db:"CREATED_AT"`
	UpdatedAt    sql.NullTime   `json:"updated_at" db:"UPDATED_AT"`
}

type IUserRepo interface {
	Insert(ctx context.Context, user UserEntity) error
	Update(ctx context.Context, user UserEntity) error
	GetUserByUserID(ctx context.Context, userID string) (UserEntity, error)
	GetUserByUserName(ctx context.Context, userName string) (UserEntity, error)
}
