package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andy-gate/artaka-tenant-server/models"
	"github.com/gin-gonic/gin"
)

func ProductList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QueryProduct
	c.BindJSON(&query)

	var products []models.Product
  
	if err := models.MPosGORM.Raw("SELECT a.user_id, b.owner_name as tenant_name, a.name, a.units, a.quantity, a.sell_cost as price from products a join subscribers b on a.user_id = b.user_id where a.user_id = ?", query.User_id).Scan(&products).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (products != nil) {
	  c.JSON(http.StatusOK, products)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}