package main

import (
	"fmt"
	"os"
	"time"

	"github.com/andy-gate/artaka-tenant-server/controllers"
	"github.com/andy-gate/artaka-tenant-server/middlewares"
	"github.com/andy-gate/artaka-tenant-server/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())

	if err := godotenv.Load(`.env`); err != nil {
		panic(err)
	}

	models.InitGormPostgres()
	defer models.MPosGORM.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api")
	api.POST("/login", controllers.Login)

	protected:= router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/tenant_list", controllers.TenantList)
	protected.POST("/tenant_list_dropdown", controllers.ActiveTenantList)
	protected.POST("/dashboard", controllers.Dashboard)
	protected.POST("/product_list", controllers.ProductList)
	protected.POST("/sales_list", controllers.SalesList)
	protected.POST("/sales_list_detail", controllers.SalesListDetail)
	protected.POST("/change_status", controllers.ChangeTenantRefCode)
	
	fmt.Printf("Listening to port %s", os.Getenv("PORT1"))
	router.Run(":" + os.Getenv("PORT1"))
}