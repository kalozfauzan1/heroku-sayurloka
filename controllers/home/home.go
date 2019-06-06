package home

import (
	"../../structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type InDB struct {
	DB *gorm.DB
}

//func (idb *InDB) ListDropPoints(c *gin.Context) {
//	var (
//		drop_point []structs.Drop_Point
//		result   gin.H
//	)
//	err := idb.DB.Find(&drop_point).Error
//	if err != nil {
//		result = gin.H{
//			"message": err,
//			"code":    500,
//			"data":    nil,
//		}
//		c.JSON(http.StatusOK, result)
//		panic(err)
//	} else {
//		result = gin.H{
//			"message": "success",
//			"code":    200,
//			"data":    drop_point,
//		}
//
//	}
//	c.JSON(http.StatusOK, result)
//}

func (idb *InDB) ListHome(c *gin.Context) {
	var (
		drop_point        []structs.Drop_Point
		promotion         []structs.Promotion
		sent_today        []structs.ListProduct
		vegetables        []structs.ListProduct
		fruits            []structs.ListProduct
		spices            []structs.ListProduct
		practical_package []structs.ListProduct
		data              gin.H
		result            gin.H
	)
	// code response api, error or no
	code := 200

	// message response api, if have error or no
	message := "Success"

	// get list promotion
	err_promotion := idb.DB.Find(&promotion).Error

	// check process get data promotions if have error
	if err_promotion != nil {
		code = 500                      // to change code api be error
		message = err_promotion.Error() // show message error to response
		panic(err_promotion)
	}

	// get list drop_point
	err_drop_point := idb.DB.Find(&drop_point).Error

	// check process get data drop_point if have error
	if err_drop_point != nil {
		code = 500                       // to change code api be error
		message = err_drop_point.Error() // show message error to response
		panic(err_drop_point)
	}

	// get list product can sent today
	err_sent_today := idb.DB.Raw("CALL ListProductCategory(1)").Scan(&sent_today).Error

	// check process get data product sent today if have error
	if err_sent_today != nil {
		code = 500                       // to change code api be error
		message = err_sent_today.Error() // show message error to response
		panic(err_sent_today)
	}

	// get list product can vegetables
	err_vegetables := idb.DB.Raw("CALL ListProductCategory(1)").Scan(&vegetables).Error

	// check process get data product vegetables if have error
	if err_vegetables != nil {
		code = 500                       // to change code api be error
		message = err_vegetables.Error() // show message error to response
		panic(err_vegetables)
	}

	// get list product can fruits
	err_fruits := idb.DB.Raw("CALL ListProductCategory(1)").Scan(&fruits).Error

	// check process get data product fruits if have error
	if err_fruits != nil {
		code = 500                   // to change code api be error
		message = err_fruits.Error() // show message error to response
		panic(err_fruits)
	}

	// get list product can spices
	err_spices := idb.DB.Raw("CALL ListProductCategory(1)").Scan(&spices).Error

	// check process get data product spices if have error
	if err_spices != nil {
		code = 500                   // to change code api be error
		message = err_spices.Error() // show message error to response
		panic(err_spices)
	}

	// get list product can practical_package
	err_practical_package := idb.DB.Raw("CALL ListProductCategory(1)").Scan(&practical_package).Error

	// check process get data product practical_package if have error
	if err_practical_package != nil {
		code = 500                              // to change code api be error
		message = err_practical_package.Error() // show message error to response
		panic(err_practical_package)
	}

	data = gin.H{
		"drop_points":       drop_point,        //list drop point
		"promotions":        promotion,         // list promotion
		"sent_today":        sent_today,        // list product can sent today
		"practical_package": practical_package, // list product practical package
		"vegetables":        vegetables,        // list product vegetables
		"fruits":            fruits,            // list product fruits
		"spices":            spices,            // list product spices
	}

	result = gin.H{
		"message": message, // getting every message response
		"code":    code,    // getting every code response
		"data":    data,
	}
	c.JSON(http.StatusOK, result)
}
