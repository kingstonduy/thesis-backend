package domain

import "context"

type RegisterRequest struct {
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Gender       string `json:"gender"`
	DateOfBirth  string `json:"dateOfBirth"`
	Street       string `json:"street"`
	City         string `json:"city"`
	CityCode     string `json:"cityCode"`
	District     string `json:"district"`
	DistrictCode string `json:"districtCode"`
	Ward         string `json:"ward"`
	WardCode     string `json:"wardCode"`
}
type RegisterResponse struct{}

type IRegisterHandler interface {
	Handle(ctx context.Context, req *RegisterRequest) (res *RegisterResponse, err error)
}
