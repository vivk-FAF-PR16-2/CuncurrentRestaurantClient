package waiter

import (
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/sendRequest"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/table"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"time"
)

type Waiter struct {
	id     int
	tables []table.Table

	conf *configuration.Configuration
}

func New(id int, conf *configuration.Configuration) *Waiter {
	return &Waiter{
		id:   id,
		conf: conf,
	}
}

func (waiter *Waiter) Run() {
	for {
		waiter.update()
	}
}

func (waiter *Waiter) update() {
	order := waiter.tryMakeOrder()
	if order != nil {
		time.Sleep(2 * time.Second)

		sendRequest.SendOrder(order, waiter.conf)
	}
}

func (waiter *Waiter) tryMakeOrder() *utils.OrderData {
	for _, tab := range waiter.tables {
		if tab.GetStatusWait() == false {
			order, err := tab.MakeOrder(waiter.id)
			if err != nil {
				continue
			}

			return order
		}
	}

	return nil
}
