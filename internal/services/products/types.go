package products

import (
	"github.com/bozd4g/comparator/internal/services/configs"
	"github.com/gocolly/colly/v2"
)

type Service interface {
	GetAll(name string) ([]Dto, error)
	GetAllByCategory(name, category string) ([]Dto, error)
	GetAllMultipleByCategory(names []string, category string) ([]MultipleDto, error)
}

type service struct {
	collector     *colly.Collector
	configService configs.Service
}

const SAMPLE_USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/537.36 Safari/537.36"