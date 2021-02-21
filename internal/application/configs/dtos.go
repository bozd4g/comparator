package configs

type Dto struct {
	Name  string    `yaml:"name"`
	Sites []SiteDto `yaml:"sites"`
}

type SiteDto struct {
	Name  string    `yaml:"name"`
	Steps []StepDto `yaml:"steps"`
}

type StepDto struct {
	Selector string `yaml:"selector"`
	IsValue  bool   `yaml:"isValue"`
}
