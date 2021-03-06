package configs

type Dto struct {
	Name  string    `yaml:"name" json:"name"`
	Sites []SiteDto `yaml:"sites" json:"sites"`
}

type SiteDto struct {
	Name  string    `yaml:"name" json:"name"`
	Steps []StepDto `yaml:"steps" json:"steps"`
}

type StepDto struct {
	Selector string `yaml:"selector" json:"selector"`
	IsValue  bool   `yaml:"isValue" json:"isValue"`
}
