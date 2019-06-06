package structs

import "github.com/jinzhu/gorm"

type Promotion struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Code        string `json:"code"`
	Discount    int    `json:"discount"`
	Status      int    `json:"status"`
}

type Promotion_Res struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Code        string `json:"code"`
	Discount    int    `json:"discount"`
	Status      int    `json:"status"`
}
