package configs

type Model struct {
	name  string `json:"name" yaml:"name"`
	sites []Site `json:"sites" yaml:"sites"`
}

type Site struct {
	name  string     `json:"name" yaml:"name"`
	steps []Selector `json:"steps" yaml:"steps"`
}

type Selector struct {
	selector string `json:"selector" yaml:"selector"`
	isValue  bool   `json:"isValue" yaml:"isValue"`
}
