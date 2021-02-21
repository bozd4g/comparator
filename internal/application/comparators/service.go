package comparators

import "github.com/bozd4g/comparator/internal/application/configs"

func New(config configs.Service) Service {
	return service{config: config}
}

func (s service) GetAll(name string) {}
