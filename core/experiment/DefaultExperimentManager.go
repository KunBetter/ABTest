package experiment

import (
	"encoding/json"
	"fmt"
)

type DefaultExperimentManager struct {
	ExperimentGroupMap map[int]ExperimentGroup
}

func (dem *DefaultExperimentManager) GetExperimentGroup(layId int) interface{} {
	return dem.ExperimentGroupMap[layId]
}

func (dem *DefaultExperimentManager) Init(confList []string) {
	for i := 0; i < len(confList); i++ {
		config := confList[i]
		dem.handlerExperimentGroup(config)
	}
}

func (dem *DefaultExperimentManager) handlerExperimentGroup(config string) {
	eg := &ExperimentGroup{}
	err := json.Unmarshal([]byte(config), &eg)
	if err != nil {
		fmt.Println("some error")
	}
	if eg != nil {
		dem.ExperimentGroupMap[eg.LayId] = *eg
	}
}
