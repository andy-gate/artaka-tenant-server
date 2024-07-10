package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	"database/sql"

	_ "github.com/lib/pq"
)

var MPosDB *sql.DB
var MPosGORM *gorm.DB
var err error

func InitGormPostgres() {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	MPosGORM, err = gorm.Open("postgres", DBURL)
	if err != nil {
		panic(err)
	}
}