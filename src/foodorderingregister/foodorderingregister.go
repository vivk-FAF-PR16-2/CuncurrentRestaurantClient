package foodorderingregister

import (
	"encoding/json"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/item"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/sendRequest"
	"log"
)

func Register(conf configuration.Configuration, container item.Container) {
	menu := make([]item.Item, container.GetLen())
	for i := 0; i < container.GetLen(); i++ {
		elem, _ := container.Get(i)
		menu[i] = *elem
	}

	info := dto.RestaurantData{
		RestaurantId: conf.RestaurantId,
		Name:         conf.RestaurantName,
		Address:      conf.DinnerHallAddr,
		MenuItems:    len(menu),
		Menu:         menu,
		Rating:       conf.RestaurantRating,
	}

	addr := "http://" + conf.FoodOrderingAddr + "/register"

	jsonBuff, err := json.Marshal(info)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	sendRequest.SendRequest(addr, jsonBuff)
}
