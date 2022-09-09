package food_api

import (
	storage "Microservices_and_API/database"
	restaurants "Microservices_and_API/restaurants_api"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getAllFoods(c echo.Context) error {
	db := storage.DB
	var foods []Food

	if result := db.Find(&foods); result.Error != nil {
		return c.String(http.StatusOK, "there are no foods in the database.")
	}

	return c.JSON(http.StatusOK, foods)
}

func getFood(c echo.Context) error {
	db := storage.DB
	food := &Food{}
	food_id := c.Param("food_id")
	if err := db.Where("id = ?", food_id).First(&food).Error; err != nil {
		return c.String(http.StatusBadRequest, "wrong food")
	}

	return c.JSON(http.StatusOK, food)
}

func createFood(c echo.Context) error {
	db := storage.DB
	restaurant := &restaurants.Restaurant{}
	food := &Food{}

	if err := c.Bind(food); err != nil {
		log.Printf("func create Restaurant: Error in binding. Error: %v", err)
		return c.String(http.StatusBadRequest, "bad request body")
	}

	if db.Where("name = ?", food.RestaurantName).First(&restaurant).Error != nil {
		return c.String(http.StatusBadRequest, "there doesn't exist a restaurant with that name")
	}

	return c.JSON(http.StatusOK, food)
}

func updateFood(c echo.Context) error {
	db := storage.DB
	food := &Food{}
	restaurant := &restaurants.Restaurant{}
	food_id := c.Param("food_id")
	if err := db.Where("id = ?", food_id).First(&food).Error; err != nil {
		log.Printf("func updateFood: Error in querying database. Error: %v", err)
		return c.String(http.StatusBadRequest, "there is no food with the given id.")
	}
	if err := c.Bind(food); err != nil {
		log.Printf("func createFood: Error in binding. Error: %v", err)
		return c.String(http.StatusBadRequest, "the given food data are not valid.")
	}

	if db.Where("name = ?", food.RestaurantName).First(&restaurant).Error != nil {
		return c.String(http.StatusBadRequest, "there doesn't exist a restaurant with that name")
	}

	if db.Model(&Food{}).Where("id = ?", food_id).Updates(&food).Error != nil {
		return c.String(http.StatusBadRequest, "bad request, check if there is a food with the same id.")
	}

	return c.JSON(http.StatusOK, food)
}

func deleteFood(c echo.Context) error {
	db := storage.DB
	food := &Food{}
	food_id := c.Param("food_id")
	if err := db.Where("id = ?", food_id).First(&food).Error; err != nil {
		log.Printf("func updateFood: Error in querying database. Error: %v", err)
		c.String(http.StatusBadRequest, "there is no food with the given id.")
	}

	db.Delete(&food_id)
	return c.JSON(http.StatusOK, food_id)
}
