package structs

import "github.com/jinzhu/gorm"

type Categorie struct {
	gorm.Model
	Name string `json:"name"`
}

type Categori_Res struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}
