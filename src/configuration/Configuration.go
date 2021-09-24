package configuration

type Configuration struct {
	DinnerHallAddr string `json:"dinnerHallAddr"`
	KitchenAddr    string `json:"kitchenAddr"`

	DistributionRout string `json:"distributionRout"`
	OrderRout        string `json:"orderRout"`

	MinOrderItems int `json:"min_order_items"`
	MaxOrderItems int `json:"max_order_items"`

	MinPriority int `json:"min_priority"`
	MaxPriority int `json:"max_priority"`

	MaxWaitMultiplier float32 `json:"max_wait_multiplier"`
}
