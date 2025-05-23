package main

import (
	"log"
	"os"

	"gimnasio-app/config"
)

func testDB() {
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
	log.Println("Conexión a la base de datos exitosa!")
}
