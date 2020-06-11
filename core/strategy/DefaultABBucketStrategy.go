package strategy

import (
	"github.com/KunBetter/ABTest/core/constant"
	"github.com/KunBetter/ABTest/core/context"
	"github.com/valyala/fastrand"
	"hash/fnv"
	"strconv"
)

type DefaultABBucketStrategy struct {
}

func (bs *DefaultABBucketStrategy) DoBucket(ctx context.ABContext, layId int, divertKey string) int {
	diversionId := getDiversionIdByKey(ctx, divertKey)
	if len(diversionId) > 0 {
		hf := fnv.New32a()
		_, err := hf.Write([]byte(strconv.Itoa(layId) + diversionId))
		if err != nil {
			return -1
		}
		hc := hf.Sum32()
		return int((hc & constant.UnsignedLongNumber) % constant.DefaultBucketAmount)
	}

	return -1
}

func getDiversionIdByKey(ctx context.ABContext, divertKey string) string {
	divertValue := ctx.ContextMap[divertKey]
	if "" == divertValue {
		divertValue = strconv.Itoa(int(fastrand.Uint32n(constant.RandomUpBound)))
		if "random" == divertKey {
			return divertValue
		}
		//LOGGER.warn("can not get divertValue from " + ctx + " by divertKey:" + divertKey);
	}

	return divertValue
}
