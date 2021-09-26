package distributionManager

import (
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/waiter"
)

var waiters []*waiter.Waiter

func SetWaiters(ws []*waiter.Waiter) {
	waiters = ws
}

func PushQueue(data *utils.DistributionData) {
	for _, w := range waiters {
		if data.WaiterID == w.GetId() {
			w.AddDistributionData(data)
			return
		}
	}
}
