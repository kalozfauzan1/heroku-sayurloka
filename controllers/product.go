package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) ListProduct(c *gin.Context) {
	var (
		lis_product []structs.ListProduct
		result      gin.H
	)

	err := idb.DB.Raw("CALL ListProduct()").Scan(&lis_product).Error
	if err != nil {
		result = gin.H{
			"message": err.Error(),
			"code":    500,
			"data":    nil,
		}
		c.JSON(http.StatusOK, result)
		panic(err)
	} else {
		result = gin.H{
			"message": "Success",
			"code":    200,
			"data":    lis_product,
		}

	}
	c.JSON(http.StatusOK, result)
}
