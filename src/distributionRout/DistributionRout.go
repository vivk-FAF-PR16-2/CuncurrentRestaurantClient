package distributionRout

import (
	"encoding/json"
	"io"
	"net/http"
	"restaurant_client/src/utils"
)

func DistributionHandler(writer http.ResponseWriter, request *http.Request) {
	var data utils.DistributionData
	var response string

	jsonData, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(writer, response, http.StatusOK)
}
