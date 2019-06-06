package structs

import "github.com/jinzhu/gorm"

type Uom struct {
	gorm.Model
	Name string `json:"name"`
}
