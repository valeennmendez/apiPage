package main


import (
	"net/http"
	//"os"
	//"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valeennmendez/api-go/connection"
	"github.com/valeennmendez/api-go/models"
	"github.com/valeennmendez/api-go/routes"
)

func main() {
	connection.ConnectionDB()

	connection.DB.AutoMigrate(&models.Patients{})
	connection.DB.AutoMigrate(&models.Admin{})
	connection.DB.AutoMigrate(&models.Appoinment{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://codepages.onrender.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Servir archivos estáticos
	r.Static("/static", "./static")
	r.StaticFile("/login.html", "./pages/login.html")

	// Rutas de autenticación
	r.POST("/register", routes.RegisterUser)
	r.POST("/login", routes.Login)
	r.GET("/validate", routes.ValidateSession)
	r.POST("/logout", routes.CloseSesion)

	// Rutas protegidas
	protected := r.Group("/")
	protected.Use(routes.AuthMiddleware())
	{
		protected.StaticFile("/index.html", "./pages/index.html")
		protected.GET("/patients", routes.GetAllPatients)
		protected.GET("/patients/:id", routes.GetPatientByID)
		protected.POST("/create", routes.CreatePatient)
		protected.PUT("/edit/:id", routes.EditPatient)
		protected.DELETE("/patients/:id", routes.DeletePacients)
		protected.GET("/total-patients", routes.TotalPatientsData)
		protected.POST("/create-appointment", routes.CreateAppoinment)
		protected.GET("/appointment-today", routes.AppointmentToday)
		protected.GET("/available-hours", routes.GetAviableHours)
		protected.GET("/appointments", routes.GetAllAppointments)
		protected.GET("/appointments-week", routes.AppointmentWeek)
		protected.POST("/approve-user/:id", routes.ApproveUser)
		protected.POST("/decline-user/:id", routes.DeclineUser)
		protected.GET("/admins", routes.GetAllAdmins)

	}
	r.GET("/search-patient", routes.SearchPatient)
	r.DELETE("/cancel-appointment/:id", routes.CancelAppointment) // <--- DEBE ESTAR PUBLICA SI O SI.


	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Corriendo",
		})
	})
/*
 	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Para desarrollo local
	}
*/

	r.Run(":8000")
}

