package configs

type Dto struct {
	Name  string    `yaml:"name" json:"name"`
	Sites []SiteDto `yaml:"sites" json:"sites"`
}

type SiteDto struct {
	Name    string `yaml:"name" json:"name"`
	Address string `yaml:"address" json:"address"`
	Search  string `yaml:"search" json:"search"`
	Link    string `yaml:"link" json:"link"`
	Value   string `yaml:"value" json:"value"`
}
