package experiment

type ExperimentGroup struct {
	LayId           int                        `json:"layid"`
	Name            string                     `json:"name"`
	DivertKey       string                     `json:"divertkey"`
	DefaultTag      string                     `json:"defaulttag"`
	WhiteListKey    string                     `json:"whitelistkey"`
	Conditions      map[string][]interface{}   `json:"conditions"`
	ConditionSetMap map[string]map[string]bool `json:"conditionsetmap"`
	Experiments     []Experiment               `json:"experiments"`
}

func (expGroup *ExperimentGroup) setExperiments(experiments []Experiment) {
	start := 0
	for i := 0; i < len(experiments); i++ {
		exp := experiments[i]
		exp.setBuckets(start, start+exp.Traffic)
		start += exp.Traffic
	}

	if start > 100 {
		//throw new Exception
	}
}
