package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andy-gate/artaka-tenant-server/models"
	"github.com/gin-gonic/gin"
)

func SalesList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QuerySales
	c.BindJSON(&query)

	var sales []models.Sales
  
	if err := models.MPosGORM.Raw("select b.owner_name as tenant_name, a.create_dtm::date, count(a.id) as total_trx, sum(a.total_bill) as total_amount from sales a join subscribers b on a.user_id = b.user_id where a.user_id = ? AND a.create_dtm > ? AND a.create_dtm < ? group by b.owner_name, a.create_dtm::date", query.User_id, query.Start_date, query.End_date).Scan(&sales).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (sales != nil) {
	  c.JSON(http.StatusOK, sales)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}

func SalesListDetail(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QuerySales
	c.BindJSON(&query)

	var sales []models.SalesDetail
  
	if err := models.MPosGORM.Raw("SELECT * from sales where user_id = ? AND create_dtm > ? AND create_dtm < ?", query.User_id, query.Start_date, query.End_date).Scan(&sales).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (sales != nil) {
	  c.JSON(http.StatusOK, sales)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}