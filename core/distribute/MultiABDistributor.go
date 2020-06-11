package distribute

import (
	"github.com/KunBetter/ABTest/core/context"
	"github.com/KunBetter/ABTest/core/entity"
	"github.com/KunBetter/ABTest/core/experiment"
	"github.com/KunBetter/ABTest/core/strategy"
)

type MultiABDistributor struct {
	AbstractABDistributor
	ExperimentManager experiment.ExperimentManager
}

func (dis *MultiABDistributor) Init(manager *experiment.MultiExperimentManager, dbs *strategy.DefaultABBucketStrategy) {
	dis.ExperimentManager = manager
	dis.ExperimentManager.Init(nil)
	dis.AbstractABDistributor.ABBucketStrategy = dbs
}

func (dis *MultiABDistributor) Distribute(abTestContext context.ABContext) entity.ABTag {
	experimentGroup := dis.ExperimentManager.GetExpGroups(abTestContext.LayId)
	if nil == experimentGroup {
		//LOGGER.info("can not find experiment groups by layId:" + abTestContext.getLayId());
		return dis.GetGlobalTag(abTestContext)
	}
	expGroups, ok := experimentGroup.([]experiment.ExperimentGroup)
	if !ok {
		//return errors.New("InitField: require a *Field")
	}
	for i := 0; i < len(expGroups); i++ {
		expGroup := expGroups[i]
		//① conditions
		conditions := expGroup.ConditionSetMap
		if dis.IsMeetCondition(conditions, abTestContext) {
			return dis.AbstractABDistributor.Distribute(abTestContext, expGroup)
		}
	}

	return dis.GetGlobalTag(abTestContext)
}
