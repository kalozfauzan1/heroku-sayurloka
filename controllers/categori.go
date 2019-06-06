package controllers

import (
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) GetCategorie(c *gin.Context) {
	var (
		categori []structs.Categorie
		result   gin.H
	)
	err := idb.DB.Find(&categori).Error
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
			"data":    categori,
		}

	}
	c.JSON(http.StatusOK, result)
}
