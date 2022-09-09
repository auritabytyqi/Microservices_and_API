package main

import (
	foods "Microservices_and_API/food_api"
	restaurants "Microservices_and_API/restaurants_api"

	storage "Microservices_and_API/database"

	"github.com/labstack/echo/v4"
)

func main() {
	db := storage.DB
	db.AutoMigrate(&restaurants.Restaurant{})
	db.AutoMigrate(&foods.Food{})

	e := echo.New()
	restaurants.InitializeRoutes(e)
	foods.InitializeRoutes(e)

	e.Start(":8080")
}
