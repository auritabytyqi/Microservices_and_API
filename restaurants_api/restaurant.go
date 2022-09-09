package restaurants

type Restaurant struct {
	Name        string `json:"name" gorm:"primaryKey"`
	Description string `json:"description"`
}
