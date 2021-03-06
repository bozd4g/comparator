package configs

type Service interface {
	GetAll() ([]Dto, error)
	GetByName(name string) (*Dto, error)
}

type service struct {
	configsPath string
}
