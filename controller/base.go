package controller

import (
	"net/http"
	"vinid_project/database"

	"github.com/gin-gonic/gin"
	"github.com/jasonwinn/geocoder"
)

type Controller struct {
	dao *database.Dao
}

func NewController(dao *database.Dao) *Controller {
	return &Controller{
		dao: dao,
	}
}

func (this *Controller) TestGeoCoder(c *gin.Context) {
	address, err := geocoder.ReverseGeocode(20.9645230000, 105.8250960000)
	if err != nil {
		panic("THERE WAS SOME ERROR!!!!!")
	}

	c.JSON(http.StatusOK, address.Street+", "+address.City)

	// address.Street 	        // 542 Marion St
	// address.City 		        // Seattle
	// address.State 	        // WA
	// address.PostalCode 	    // 98104
	// address.County 	        // King
	// address.CountryCode       // US
}
