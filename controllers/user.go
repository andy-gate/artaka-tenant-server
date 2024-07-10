package controllers

import (
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andy-gate/artaka-tenant-server/models"
	"github.com/andy-gate/artaka-tenant-server/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
  
	var query models.QueryUser
	c.BindJSON(&query)

	b, err := json.Marshal(query)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

	var jsonStr = []byte(b)

	apiUrl := "https://artaka999.com/api/admin/loginTenant"

	req, err := http.NewRequest("POST",apiUrl,bytes.NewBuffer(jsonStr))
	if err != nil {   
		fmt.Printf("Request Failed: %s", err)
		return
	}

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

	defer resp.Body.Close()

	user := models.UserResponse{}
  	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &user)
	if(user.Success) {
		token, _ := utils.GenerateToken(user.Data.Id, user.Data.Username, user.Data.Referral_code)
		data := gin.H{"id": user.Data.Id, "username": user.Data.Username, "referral_code": user.Data.Referral_code, "token": token}
		c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "data": nil})
	}
}