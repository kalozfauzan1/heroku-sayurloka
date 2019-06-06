package structs

import "github.com/jinzhu/gorm"

type Remarks struct {
	gorm.Model
	Name string `json:"name"`
}
