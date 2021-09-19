package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"restaurant_client/src/handler"
	"restaurant_client/src/utils"
	"restaurant_client/src/utils/item"
	"restaurant_client/src/utils/singleton"
)

const (
	ConfPath  = "./conf/configuration.json"
	ItemsPath = "./conf/items.json"
)

func main() {
	conf := GetConf()
	container := GetItemContainer()

	singleton.Singleton().Set("items", container)

	http.HandleFunc(conf.DistributionRout, handler.DistributionHandler)

	err := http.ListenAndServe(conf.DinnerHallAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConf() utils.Configuration {
	var conf utils.Configuration

	confFile, _ := os.Open(ConfPath)
	defer confFile.Close()

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
	defer itemListFile.Close()

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
