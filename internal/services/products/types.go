package products

import "github.com/bozd4g/comparator/internal/services/configs"

type Service interface {
	GetAll(name string) ([]Dto, error)
	GetAllByCategory(name, category string) ([]Dto, error)
}

type service struct {
	configService configs.Service
}
