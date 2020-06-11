package distribute

import "github.com/KunBetter/ABTest/core/context"
import "github.com/KunBetter/ABTest/core/entity"

type ABDistributor interface {
	Distribute(abTestContext context.ABContext) entity.ABTag
}
