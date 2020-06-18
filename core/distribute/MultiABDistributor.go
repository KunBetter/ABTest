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

func (dis *MultiABDistributor) Init(manager *experiment.DefaultExperimentManager, strategy *strategy.DefaultABBucketStrategy) {
	dis.ExperimentManager = manager
	dis.AbstractABDistributor.ABBucketStrategy = strategy
}

func (dis *MultiABDistributor) Distribute(abTestContext context.ABContext) entity.ABTag {
	experimentGroups := dis.ExperimentManager.GetExpGroups(abTestContext.LayId)
	if nil == experimentGroups {
		//LOGGER.info("can not find experiment groups by layId:" + abTestContext.getLayId());
		return dis.GetGlobalTag(abTestContext)
	}
	for i := 0; i < len(experimentGroups); i++ {
		expGroup := experimentGroups[i]
		//① conditions
		conditions := expGroup.Conditions
		if dis.IsMeetCondition(conditions, abTestContext) {
			return dis.AbstractABDistributor.Distribute(abTestContext, expGroup)
		}
	}

	return dis.GetGlobalTag(abTestContext)
}
