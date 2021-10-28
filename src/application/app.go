package application

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/configuration"
	"github.com/vivk-FAF-PR16-2/RestaurantKitchen/src/http/controller"
	"log"
	"net/http"
)

type IApp interface {
	Start()
	Shutdown()
}

type clientApp struct {
	server *http.Server
}

func New(conf configuration.Configuration) IApp {
	router := gin.New()

	ctrl := controller.NewController(conf)
	ctrl.RegisterRoutes(router)

	return &clientApp{
		server: &http.Server{
			Addr:    conf.DinnerHallAddr,
			Handler: router,
		},
	}
}

func (app *clientApp) Start() {
	if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error while running `food ordering` server: %v\n", err)
	}
}

func (app *clientApp) Shutdown() {
	if err := app.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Unable to shutdown `food ordering` server: %v\n", err)
	}
}
