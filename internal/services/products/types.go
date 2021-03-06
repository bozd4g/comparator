package products

import "github.com/bozd4g/comparator/internal/services/configs"

type Service interface {
	GetAll(name string)
}

type service struct {
	config configs.Service
}
