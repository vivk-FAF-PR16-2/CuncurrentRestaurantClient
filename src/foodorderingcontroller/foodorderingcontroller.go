package foodorderingcontroller

import (
	"errors"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/sendRequest"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/tableIdCounter"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"time"
)

type data struct {
	order          utils.OrderData
	createdTime    int64
	cookingTime    int
	cookingDetails []utils.CookingDetails

	status bool
}

type foodOrderingController struct {
	orders []data

	manager *tableIdCounter.TableIdCounter
	conf    *configuration.Configuration

	id int
}

var instance *foodOrderingController

func Get() *foodOrderingController {
	if instance == nil {
		instance = &foodOrderingController{
			orders: make([]data, 0),
			id:     -1,
		}
	}
	return instance
}

func (c *foodOrderingController) Setup(manager *tableIdCounter.TableIdCounter, conf *configuration.Configuration) {
	c.manager = manager
	c.conf = conf
}

func (c *foodOrderingController) Add(inData dto.OrderInData) dto.OrderOutData {
	orderId := c.manager.Get()
	orderData := utils.OrderData{
		OrderID:  orderId,
		TableID:  c.id,
		WaiterID: c.id,
		Items:    inData.Items,
		Priority: inData.Priority,
		MaxWait:  float32(inData.MaxWait),
	}

	c.orders = append(c.orders, data{
		order:       orderData,
		createdTime: inData.CreatedTime,
		status:      false,
	})
	sendRequest.SendOrder(&orderData, c.conf)

	return dto.OrderOutData{
		RestaurantId:         c.conf.RestaurantId,
		OrderId:              orderId,
		EstimatedWaitingTime: inData.MaxWait,
		RegisteredTime:       time.Now().Unix(),
	}
}

func (c *foodOrderingController) SetReady(id int, cookingTime int, cookingDetails []utils.CookingDetails) error {
	index := c.findIndex(id)
	if index == -1 {
		return errors.New("error: wrong ID")
	}

	c.orders[index].status = true
	c.orders[index].cookingTime = cookingTime
	c.orders[index].cookingDetails = cookingDetails
	return nil
}

func (c *foodOrderingController) Get(id int) (dto.OrderStatusData, error) {
	index := c.findIndex(id)
	if index == -1 {
		return dto.OrderStatusData{}, errors.New("error: wrong ID")
	}

	orderData := c.orders[index]
	isReady := orderData.status

	var estimatedWaitingTime int
	var preparedTime int64
	if isReady {
		estimatedWaitingTime = 0
		preparedTime = time.Now().Unix()
	} else {
		estimatedWaitingTime = 1 // TODO: Create method for better time calculation... But how???
		preparedTime = 0
	}

	orderStatus := dto.OrderStatusData{
		OrderId:              orderData.order.OrderID,
		IsReady:              isReady,
		EstimatedWaitingTime: estimatedWaitingTime,
		Priority:             orderData.order.Priority,
		MaxWait:              int(orderData.order.MaxWait),
		CreatedTime:          orderData.createdTime,
		RegisteredTime:       time.Now().Unix(),
		PreparedTime:         preparedTime,
		CookingTime:          orderData.cookingTime,
		CookingDetails:       orderData.cookingDetails,
	}

	return orderStatus, nil
}

func (c *foodOrderingController) findIndex(id int) int {
	resultIndex := -1

	for i := range c.orders {
		if c.orders[i].order.OrderID == id {
			resultIndex = i
			break
		}
	}

	return resultIndex
}
