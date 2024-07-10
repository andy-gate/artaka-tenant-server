package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andy-gate/artaka-tenant-server/models"
	"github.com/gin-gonic/gin"
)

func TenantList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QueryTenant
	c.BindJSON(&query)

	var tenants []models.Tenant
  
	if err := models.MPosGORM.Raw("SELECT a.user_id, b.nama, b.address, a.referral_code from subscribers a join (SELECT DISTINCT on (user_id) * from outlets where outlet_id = 'OTL-001') b on a.user_id = b.user_id where referral_code ilike '%'||?||'%' order by length(referral_code), b.nama asc", query.Referral_code).Scan(&tenants).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (tenants != nil) {
	  c.JSON(http.StatusOK, tenants)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}

func ActiveTenantList(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QueryTenant
	c.BindJSON(&query)

	var tenants []models.Tenant
  
	if err := models.MPosGORM.Raw("SELECT a.user_id, b.outlet_id, a.owner_name as nama, b.address, a.referral_code from subscribers a join outlets b on a.user_id = b.user_id AND b.outlet_id = 'OTL-001' where referral_code ilike ?", query.Referral_code).Scan(&tenants).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	if (tenants != nil) {
	  c.JSON(http.StatusOK, tenants)
	} else {
	  c.JSON(http.StatusOK, json.RawMessage(`[]`))
	}	
}

func ChangeTenantRefCode(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var tenant models.Tenant
	c.BindJSON(&tenant)
  
	if err := models.MPosGORM.Raw("UPDATE subscribers SET referral_code = ? WHERE user_id = ? RETURNING user_id, referral_code", tenant.Referral_code, tenant.User_id).Scan(&tenant).Error; err != nil {
		fmt.Printf("error list tenant: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  
	c.JSON(http.StatusOK, tenant)
}