package response

type Categories struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []CategoriesData `json:"data"`
}

type CategoriesData struct {
	Id int `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
