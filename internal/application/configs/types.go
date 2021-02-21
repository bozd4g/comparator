package configs

type Service interface {
	GetAll() []Model
	GetByName(name string) Model
}

type service struct{}
