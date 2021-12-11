package products

import (
	"github.com/bozd4g/comparator/internal/services/configs"
)

type Servicer interface {
	GetAll(name string) ([]Dto, error)
	GetAllByConfig(name, config string) ([]Dto, error)
	GetAllMultipleByConfig(names []string, config string) ([]MultipleDto, error)
}

type Service struct {
	configService configs.Servicer
}
