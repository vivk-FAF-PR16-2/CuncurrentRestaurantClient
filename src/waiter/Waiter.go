package waiter

import (
	"fmt"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/queue"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/random"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/sendRequest"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/table"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"time"
)

type Waiter struct {
	id     int
	tables []*table.Table
	queue  *queue.Queue

	conf *configuration.Configuration
}

func New(id int, conf *configuration.Configuration) *Waiter {
	return &Waiter{
		id:    id,
		conf:  conf,
		queue: queue.New(),
	}
}

// region Public property

func (waiter *Waiter) GetId() int {
	return waiter.id
}

// endregion

// region Public methods

func (waiter *Waiter) Run() {
	for {
		waiter.update()
	}
}

func (waiter *Waiter) AddTable(table *table.Table) {
	if contains(waiter.tables, table) {
		return
	}

	waiter.tables = append(waiter.tables, table)
}

func (waiter *Waiter) AddDistributionData(data *utils.DistributionData) {
	waiter.queue.Push(data)
}

// endregion

// region Private methods

func (waiter *Waiter) update() {
	orderDone := waiter.queue.Len() != 0
	if orderDone {
		for waiter.queue.Len() != 0 {
			data := waiter.popData()
			tab := waiter.getTableById(data.TableID)
			tab.GetOrder(data)

			fmt.Println("Time of getting order = ", time.Now())
			fmt.Println(*data)
			fmt.Println()
		}
	}

	for _, tab := range waiter.tables {
		if tab.GetStatusMakingOrder() {
			err := tab.StartMakeOrder()
			if err != nil {
				continue
			}

			timeToMakeOrder := random.Range(waiter.conf.MinMakeOrder, waiter.conf.MaxMakeOrder)
			durationToMakeOrder := time.Duration(timeToMakeOrder)
			time.Sleep(configuration.TimeUnit * durationToMakeOrder)

			order, err := tab.FinishMakeOrder(waiter.id)
			if err != nil {
				continue
			}

			sendRequest.SendOrder(order, waiter.conf)

			fmt.Println("Time of sending order = ", time.Now())
			fmt.Println(*order)
			fmt.Println()
		}
	}
}

func (waiter *Waiter) getTableById(id int) *table.Table {
	for _, t := range waiter.tables {
		if t.GetId() == id {
			return t
		}
	}
	return nil
}

func (waiter *Waiter) popData() *utils.DistributionData {
	dataRef := waiter.queue.Pop()
	data := dataRef.(*utils.DistributionData)

	return data
}

// region Utils

func contains(s []*table.Table, e *table.Table) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// endregion

// endregion
