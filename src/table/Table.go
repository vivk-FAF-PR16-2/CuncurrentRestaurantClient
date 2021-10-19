package table

import (
	"errors"
	"fmt"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/item"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/random"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/ratingSystem"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/tableIdCounter"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"math/rand"
	"sync"
	"time"
)

type Status int

const (
	Free        Status = 0
	MakingOrder        = 1
	Wait               = 2
)

type Table struct {
	id     int
	status Status

	orderStatus chan bool

	mutex   sync.Mutex
	manager *tableIdCounter.TableIdCounter
	menu    *item.Container
	conf    *configuration.Configuration
	rate    *ratingSystem.RatingSystem
}

func New(
	id int,
	manager *tableIdCounter.TableIdCounter,
	menu *item.Container,
	conf *configuration.Configuration) *Table {
	return &Table{
		id:          id,
		manager:     manager,
		menu:        menu,
		conf:        conf,
		orderStatus: make(chan bool),
	}
}

// region Public property

func (table *Table) GetId() int {
	return table.id
}

func (table *Table) GetStatus() Status {
	return table.status
}

func (table *Table) GetStatusMakingOrder() bool {
	return table.status == MakingOrder
}

// endregion

// region Public methods

func (table *Table) Run() {
	for {
		table.update()
	}
}

func (table *Table) StartMakeOrder() error {
	table.mutex.Lock()

	if table.status != MakingOrder {
		return errors.New("bad table status")
	}
	table.status = Wait

	table.mutex.Unlock()

	return nil
}

func (table *Table) FinishMakeOrder(waiterId int) (*utils.OrderData, error) {
	priority := table.getPriority()
	count := table.getOrderCount()

	items := make([]int, count)
	var maxWait int = 0

	for i := 0; i < count; i++ {
		menuLen := table.menu.GetLen()
		index := rand.Intn(menuLen)
		tab, ok := table.menu.Get(index)
		if ok != true {
			return nil, errors.New("bad array index")
		}

		items[i] = tab.Id

		if maxWait < tab.PreparationTime {
			maxWait = tab.PreparationTime
		}
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

	return order, nil
}

func (table *Table) GetOrder(dist *utils.DistributionData) {
	<-table.orderStatus

	now := time.Now().Unix()
	rating := ratingSystem.Calculate(dist.PickUpTime, now, dist.MaxWait)

	fmt.Printf("%s = %d\n", "Raiting for order", rating)

	table.rate.Add(rating)

	fmt.Printf("%s = %f\n", "Raiting for all", table.rate.Return())
}

func (table *Table) SetRatingSystem(rate *ratingSystem.RatingSystem) {
	table.rate = rate
}

// endregion

// region Private methods

func (table *Table) update() {
	table.free()
	table.makingOrder()
}

func (table *Table) free() {
	table.status = Free
	time.Sleep(configuration.TimeUnit)
}

func (table *Table) makingOrder() {
	table.status = MakingOrder
	table.orderStatus <- true
}

func (table *Table) getPriority() int {
	minPriority := table.conf.MinPriority
	maxPriority := table.conf.MaxPriority

	priority := random.Range(minPriority, maxPriority)

	return priority
}

func (table *Table) getOrderCount() int {
	minOrderItems := table.conf.MinOrderItems
	maxOrderItems := table.conf.MaxOrderItems

	orderItemsCount := random.Range(minOrderItems, maxOrderItems)

	return orderItemsCount
}

// endregion
