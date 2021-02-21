package comparators

import "github.com/bozd4g/comparator/internal/application/configs"

type Service interface {
	GetAll(name string)
}

type service struct {
	config configs.Service
}
