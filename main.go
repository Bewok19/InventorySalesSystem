package main

import (
	"log"
	"myapp/config"
	"myapp/router"
)

func main() {
    // Inisialisasi database
    config.InitDB()
    log.Println("Database connection established")

    // Inisialisasi router
    r := router.SetupRouter()

    // Jalankan server
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
