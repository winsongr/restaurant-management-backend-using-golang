package routes


import (
	controller "restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menus_id", controller.GetMenus())
	incomingRoutes.POST("/menus", controller.CreateMenus())
	incomingRoutes.PATCH("/menus/:menus_id", controller.UpdateMenus())

}
