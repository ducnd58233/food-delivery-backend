package subscriber

import (
	"context"
	component "food-delivery/components"
)

func Setup(appCtx component.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
}