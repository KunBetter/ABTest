package strategy

import "github.com/KunBetter/ABTest/core/context"

type ABBucketStrategy interface {
	DoBucket(ctx context.ABContext, layId int, divertKey string) int
}
