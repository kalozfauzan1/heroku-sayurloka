package structs

import "github.com/jinzhu/gorm"

type Denom struct {
	gorm.Model
	Name string `json:"name"`
}
