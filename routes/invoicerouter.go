package routes


import (
	controller "restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoices_id", controller.GetInvoices())
	incomingRoutes.POST("/invoices", controller.CreateInvoices())
	incomingRoutes.PATCH("/invoices/:invoices_id", controller.UpdateInvoices())

}
