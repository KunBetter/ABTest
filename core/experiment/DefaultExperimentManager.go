package experiment

import (
	"encoding/json"
	"fmt"
)

type DefaultExperimentManager struct {
	ExpGroupMap map[int]ExperimentGroup
}

func (manager *DefaultExperimentManager) GetExpGroups(layId int) interface{} {
	return manager.ExpGroupMap[layId]
}

func (manager *DefaultExperimentManager) LoadConfig(configs []string) {
	for i := 0; i < len(configs); i++ {
		config := configs[i]
		manager.handlerExpGroup(config)
	}
}

func (manager *DefaultExperimentManager) handlerExpGroup(config string) {
	expGroup := &ExperimentGroup{}
	err := json.Unmarshal([]byte(config), &expGroup)
	if err != nil {
		fmt.Println("some error")
	}
	if expGroup != nil {
		manager.ExpGroupMap[expGroup.LayId] = *expGroup
	}
}
