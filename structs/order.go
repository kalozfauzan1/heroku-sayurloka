package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Order_Code    string    `json:"order_code"`
	Drop_Point_Id int       `json:"drop_point_id"`
	Customer_Id   string    `json:"customer_id"`
	Discount      int       `json:"discount"`
	Delivery_Fee  int       `json:"delivery_fee"`
	Total_Price   int       `json:"total_price"`
	Order_Date    time.Time `json:"order_date"`
	Expired_Date  string    `json:"expired_date"`
	Delivery_Date string    `json:"delivery_date"`
	Delivery_Time string    `json:"delivery_time"`
	Note          string    `json:"note"`
	Status        int       `json:"status"`
}

type Order_Detail_Receiver struct {
	gorm.Model
	Order_Code   string `json:"order_code"`
	Name         string `json:"name"`
	Phone_Number string `json:"phone_number"`
	Address      string `json:"address"`
	Note_Address string `json:"note_address"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

type Order_Detail_Product struct {
	gorm.Model
	Order_Code      string `json:"order_code"`
	Name            string `json:"name"`
	Category        string `json:"category"`
	Species         string `json:"species"`
	Uom             string `json:"uom"`
	Initial_Price   int    `json:"initial_price"`
	Packaging_Price int    `json:"packaging_price"`
	Order_Price     int    `json:"order_price"`
	Qty             int    `json:"qty"`
	Subtotal        int    `json:"subtotal"`
	Discount        int    `json:"discount"`
	Total           int    `json:"total"`
}

type Order_Detail_Payment struct {
	gorm.Model
	Order_Code         string    `json:"order_code"`
	Payment_Method     int       `json:"payment_method"`
	Transaction_Time   time.Time `json:"transaction_time"`
	Transaction_Status string    `json:"transaction_status"`
	Transaction_Id     string    `json:"transaction_id"`
	Status_Message     string    `json:"status_message"`
	Status_Code        int       `json:"status_code"`
	Signature_Key      string    `json:"signature_key"`
	Settlemen_Time     time.Time `json:"settlemen_time"`
	Payment_Type       string    `json:"payment_type"`
	Order_Id           string    `json:"order_id"`
	Gross_Amount       string    `json:"gross_amount"`
	Fraud_Status       string    `json:"fraud_status"`
	Currency           string    `json:"currency"`
	Approval_Code      string    `json:"approval_code"`
	Json_Code          string    `json:"json_code"`
}

type Order_Detail_Customer struct {
	gorm.Model
	Name         string
	Address      string
	Note_Address string
}

type Order_Detail_Deliverie struct {
	gorm.Model
	Driver        string
	Delivery_Date string
	Delivery_Time int
}

type Order_Detail_Payment_Method struct {
	gorm.Model
	Name               string
	Type               int
	Transaction_Number string
	Photo              string
}

type SummaryOrder struct {
	Id                int    `json:"id"`
	Order_Code        string `json:"order_code"`
	Order_Date        string `json:"order_date"`
	Delivery_Date     string `json:"delivery_date"`
	Delivery_Time     string `json:"delivery_time"`
	Total_Qty_Product int    `json:"total_qty_product"`
	Total_Price_Order int    `json:"total_price_order"`
	Status            int    `json:"status_order"`
}

type DetailPesanan struct {
	Id                int           `json:"id"`
	Order_Code        string        `json:"order_code"`
	Order_Date        string        `json:"order_date"`
	Expired_Date      string        `json:"expired_date"`
	Total_Price       int           `json:"total_price"`
	Payment_Method    int           `json:"payment_method"`
	Payment_Type      int           `json:"payment_type"`
	Status_Message    string        `json:"status_message"`
	Status_Code       string        `json:"status_code"`
	Rincian_Pemesanan []interface{} `json:"rincian_pemesanan"`
	Alamat_Pengiriman []interface{} `json:"alamat_pengiriman"`
	Waktu_Pengiriman  []interface{} `json:"waktu_pengiriman"`
}

type ResponseOrder struct {
	Orders           interface{} `json:"orders"`
	Detail_Product   interface{} `json:"detail_product"`
	Detail_Payment   interface{} `json:"detail_payment"`
	Detail_Receivers interface{} `json:"detail_receivers"`
}
