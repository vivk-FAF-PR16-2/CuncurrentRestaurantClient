package table

import (
	"errors"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/item"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/tableIdCounter"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"math/rand"
	"sync"
	"time"
)

type Status int

const (
	Wait      Status = 0
	MakeOrder        = 1
)

type Table struct {
	id      int
	status  Status
	mutex   sync.Mutex
	manager *tableIdCounter.TableIdCounter
	menu    *item.Container
	conf    *configuration.Configuration
}

func New(id int) *Table {
	return &Table{
		id:     id,
		status: MakeOrder,
	}
}

func (table *Table) GetId() int {
	return table.id
}

func (table *Table) GetStatus() Status {
	return table.status
}

func (table *Table) GetStatusWait() bool {
	return table.status == Wait
}

func (table *Table) getPriority() int {
	minPriority := table.conf.MinPriority
	maxPriority := table.conf.MaxPriority

	priorityDiff := maxPriority - minPriority

	priority := rand.Intn(priorityDiff) + maxPriority

	return priority
}

func (table *Table) getOrderCount() int {
	minOrderItems := table.conf.MinOrderItems
	maxOrderItems := table.conf.MaxOrderItems

	orderItemsDiff := maxOrderItems - minOrderItems

	orderItemsCount := rand.Intn(orderItemsDiff) + maxOrderItems

	return orderItemsCount
}

func (table *Table) MakeOrder(waiterId int) (*utils.OrderData, error) {
	table.mutex.Lock()
	defer table.mutex.Unlock()

	if table.status != MakeOrder {
		return nil, errors.New("bad table status")
	}

	priority := table.getPriority()
	count := table.getOrderCount()

	items := make([]int, count)
	var maxWait int = 0

	for i := 0; i < count; i++ {
		max := table.menu.GetLen()
		index := rand.Intn(max)
		item, ok := table.menu.Get(index)
		if ok != true {
			return nil, errors.New("bad array index")
		}

		items[i] = item.Id
		maxWait += item.PreparationTime
	}

	finalMaxWait := float32(maxWait) * table.conf.MaxWaitMultiplier
	pickUpTime := time.Now().Unix()

	order := &utils.OrderData{
		OrderID:    table.manager.Get(),
		TableID:    table.id,
		WaiterID:   waiterId,
		Items:      items,
		Priority:   priority,
		MaxWait:    finalMaxWait,
		PickUpTime: pickUpTime,
	}

	table.status = Wait

	return order, nil
}
