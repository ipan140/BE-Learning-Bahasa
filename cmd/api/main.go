package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"learning-bahasa-be/app/modules/auth"
	"learning-bahasa-be/app/modules/customtable"
	"learning-bahasa-be/app/modules/phrase"
	"learning-bahasa-be/app/modules/preference"
	"learning-bahasa-be/app/modules/progress"
	"learning-bahasa-be/app/modules/target"
	"learning-bahasa-be/common/middleware"
	"learning-bahasa-be/common/utils"
	"learning-bahasa-be/config"
)

func main() {
	// 1. Muat konfigurasi .env
	cfg := config.InitConfig()

	// 2. Inisialisasi Database PostgreSQL & Auto-Migration
	db := config.InitDatabase(cfg.DatabaseURL)

	// 3. Inisialisasi Echo Engine
	e := echo.New()
	e.HideBanner = true

	// 4. Pasang Middleware Global
	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${uri} (${latency_human})\n",
	}))
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.CORSMiddleware())

	// Health Check Endpoint
	e.GET("/", func(c echo.Context) error {
		return utils.OK(c, map[string]interface{}{
			"service": "Learning Bahasa REST API",
			"status":  "running",
			"runtime": "Go (Echo Framework) - Modular Monolith",
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// 5. Inisialisasi API Route Group & Auth Middleware
	api := e.Group("/api")
	jwtAuth := middleware.AuthMiddleware(cfg.JWTSecret)

	// 6. Inisialisasi Setiap Modul Mandiri (Feature-First)
	log.Println("Menginisialisasi modul-modul API...")
	auth.InitModule(api, db, cfg.JWTSecret, jwtAuth)
	target.InitModule(api, db, jwtAuth)
	customtable.InitModule(api, db, jwtAuth)
	progress.InitModule(api, db, jwtAuth)
	phrase.InitModule(api, db, jwtAuth)
	preference.InitModule(api, db, jwtAuth)

	// 7. Jalankan Server
	log.Printf("🚀 Server Learning Bahasa Backend berjalan di port :%s\n", cfg.Port)
	if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server berhenti dengan error: %v", err)
	}
}
