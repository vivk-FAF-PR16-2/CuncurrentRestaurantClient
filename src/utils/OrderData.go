package utils

import "fmt"

type OrderData struct {
	OrderID    int     `json:"order-id"`
	TableID    int     `json:"table-id"`
	WaiterID   int     `json:"waiter-id"`
	Items      []int   `json:"items"`
	Priority   int     `json:"priority"`
	MaxWait    float32 `json:"max-wait"`
	PickUpTime int64   `json:"pick-up-time"`
}

func (data OrderData) String() string {
	var result string

	result += fmt.Sprintln("OrderID = ", data.OrderID)
	result += fmt.Sprintln("TableID = ", data.TableID)
	result += fmt.Sprintln("WaiterID = ", data.WaiterID)
	result += fmt.Sprintln("Items = ", data.Items)
	result += fmt.Sprintln("Priority = ", data.Priority)
	result += fmt.Sprintln("MaxWait = ", data.MaxWait)
	result += fmt.Sprintln("PickUpTime = ", data.PickUpTime)

	return result
}
