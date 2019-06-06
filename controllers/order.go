package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) ListOrder(c *gin.Context) {
	var (
		detail_product   []structs.Order_Detail_Product
		detail_payment   structs.Order_Detail_Payment
		detail_receivers structs.Order_Detail_Receiver
		summary_order    structs.SummaryOrder
		_response        structs.ResponseOrder

		result gin.H
	)
	customer_id := c.Param("customer_id")

	err := idb.DB.Raw("CALL MyOrder(?)", customer_id).Scan(&detail_payment).Scan(&detail_product).Scan(&detail_receivers).Scan(&summary_order).Error
	if err != nil {
		result = gin.H{
			"message": err.Error(),
			"code":    500,
			"data":    nil,
		}
		c.JSON(http.StatusOK, result)
		panic(err)
	} else {
		_response.Detail_Product = detail_product
		_response.Detail_Payment = detail_payment
		_response.Detail_Receivers = detail_receivers
		_response.Orders = summary_order
		result = gin.H{
			"message": "Success",
			"code":    200,
			"data":    _response,
		}

	}
	c.JSON(http.StatusOK, result)
}
