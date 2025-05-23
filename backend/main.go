package main

import (
	"log"
	"os"

	"gimnasio-app/config"
	"gimnasio-app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configurar variables de entorno por defecto
	if os.Getenv("DB_HOST") == "" {
		os.Setenv("DB_HOST", "localhost")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "3306")
	}
	if os.Getenv("DB_USER") == "" {
		os.Setenv("DB_USER", "root")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "root")
	}
	if os.Getenv("DB_NAME") == "" {
		os.Setenv("DB_NAME", "gimnasio_db")
	}

	// Inicializar la base de datos
	config.InitDB()

	// Configurar el router
	r := gin.Default()

	// Configurar CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Rutas públicas
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	// Rutas de actividades (públicas)
	r.GET("/actividades", controllers.GetActivities)
	r.GET("/actividades/:id", controllers.GetActivityByID)

	// Rutas protegidas (requieren autenticación)
	authorized := r.Group("/")
	authorized.Use( /* middleware de autenticación */ )
	{
		// Rutas de inscripciones
		authorized.POST("/inscripciones", controllers.CreateEnrollment)
		authorized.GET("/mis-inscripciones/:userID", controllers.GetUserEnrollments)
		authorized.DELETE("/inscripciones/:id", controllers.CancelEnrollment)

		// Rutas de administrador
		admin := authorized.Group("/admin")
		admin.Use( /* middleware de verificación de rol admin */ )
		{
			admin.POST("/actividades", controllers.CreateActivity)
			admin.PUT("/actividades/:id", controllers.UpdateActivity)
			admin.DELETE("/actividades/:id", controllers.DeleteActivity)
		}
	}

	// Iniciar el servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
