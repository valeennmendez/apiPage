package connection

import (
	"fmt"

	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// const DSN = "root:@tcp(127.0.0.1:3306)/Acneclinic?charset=utf8mb4&parseTime=True&loc=UTC"
const DSN = "postgresql://postgres.vfbshxvqidlrawostjvn:Repeatrave10_@aws-0-us-west-1.pooler.supabase.com:6543/postgres"

func ConnectionDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Base de Datos corriendo...")

}
 
