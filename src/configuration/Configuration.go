package configuration

type Configuration struct {
	DinnerHallAddr string `json:"dinnerHallAddr"`
	KitchenAddr    string `json:"kitchenAddr"`

	DistributionRout string `json:"distributionRout"`
	OrderRout        string `json:"orderRout"`
}
