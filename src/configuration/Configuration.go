package configuration

type Configuration struct {
	DinnerHallAddr string `json:"dinner_hall_addr"`
	KitchenAddr    string `json:"kitchen_addr"`

	DistributionRout string `json:"distribution_rout"`
	OrderRout        string `json:"order_rout"`
	OrderV2Rout      string `json:"order_v2_rout"`

	FoodOrderRout   string `json:"food_order_rout"`
	FoodOrderIdRout string `json:"food_order_id_rout"`

	RestaurantId     int     `json:"restaurant_id"`
	RestaurantName   string  `json:"restaurant_name"`
	RestaurantRating float32 `json:"restaurant_rating"`
	FoodOrderingAddr string  `json:"food_ordering_addr"`

	TableCount  int `json:"table_count"`
	WaiterCount int `json:"waiter_count"`

	MinMakeOrder int `json:"min_make_order"`
	MaxMakeOrder int `json:"max_make_order"`

	MinOrderItems int `json:"min_order_items"`
	MaxOrderItems int `json:"max_order_items"`

	MinPriority int `json:"min_priority"`
	MaxPriority int `json:"max_priority"`

	MaxWaitMultiplier float32 `json:"max_wait_multiplier"`

	TimeUnitMillisecondMultiplier int `json:"time_unit_millisecond_multiplier"`
}
