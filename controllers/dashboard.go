package controllers

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/andy-gate/artaka-tenant-server/models"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QueryDashboard
	c.BindJSON(&query)

	var dashboard models.Dashboard
  
	if err := models.MPosGORM.Raw("SELECT count(a.*) as tenant_count from subscribers a join (SELECT DISTINCT on (user_id) * from outlets) b on a.user_id = b.user_id where referral_code ilike ?", query.Referral_code).Scan(&dashboard).Error; err != nil {
		fmt.Printf("error count tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	
	if err := models.MPosGORM.Raw("SELECT count(a.*) as sales_count FROM sales a JOIN (SELECT user_id FROM subscribers where referral_code ilike ?) b ON a.user_id=b.user_id", query.Referral_code).Scan(&dashboard).Error; err != nil {
		fmt.Printf("error count sales: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	
	if err := models.MPosGORM.Raw("SELECT sum(a.total_bill) as sales_total FROM sales a JOIN (SELECT user_id FROM public.subscribers where referral_code ilike ?) b ON a.user_id=b.user_id", query.Referral_code).Scan(&dashboard).Error; err != nil {
		fmt.Printf("error count sales: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := models.MPosGORM.Raw("SELECT count(a.*) as inventory_count FROM products a JOIN (SELECT user_id FROM subscribers where referral_code ilike ?) b ON a.user_id=b.user_id", query.Referral_code).Scan(&dashboard).Error; err != nil {
		fmt.Printf("error count sales: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dashboard)
}