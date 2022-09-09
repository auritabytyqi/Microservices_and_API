package restaurants

import (
	"net/http"

	storage "Microservices_and_API/database"
	"log"

	"github.com/labstack/echo/v4"
)

func CreateRestaurant(c echo.Context) error {
	db := storage.DB
	restaurant := &Restaurant{}

	if err := c.Bind(restaurant); err != nil {
		log.Printf("func createRestaurant: Error in binding. Error:%v ", err)
		return c.String(http.StatusBadRequest, "bad request: the given restaurant is not of the desired format")
	}

	if restaurant.Name == "" {
		return c.String(http.StatusBadRequest, "bad request:restaurant name can't be empty string")
	}
	if db.Create(&restaurant).Error != nil {
		return c.String(http.StatusBadRequest, "there was a problem with creating the restaurant, make sure that there is a restaurant name and it isn't duplicated")
	}

	return c.JSON(http.StatusOK, restaurant)
}

func getRestaurantByName(c echo.Context) error {
	db := storage.DB
	restaurant := &Restaurant{}
	restaurant_name := c.Param("restaurant_name")
	if err := db.Where("name = ?", restaurant_name).First(&restaurant).Error; err != nil {
		log.Printf("func getRestaurantByName: Error in querying database. Error: %v", err)
		return c.String(http.StatusBadRequest, "bad request: there are no restaurants with that name.")
	}

	return c.JSON(http.StatusOK, restaurant)
}

func updateRestaurant(c echo.Context) error {
	db := storage.DB
	restaurant := &Restaurant{}
	restaurant_name := c.Param("restaurant_name")
	if err := db.Where("name = ?", restaurant_name).First(&restaurant).Error; err != nil {
		log.Printf("func updateRestaurant: Error in querying database. Error: %v", err)
		return c.String(http.StatusBadRequest, "bad request: there are no restaurants with that name")
	}
	if err := c.Bind(restaurant); err != nil {
		log.Printf("func updateRestaurant: Error in binding. Error:%v ", err)
		return c.String(http.StatusBadRequest, "the given restaurant is not of the desired format")
	}
	if restaurant.Name == "" {
		return c.String(http.StatusBadRequest, "bad request: restaurant name can't be empty string")
	}
	if db.Model(&Restaurant{}).Where("name = ?", restaurant_name).Updates(&restaurant).Error != nil {
		return c.String(http.StatusOK, "bad request: you can't update to this food, check if this restaurant_name already taken.")
	}

	return c.JSON(http.StatusOK, restaurant)
}

func deleteRestaurant(c echo.Context) error {
	db := storage.DB
	restaurant := &Restaurant{}
	restaurant_name := c.Param("restaurant_name")
	if err := db.Where("name = ?", restaurant_name).First(&restaurant).Error; err != nil {
		log.Printf("func deleteRestaurant: Error in querying database. Error: %v", err)
		return c.String(http.StatusBadRequest, "bad request: there are no restaurants with that name.")
	}
	db.Delete(&restaurant)
	return c.JSON(http.StatusOK, restaurant)
}
