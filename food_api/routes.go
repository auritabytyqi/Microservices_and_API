package food_api

import "github.com/labstack/echo/v4"

func InitializeRoutes(e *echo.Echo) {
	e.GET("/api/foods", getAllFoods)
	e.GET("/api/foods/:food_id", getFood)
	e.POST("/api/foods", createFood)
	e.PATCH("/api/foods/:food_id", updateFood)
	e.DELETE("/api/foods/:food_id", deleteFood)
}
