package main

import (
	"encoding/json"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/distributionRout"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/item"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/singleton"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	ConfPath  = "./conf/configuration.json"
	ItemsPath = "./conf/items.json"
)

func main() {
	conf := GetConf()
	container := GetItemContainer()

	singleton.Singleton().Set("items", container)

	http.HandleFunc(conf.DistributionRout, distributionRout.DistributionHandler)

	err := http.ListenAndServe(conf.DinnerHallAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConf() configuration.Configuration {
	var conf configuration.Configuration

	confFile, _ := os.Open(ConfPath)
	defer func(confFile *os.File) {
		_ = confFile.Close()
	}(confFile)

	jsonData, err := io.ReadAll(confFile)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return conf
	}

	err = json.Unmarshal(jsonData, &conf)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return conf
	}

	return conf
}

func GetItemContainer() *item.Container {
	var itemList []item.Item

	itemListFile, _ := os.Open(ItemsPath)
	defer func(itemListFile *os.File) {
		_ = itemListFile.Close()
	}(itemListFile)

	jsonData, err := io.ReadAll(itemListFile)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return nil
	}

	err = json.Unmarshal(jsonData, &itemList)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return nil
	}

	return item.NewContainer(itemList)
}
