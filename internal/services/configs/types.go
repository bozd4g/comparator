package configs

type Servicer interface {
	GetAll() ([]Dto, error)
	GetByName(name string) (*Dto, error)
}

type Service struct {
	configsPath string
}
