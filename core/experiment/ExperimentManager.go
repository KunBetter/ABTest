package experiment

type ExperimentManager interface {
	LoadConfig(confList []string)
	GetExpGroups(layId int) interface{}
}
