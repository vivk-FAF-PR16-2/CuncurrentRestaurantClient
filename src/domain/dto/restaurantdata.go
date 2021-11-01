package dto

import "github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/item"

type RestaurantData struct {
	RestaurantId int         `json:"restaurant_id"`
	Name         string      `json:"name"`
	Address      string      `json:"address"`
	MenuItems    int         `json:"menu_items"`
	Menu         []item.Item `json:"menu"`
	Rating       float32     `json:"rating"`
}
