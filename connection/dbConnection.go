package connection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DSN = "root:eTmzFGVqnyRPratogumTycVamwXDKfgh@tcp(junction.proxy.rlwy.net:36199)/railway?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectionDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Base de Datos corriendo...")

} 
