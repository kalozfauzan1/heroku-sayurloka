package list_product

import (
	"../../structs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) ListProduct(c *gin.Context) {
	var (
		list_product []structs.ListProduct // get struct list product
		result       gin.H
	)
	// code response api, error or no
	code := 200

	// message response api, if have error or no
	message := "Success"

	// get type for type response
	param_type := c.Param("type")

	fmt.Println(param_type)

	if param_type == "1" { // type vegetables
		// get list product vegetables
		err_list_product := idb.DB.Raw("CALL ListProductCategory(1)").Scan(&list_product).Error

		// check process get data product vegetables if have error
		if err_list_product != nil {
			code = 500                         // to change code api be error
			message = err_list_product.Error() // show message error to response
			panic(err_list_product)
		}

	} else {
		// type 1 = vegetable, 2 = fruits, 3 = spices, 4 = basic food, 5 = practical package
		// get list product where type
		err_list_product := idb.DB.Raw("CALL ListProductCategory(?)", param_type).Scan(&list_product).Error

		// check process get data product where type if have error
		if err_list_product != nil {
			code = 500                         // to change code api be error
			message = err_list_product.Error() // show message error to response
			panic(err_list_product)
		}
	}

	result = gin.H{
		"message": message,
		"code":    code,
		"data":    list_product,
	}

	c.JSON(http.StatusOK, result)
}
