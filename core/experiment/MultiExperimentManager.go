package experiment

import (
	"encoding/json"
	"fmt"
)

type MultiExperimentManager struct {
	DefaultExperimentManager
}

func (manager *MultiExperimentManager) Init() {
	manager.ExpGroupMap = make(map[int][]*ExperimentGroup)
}

func (manager *MultiExperimentManager) GetExpGroups(layId int) []*ExperimentGroup {
	return manager.ExpGroupMap[layId]
}

func (manager *MultiExperimentManager) handlerExpGroup(config string) {
	var expGroups []*ExperimentGroup
	err := json.Unmarshal([]byte(config), &expGroups)
	if err != nil {
		fmt.Println("parse error")
	}
	if nil != expGroups && len(expGroups) > 0 {
		expGroup0 := expGroups[0]
		for i := 0; i < len(expGroups); i++ {
			expGroup := expGroups[i]
			if expGroup0.LayId != expGroup.LayId {
				//new Exception("layids are not consistent")
			}
		}

		manager.ExpGroupMap[expGroup0.LayId] = expGroups
	}
}
