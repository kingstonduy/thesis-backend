package domain

import "context"

type GetUserInformationRequest struct {
	UserID string `json:"userId"`
}
type GetUserInformationResponse struct {
	UserID       string `json:"userId"`
	UserName     string `json:"userName"`
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

type IGetUserInformationHandler interface {
	Handle(ctx context.Context, req *GetUserInformationRequest) (res *GetUserInformationResponse, err error)
}
