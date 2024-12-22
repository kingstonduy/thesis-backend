package configuration

import (
	"github.com/kingstonduy/go-core/mapping"
	"github.com/kingstonduy/go-core/mapping/mapstructure"
)

func GetMapper() mapping.Mapper {
	return mapstructure.NewMapStructure()
}
