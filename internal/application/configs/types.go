package configs

type Service interface {
	GetAll() ([]Model, error)
	GetByName(name string) (*Model, error)
}

type service struct {
	configsPath string
}
