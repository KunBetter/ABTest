package core

import (
	"github.com/KunBetter/ABTest/core/context"
	"github.com/KunBetter/ABTest/core/distribute"
	"github.com/KunBetter/ABTest/core/experiment"
	"github.com/KunBetter/ABTest/core/strategy"
	"strconv"
)

type ABTest struct {
	manager     *experiment.DefaultExperimentManager
	strategy    *strategy.DefaultABBucketStrategy
	distributer *distribute.MultiABDistributor
}

func (ab *ABTest) Init() {
	ab.manager = &experiment.DefaultExperimentManager{}
	ab.manager.Init()

	ab.strategy = &strategy.DefaultABBucketStrategy{}

	ab.distributer = &distribute.MultiABDistributor{}
	ab.distributer.Init(ab.manager, ab.strategy)
}

func (ab *ABTest) LoadConfig(configs []string) {
	ab.manager.LoadConfig(configs)
}

func (ab *ABTest) Distribute(req map[string]string) map[string]string {
	abContext := &context.ABContext{
		ContextMap: make(map[string]string),
	}
	for k, v := range req {
		abContext.ContextMap[k] = v
	}
	layId, _ := req["layId"]
	id, _ := strconv.Atoi(layId)
	abContext.LayId = id

	abTag := ab.distributer.Distribute(*abContext)

	tagMap := make(map[string]string)
	tagMap["tag"] = abTag.Tag
	tagMap["logTag"] = abTag.LogTag
	tagMap["traceTag"] = abTag.TraceTag

	return tagMap
}
