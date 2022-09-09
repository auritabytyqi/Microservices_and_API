package restaurants

import "github.com/labstack/echo/v4"

func InitializeRoutes(e *echo.Echo) {
	e.POST("/api/restaurants", CreateRestaurant)
	e.GET("/api/restaurants/:restaurant_name", getRestaurantByName)
	e.PATCH("/api/restaurants/:restaurant_name", updateRestaurant)
	e.DELETE("/api/restaurants/:restaurants_name", deleteRestaurant)
}
