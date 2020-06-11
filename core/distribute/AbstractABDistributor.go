package distribute

import (
	"github.com/KunBetter/ABTest/core/context"
	"github.com/KunBetter/ABTest/core/entity"
	"github.com/KunBetter/ABTest/core/experiment"
	"github.com/KunBetter/ABTest/core/strategy"
	"github.com/KunBetter/ABTest/core/util"
)

type AbstractABDistributor struct {
	ABBucketStrategy strategy.ABBucketStrategy
}

func (dis *AbstractABDistributor) Distribute(abTestContext context.ABContext, experimentGroup experiment.ExperimentGroup) entity.ABTag {
	//② whitelist
	whiteKey := experimentGroup.WhiteListKey
	if "" == whiteKey {
		whiteKey = experimentGroup.DivertKey
	}

	divertValue := abTestContext.ContextMap[whiteKey]
	for i := 0; i < len(experimentGroup.Experiments); i++ {
		exp := experimentGroup.Experiments[i]
		if _, ok := exp.WhiteSet[divertValue]; ok {
			return dis.GetExperimentTag(abTestContext, experimentGroup, exp)
		}
	}

	//③0%
	sum := 0
	for i := 0; i < len(experimentGroup.Experiments); i++ {
		exp := experimentGroup.Experiments[i]
		sum += exp.Traffic
	}
	if sum == 0 {
		return dis.GetDefaultTag(abTestContext, experimentGroup)
	}

	//④ 100%
	for i := 0; i < len(experimentGroup.Experiments); i++ {
		exp := experimentGroup.Experiments[i]
		if 100 == exp.Traffic {
			return dis.GetExperimentTag(abTestContext, experimentGroup, exp)
		}
	}

	//⑤ distribute
	bucket := dis.ABBucketStrategy.DoBucket(abTestContext, experimentGroup.LayId, experimentGroup.DivertKey)
	for i := 0; i < len(experimentGroup.Experiments); i++ {
		exp := experimentGroup.Experiments[i]
		if _, ok := exp.Buckets[bucket]; ok {
			return dis.GetExperimentTag(abTestContext, experimentGroup, exp)
		}
	}

	return dis.GetDefaultTag(abTestContext, experimentGroup)
}

func (dis *AbstractABDistributor) GetExperimentTag(abTestContext context.ABContext, experimentGroup experiment.ExperimentGroup, experiment experiment.Experiment) entity.ABTag {
	logTag := experiment.LogTag
	if logTag == "" {
		logTag = experiment.Tag
	}

	return entity.ABTag{
		Tag:      experiment.Tag,
		LogTag:   logTag,
		TraceTag: util.AppendTag(abTestContext.TraceTag, experimentGroup.LayId, logTag),
	}
}

func (dis *AbstractABDistributor) GetDefaultTag(abTestContext context.ABContext, experimentGroup experiment.ExperimentGroup) entity.ABTag {
	return entity.ABTag{
		Tag:      experimentGroup.DefaultTag,
		LogTag:   experimentGroup.DefaultTag,
		TraceTag: util.AppendTag(abTestContext.TraceTag, experimentGroup.LayId, experimentGroup.DefaultTag),
	}
}

func (dis *AbstractABDistributor) GetGlobalTag(abTestContext context.ABContext) entity.ABTag {
	return entity.ABTag{
		Tag:      abTestContext.GlobalTag,
		LogTag:   abTestContext.GlobalTag,
		TraceTag: util.AppendTag(abTestContext.TraceTag, abTestContext.LayId, abTestContext.GlobalTag),
	}
}

func (dis *AbstractABDistributor) IsMeetCondition(conditions map[string]map[string]bool, abTestContext context.ABContext) bool {
	if len(conditions) == 0 {
		return true
	}

	flag := true
	for k, v := range conditions {
		if _, ok := v[abTestContext.ContextMap[k]]; !ok {
			if _, okk := v["*"]; !okk {
				flag = false
				break
			}
		}
	}
	return flag
}
