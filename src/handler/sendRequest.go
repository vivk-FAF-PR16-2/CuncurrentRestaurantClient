package handler

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const (
	method = "POST"
)

func SendRequest(addr string, body []byte) {
	request, err := http.NewRequest(method, addr, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("exit: %s\n", err.Error())
		return
	}
	defer response.Body.Close()

	fmt.Println(response.Status)

}
