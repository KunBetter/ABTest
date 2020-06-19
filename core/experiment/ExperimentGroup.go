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

func (expGroup *ExperimentGroup) SetBucket() {
	start := 0
	for i := 0; i < len(expGroup.Experiments); i++ {
		exp := &expGroup.Experiments[i]
		exp.setBuckets(start, start+exp.Traffic)
		start += exp.Traffic

		exp.WhiteSet = make(map[string]bool)
		for i := 0; i < len(exp.Whitelist); i++ {
			exp.WhiteSet[exp.Whitelist[i]] = true
		}
	}

	if start > 100 {
		//throw new Exception
	}
}
