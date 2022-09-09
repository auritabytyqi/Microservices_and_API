package food_api

import restaurants "Microservices_and_API/restaurants_api"

type Food struct {
	Id             string                 `json:"id"`
	Name           string                 `json:"name"`
	Descriptiom    int                    `json:"description"`
	RestaurantName string                 `json:""restaurant_name`
	Restaurant     restaurants.Restaurant `gorm:"foreignKey:RestaurantName;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
