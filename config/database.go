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
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Printf("⚠️ PERINGATAN: Belum dapat terhubung ke PostgreSQL (%v)\n", err)
		log.Println("👉 Server tetap berjalan. Sesuaikan password database di file .env untuk mengaktifkan koneksi database.")
		return nil
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
		log.Printf("⚠️ Peringatan saat auto-migrate tabel: %v\n", err)
	} else {
		log.Println("✓ Auto-Migration tabel database selesai.")
	}

	return DB
}
