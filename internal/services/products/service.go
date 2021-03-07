package products

import (
	"github.com/bozd4g/comparator/internal/services/configs"
)

func New(config configs.Service) Service {
	return service{configService: config}
}

func (s service) GetAll(name string) ([]Dto, error) {
	_, err := s.configService.GetAll()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s service) GetAllByCategory(name, category string) ([]Dto, error) {
	_, err := s.configService.GetByName(category)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
