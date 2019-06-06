package structs

import "github.com/jinzhu/gorm"

type Specie struct {
	gorm.Model
	Name string `json:"name"`
}
