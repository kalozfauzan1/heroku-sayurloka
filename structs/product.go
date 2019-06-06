package structs

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Category_Id           int    `json:"category_id"`
	Category_Name         string `json:"category_name"`
	English_Name          string `json:"english_name"`
	Indonesia_Name        string `json:"indonesia_name"`
	Photo                 string `json:"photo"`
	Description           string `json:"description"`
	Health_Benefits       string `json:"health_benefits"`
	Proucer_Description   string `json:"proucer_description"`
	Photo_Documentation_1 string `json:"photo_documentation_1"`
	Photo_Documentation_2 string `json:"photo_documentation_2"`
}

type ListProduct struct {
	Product_Id             int    `json:"product_id"`
	Product_English_Name   string `json:"product_english_name"`
	Product_Indonesia_Name string `json:"product_indonesia_name"`
	Product_Photo          string `json:"product_photo"`
	Product_Description    string `json:"product_description"`
	Product_Health_Benefit string `json:"product_health_benefit"`
	Category_Id            string `json:"category_id"`
	Category_Name          string `json:"category_name"`
	Species_Id             string `json:"species_id"`
	Species_Name           string `json:"species_name"`
	Uom_Id                 string `json:"uom_id"`
	Uom_Name               string `json:"uom_name"`
	Remark_Id              int    `json:"remark_id"`
	Remark_Name            string `json:"remark_name"`
	Denom_Id               int    `json:"denom_id"`
	Denom_Name             string `json:"denom_name"`
	Weight                 string `json:"weight"`
	Initial_Price          int    `json:"initial_price"`
	Packaging_Price        int    `json:"packging_price"`
	Discount               int    `json:"discount"`
	Order_Price            int    `json:"order_price"`
	Producer_Description   string `json:"producer_description"`
	Photo_Documentation_1  string `json:"photo_documentation_1"`
	Photo_Documentation_2  string `json:"photo_documentation_2"`
	Stock                  int    `json:"stock"`
	Status_Available       int    `json:"status_available"`
	Available_At           string `json:"available_at"`
	Drop_point_Id          int    `json:"drop_point_id"`
}
