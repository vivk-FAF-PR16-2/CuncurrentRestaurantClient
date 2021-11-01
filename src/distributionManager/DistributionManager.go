package distributionManager

import (
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/foodorderingcontroller"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/waiter"
	"log"
)

var waiters []*waiter.Waiter

func SetWaiters(ws []*waiter.Waiter) {
	waiters = ws
}

func PushQueue(data *utils.DistributionData) {
	if data.WaiterID == -1 {
		err := foodorderingcontroller.Get().SetReady(data.OrderID, data.CookingTime, data.CookingDetails)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		return
	}

	for _, w := range waiters {
		if data.WaiterID == w.GetId() {
			w.AddDistributionData(data)
			return
		}
	}
}
