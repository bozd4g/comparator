package configs

type Model struct {
	Name  string `yaml:"name"`
	Sites []Site `yaml:"sites"`
}

type Site struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Selector string `yaml:"selector"`
	IsValue  bool   `yaml:"isValue"`
}
