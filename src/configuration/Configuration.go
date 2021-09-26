package configuration

type Configuration struct {
	DinnerHallAddr string `json:"dinner_hall_addr"`
	KitchenAddr    string `json:"kitchen_addr"`

	DistributionRout string `json:"distribution_rout"`
	OrderRout        string `json:"order_rout"`

	TimeUnitMod float32 `json:"time_unit_mod"`

	TableCount  int `json:"table_count"`
	WaiterCount int `json:"waiter_count"`

	MinMakeOrder int `json:"min_make_order"`
	MaxMakeOrder int `json:"max_make_order"`

	MinOrderItems int `json:"min_order_items"`
	MaxOrderItems int `json:"max_order_items"`

	MinPriority int `json:"min_priority"`
	MaxPriority int `json:"max_priority"`

	MaxWaitMultiplier float32 `json:"max_wait_multiplier"`
}
