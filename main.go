package main

import (
	"myapp/app"
	"myapp/config"

	"github.com/gin-gonic/gin"
)

func main() {
    db := config.ConnectDatabase() // Pastikan koneksi DB berhasil
    r := gin.Default()
    app.RegisterRoutes(r, db) // Pastikan RegisterRoutes dipanggil di sini
    r.Run(":8080")            // Jalankan server pada port 8080
}
