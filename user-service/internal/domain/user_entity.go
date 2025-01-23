package domain

import (
	"context"
	"time"
)

type UserEntity struct {
	UserID       string    `json:"user_id" db:"USER_ID"`
	UserName     string    `json:"user_name" db:"USER_NAME"`
	UserPassword string    `json:"user_password" db:"USER_PASSWORD"`
	UserImage    string    `json:"user_image" db:"USER_IMAGE"`
	FirstName    string    `json:"first_name" db:"FIRST_NAME"`
	LastName     string    `json:"last_name" db:"LAST_NAME"`
	DateOfBirth  string    `json:"date_of_birth" db:"DATE_OF_BIRTH"`
	Gender       string    `json:"gender" db:"GENDER"`
	Email        string    `json:"email" db:"EMAIL"`
	PhoneNumber  string    `json:"phone_number" db:"PHONE_NUMBER"`
	Street       string    `json:"street" db:"STREET"`
	City         string    `json:"city" db:"CITY"`
	CityCode     string    `json:"city_code" db:"CITY_CODE"`
	District     string    `json:"district" db:"DISTRICT"`
	DistrictCode string    `json:"district_code" db:"DISTRICT_CODE"`
	Ward         string    `json:"ward" db:"WARD"`
	WardCode     string    `json:"ward_code" db:"WARD_CODE"`
	CreatedAt    time.Time `json:"created_at" db:"CREATED_AT"`
	UpdatedAt    time.Time `json:"updated_at" db:"UPDATED_AT"`
}

type IUserRepo interface {
	Insert(ctx context.Context, user UserEntity) error
	Update(ctx context.Context, cols map[string]interface{}, conditions map[string]interface{}) error
	GetUserByUserID(ctx context.Context, userID string) (UserEntity, error)
	GetUserByUserName(ctx context.Context, userName string) (UserEntity, error)
}
