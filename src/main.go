package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"restaurant_client/src/handler"
	"restaurant_client/src/utils"
)

const (
	ConfPath = "./conf/configuration.json"
)

func main() {
	conf := GetConf()

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
