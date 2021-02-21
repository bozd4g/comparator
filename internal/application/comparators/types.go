package comparators

type Service interface {
	GetAll(name string)
}

type service struct {}