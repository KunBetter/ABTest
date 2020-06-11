package experiment

type ExperimentManager interface {
	Init(confList []string)
	GetExperimentGroup(layId int) interface{}
}
