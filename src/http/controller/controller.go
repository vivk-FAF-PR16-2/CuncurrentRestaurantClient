package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/distributionManager"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"io"
	"log"
	"net/http"
)

type IController interface {
	RegisterRoutes(r *gin.Engine)
}

type controller struct {
	conf configuration.Configuration
}

func NewController(conf configuration.Configuration) IController {
	return &controller{
		conf: conf,
	}
}

func (c *controller) RegisterRoutes(r *gin.Engine) {
	r.POST(c.conf.DistributionRout, c.distribution)
}

func (c *controller) distribution(ctx *gin.Context) {
	var data utils.DistributionData

	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		message := fmt.Sprintf("error from `%s` route: %v\n", c.conf.DistributionRout, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		log.Panic(message)
		return
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		message := fmt.Sprintf("error from `%s` route: %v\n", c.conf.DistributionRout, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		log.Panic(message)
		return
	}

	distributionManager.PushQueue(&data)
	ctx.JSON(http.StatusOK, nil)
}

func (c *controller) foodOrder(ctx *gin.Context) {
	var inData dto.OrderInData
	var outData dto.OrderOutData

	err := json.NewDecoder(ctx.Request.Body).Decode(&inData)
	if err != nil {
		message := fmt.Sprintf("error from `%s` route: %v\n", c.conf.FoodOrderRout, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		log.Panic(message)
		return
	}

	// TODO: `order out data` calculation

	ctx.JSON(http.StatusOK, &outData)
}
