package main

import (
	"os"
	"restaurant/database"
	"restaurant/routes"
	"restaurant/middleware"
	"go.mongodb.org/mongo-driver/mongo"

)
func main(){
port := os.Getenv("PORT")

if port == "" {
	port ="8000"
	
}
	router:= gin.new()
	router.Use(gin.Logger())
	route

}