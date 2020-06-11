package experiment

type ExperimentManager interface {
	Init(confList []string)
	GetExpGroups(layId int) interface{}
}
