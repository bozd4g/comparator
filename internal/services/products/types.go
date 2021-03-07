package products

import (
	"github.com/bozd4g/comparator/internal/services/configs"
	"github.com/gocolly/colly/v2"
)

type Service interface {
	GetAll(name string) ([]Dto, error)
	GetAllByCategory(name, category string) ([]Dto, error)
}

type service struct {
	collector     *colly.Collector
	configService configs.Service
}
