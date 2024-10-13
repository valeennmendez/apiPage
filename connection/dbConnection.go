import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

func ConnectionDB() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		panic("Error cargando el archivo .env")
	}

	// Leer las variables de entorno
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// Construir el DSN
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require", user, password, host, port, dbname)

	// Abrir la conexión con GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Base de Datos corriendo...")
}
