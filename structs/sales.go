package structs

import "github.com/jinzhu/gorm"

type Sale struct {
	gorm.Model
	Product_Id       int    `json:"product_id"`
	Species_Id       int    `json:"species_id"`
	Uom_Id           int    `json:"uom_id"`
	Initial_Price    int    `json:"initial_price"`
	Packaging_Price  int    `json:"packaging_price"`
	Order_Price      int    `json:"order_price"`
	Remark_Id        int    `json:"remark_id"`
	Denom_Id         int    `json:"denom_id"`
	Weight           string `json:"weight"`
	Stock            int    `json:"stock"`
	Discount         int    `json:"discount"`
	Status_Available int    `json:"status_available"`
	Available_At     string `json:"available_at"`
	Drop_Point_Id    int    `json:"drop_point_id"`
}
