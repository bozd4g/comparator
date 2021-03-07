package configs

type Dto struct {
	Name  string    `yaml:"name" json:"name"`
	Sites []SiteDto `yaml:"sites" json:"sites"`
}

type SiteDto struct {
	Name    string    `yaml:"name" json:"name"`
	Address string    `yaml:"address" json:"address"`
	Steps   []StepDto `yaml:"steps" json:"steps"`
}

type StepDto struct {
	Id       int    `yaml:"id" json:"id"`
	Selector string `yaml:"selector" json:"selector"`
	Action   Action `yaml:"action" json:"action"`
}

type Action string

const (
	SEARCH Action = "SEARCH"
	CLICK  Action = "CLICK"
	LINK  Action = "LINK"
	VALUE  Action = "VALUE"
)

func (a Action) String() string {
	switch a {
	case SEARCH:
		return "SEARCH"
	case CLICK:
		return "CLICK"
	case LINK:
		return "LINK"
	case VALUE:
		return "VALUE"
	default:
		return ""
	}
}
