package experiment

import (
	"encoding/json"
	"fmt"
)

type DefaultExperimentManager struct {
	ExpGroupMap map[int][]*ExperimentGroup
}

func (manager *DefaultExperimentManager) Init() {
	manager.ExpGroupMap = make(map[int][]*ExperimentGroup)
}

func (manager *DefaultExperimentManager) GetExpGroups(layId int) []*ExperimentGroup {
	return manager.ExpGroupMap[layId]
}

func (manager *DefaultExperimentManager) LoadConfig(configs []string) {
	for i := 0; i < len(configs); i++ {
		config := configs[i]
		manager.handlerExpGroup(config)
	}
}

func (manager *DefaultExperimentManager) handlerExpGroup(config string) {
	var expGroup ExperimentGroup
	err := json.Unmarshal([]byte(config), &expGroup)
	if err != nil {
		fmt.Println("some error")
	}

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

	manager.ExpGroupMap[expGroup.LayId] = []*ExperimentGroup{&expGroup}
}
