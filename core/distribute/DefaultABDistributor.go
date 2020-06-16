package distribute

import (
	"github.com/KunBetter/ABTest/core/context"
	"github.com/KunBetter/ABTest/core/entity"
	"github.com/KunBetter/ABTest/core/experiment"
	"github.com/KunBetter/ABTest/core/strategy"
)

type DefaultABDistributor struct {
	AbstractABDistributor
	ExperimentManager experiment.ExperimentManager
}

func (dis *DefaultABDistributor) Init(manager *experiment.DefaultExperimentManager, strategy *strategy.DefaultABBucketStrategy) {
	dis.ExperimentManager = manager
	dis.AbstractABDistributor.ABBucketStrategy = strategy
}

func (dis *DefaultABDistributor) Distribute(abTestContext context.ABContext) entity.ABTag {
	experimentGroup := dis.ExperimentManager.GetExpGroups(abTestContext.LayId)
	if nil == experimentGroup {
		//LOGGER.info("can not find experiment group by layId:" + abTestContext.getLayId());
		return dis.GetGlobalTag(abTestContext)
	}
	//① conditions
	eg, ok := experimentGroup.(experiment.ExperimentGroup)
	if !ok {
		//return errors.New("InitField: require a *Field")
	}
	conditions := eg.ConditionSetMap
	if dis.IsMeetCondition(conditions, abTestContext) {
		return dis.AbstractABDistributor.Distribute(abTestContext, eg)
	}

	return dis.GetGlobalTag(abTestContext)
}
