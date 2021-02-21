package configs

type Model struct {
	name  string `yaml:"name"`
	sites []Site `yaml:"sites"`
}

type Site struct {
	name  string     `yaml:"name"`
	steps []Selector `yaml:"steps"`
}

type Selector struct {
	selector string `yaml:"selector"`
	isValue  bool   `yaml:"isValue"`
}
