package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/distributionManager"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/domain/dto"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/foodorderingcontroller"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/utils"
	"io"
	"log"
	"net/http"
	"strconv"
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
	r.POST(c.conf.OrderV2Rout, c.foodOrder)
	r.GET(c.conf.OrderIDV2Rout)
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

	foodOrdering := foodorderingcontroller.Get()
	outData = foodOrdering.Add(inData)

	ctx.JSON(http.StatusOK, &outData)
}

func (c *controller) getFoodOrder(ctx *gin.Context) {
	var statusData dto.OrderStatusData

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		message := fmt.Sprintf("error from `%s` route: %v\n", c.conf.OrderIDV2Rout, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		log.Panic(message)
		return
	}

	foodOrdering := foodorderingcontroller.Get()
	statusData, err = foodOrdering.Get(id)
	if err != nil {
		message := fmt.Sprintf("error from `%s` route: %v\n", c.conf.OrderIDV2Rout, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message,
		})
		log.Panic(message)
		return
	}

	ctx.JSON(http.StatusOK, &statusData)
}
