package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetPromotion(c *gin.Context) {
	var (
		promotion []structs.Promotion
		result    gin.H
	)
	err := idb.DB.Find(&promotion).Error
	if err != nil {
		result = gin.H{
			"message": err,
			"code":    500,
			"data":    nil,
		}
		c.JSON(http.StatusOK, result)
		panic(err)
	} else {
		result = gin.H{
			"message": "success",
			"code":    200,
			"data":    promotion,
		}

	}
	c.JSON(http.StatusOK, result)
}
