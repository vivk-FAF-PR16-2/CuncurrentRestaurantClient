package sendRequest

import (
	"encoding/json"
	"fmt"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
)

const (
	HttpAddr = "http://"
)

func SendOrder(order *utils.OrderData, conf *configuration.Configuration) {
	addr := HttpAddr + conf.KitchenAddr + conf.OrderRout

	jsonBuff, err := json.Marshal(*order)
	if err != nil {
		fmt.Println(err)
		return
	}

	SendRequest(addr, jsonBuff)
}
