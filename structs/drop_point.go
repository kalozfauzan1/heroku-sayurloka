package structs

import "github.com/jinzhu/gorm"

type Drop_Point struct {
	gorm.Model
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
	Address  string `json:"address"`
}
