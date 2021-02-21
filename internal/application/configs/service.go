package configs

func New() Service {
	return service{}
}

func (s service) GetAll() []Model {
	return []Model{}
}

func (s service) GetByName(name string) Model {
	return Model{}
}
