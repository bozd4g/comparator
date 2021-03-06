package products

import "github.com/bozd4g/comparator/internal/services/configs"

func New(config configs.Service) Service {
	return service{config: config}
}

func (s service) GetAll(name string) {}
