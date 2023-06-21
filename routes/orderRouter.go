package routes

import (
	controller "github.com/atifali-pm/go-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
