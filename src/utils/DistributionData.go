package utils

type DistributionData struct {
	OrderID        int              `json:"order_id"`
	TableID        int              `json:"table_id"`
	WaiterID       int              `json:"waiter_id"`
	Items          []int            `json:"items"`
	Priority       int              `json:"priority"`
	MaxWait        float32          `json:"max_wait"`
	PickUpTime     int64            `json:"pick_up_time"`
	CookingTime    int              `json:"cooking_time"`
	CookingDetails []CookingDetails `json:"cooking_details"`
}
