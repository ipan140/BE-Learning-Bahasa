package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"learning-bahasa-be/app/modules/auth"
	"learning-bahasa-be/app/modules/customtable"
	"learning-bahasa-be/app/modules/phrase"
	"learning-bahasa-be/app/modules/preference"
	"learning-bahasa-be/app/modules/progress"
	"learning-bahasa-be/app/modules/target"
)

var DB *gorm.DB

func InitDatabase(databaseURL string) *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database PostgreSQL: %v", err)
	}

	log.Println("✓ Berhasil terhubung ke database PostgreSQL.")

	// Auto-Migrate 6 tabel fitur Learning Bahasa
	log.Println("Menjalankan Auto-Migration 6 tabel...")
	err = DB.AutoMigrate(
		&auth.User{},
		&target.UserTarget{},
		&customtable.UserCustomTable{},
		&progress.UserStudyProgress{},
		&phrase.UserCustomPhrase{},
		&phrase.UserPhraseCategory{},
		&preference.UserPreference{},
	)
	if err != nil {
		log.Fatalf("Gagal auto-migrate tabel database: %v", err)
	}

	log.Println("✓ Auto-Migration tabel selesai.")
	return DB
}
