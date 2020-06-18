package experiment

type ExperimentGroup struct {
	LayId        int               `json:"layid"`
	Name         string            `json:"name"`
	DivertKey    string            `json:"divertkey"`
	DefaultTag   string            `json:"defaulttag"`
	WhiteListKey string            `json:"whitelistkey"`
	Conditions   map[string]string `json:"conditions"`
	Experiments  []Experiment      `json:"experiments"`
}
