package domain

import "context"

type CityEntity struct {
	Id   string `json:"Id" db:"CITY_CODE"`
	Name string `json:"Name" db:"CITY_NAME"`
}
type DistrictEntity struct {
	Id   string `json:"Id" db:"DISTRICT_CODE"`
	Name string `json:"Name" db:"DISTRICT_NAME"`
}
type WardEntity struct {
	Id   string `json:"Id" db:"WARD_CODE"`
	Name string `json:"Name" db:"WARD_NAME"`
}

type ILocationRepo interface {
	GetCity(ctx context.Context, cityCode string) (string, error)
	GetDistrict(ctx context.Context, districtCode string) (string, error)
	GetWard(ctx context.Context, wardCode string) (string, error)
}
