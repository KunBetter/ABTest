package experiment

import (
	"encoding/json"
	"fmt"
)

type MultiExperimentManager struct {
	DefaultExperimentManager
	ExperimentGroupMap map[int][]ExperimentGroup
}

func (dem *MultiExperimentManager) GetExperimentGroup(layId int) interface{} {
	return dem.ExperimentGroupMap[layId]
}

func (mem *MultiExperimentManager) handlerExperimentGroup(config string) {
	var egs []ExperimentGroup
	err := json.Unmarshal([]byte(config), &egs)
	if err != nil {
		fmt.Println("parse error")
	}
	if nil != egs && len(egs) > 0 {
		eg0 := egs[0]
		for i := 0; i < len(egs); i++ {
			eg := egs[i]
			if eg0.LayId != eg.LayId {
				//new Exception("layids are not consistent")
			}
		}

		mem.ExperimentGroupMap[eg0.LayId] = egs
	}
}
