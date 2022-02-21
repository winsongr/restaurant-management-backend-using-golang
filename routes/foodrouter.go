package routes

import (
	controller  "restaurant/controllers"

	"github.com/gin-gonic/gin"
)
//go mod tidy
func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods",controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id",controller.GetFood())
	incomingRoutes.POST("/foods",controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id",controller.UpdateFood())




}