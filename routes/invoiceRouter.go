package routes

import (
	controller "github.com/atifali-pm/go-resturant-management/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CeateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}
