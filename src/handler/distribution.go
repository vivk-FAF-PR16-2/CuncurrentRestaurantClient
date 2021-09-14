package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restaurant_client/src/utils"
)

func DistributionHandler(writer http.ResponseWriter, request *http.Request) {
	var req utils.Request

	jsonData, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jsonData, &req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = fmt.Fprintf(writer, string(jsonData))
}
