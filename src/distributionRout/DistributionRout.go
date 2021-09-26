package distributionRout

import (
	"encoding/json"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/distributionManager"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"io"
	"net/http"
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

	distributionManager.PushQueue(&data)
	http.Error(writer, response, http.StatusOK)
}
