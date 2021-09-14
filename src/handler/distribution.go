package handler

import (
	"fmt"
	"net/http"
)

func DistributionHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Distribution")
}
