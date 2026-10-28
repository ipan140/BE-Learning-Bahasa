# Learning Bahasa - Backend REST API (Go + Echo)

Backend REST API mandiri dan berkinerja tinggi untuk platform **Learning Bahasa**, dibangun menggunakan **Go (Golang)** dan framework **Echo v4**. Mengadopsi arsitektur **Feature-First Modular Monolith** dengan pendekatan **Clean 5-Tier Layered Architecture** pada setiap modul bisnisnya.

Database dirancang **100% kompatibel dan identik dengan skema database Supabase** menggunakan **PostgreSQL** dan **GORM Auto-Migration**.

---

## 📑 Daftar Isi
- [Arsitektur & Pola Desain](#-arsitektur--pola-desain)
- [Struktur Direktori](#-struktur-direktori)
- [Skema Database & Model](#-skema-database--model)
- [Panduan Memulai (Lokal)](#-panduan-memulai-lokal)
- [Dokumentasi Interaktif (Swagger UI)](#-dokumentasi-interaktif-swagger-ui)
- [Daftar Endpoint REST API](#-daftar-endpoint-rest-api)
- [Panduan Deploy ke VPS (Linux Production)](#-panduan-deploy-ke-vps-linux-production)
- [Variabel Environment (.env)](#-variabel-environment-env)

---

## 🏗️ Arsitektur & Pola Desain

### 1. Modular Monolith (Feature-First)
- **Single Deployable Unit:** Seluruh backend dikompilasi menjadi satu file binary executable mandiri (`server` / `server.exe`). Sangat hemat sumber daya (RAM < 30 MB pada VPS).
- **Domain Encapsulation:** Logika bisnis dipisahkan mandiri per fitur dalam folder `app/modules/` (`auth`, `target`, `customtable`, `progress`, `phrase`, `preference`).

### 2. Clean 5-Tier Layered Architecture (per Modul)
Setiap modul menerapkan pemisahan tanggung jawab yang tegas:
```text
[ HTTP Request dari Frontend / Client ]
                  │
                  ▼
         route.go & handler.go      ──► Presentation Layer (Routing Echo, Binding, Swagger Annotations, JSON Response)
                  │
                  ▼
              service.go            ──► Business Logic Layer (Validasi Aturan & Alur Data Bisnis)
                  │
                  ▼
            repository.go           ──► Data Access Layer (GORM DB Queries & Clauses)
                  │
                  ▼
               model.go             ──► Entity Layer (GORM Struct, JSON Tags, Skema Pemetaan Tabel)
```

---

## 📁 Struktur Direktori

```text
D:\BE learning-bahasa\
├── cmd/
│   └── api/
│       └── main.go                 # Entry point: Inisialisasi Echo, DB, Swagger UI, & Router
├── config/
│   ├── config.go                   # Load konfigurasi murni dari .env (Zero hardcoded secret)
│   └── database.go                 # Inisialisasi PostgreSQL GORM & AutoMigrate 7 tabel
├── common/
│   ├── middleware/
│   │   ├── auth.go                 # Middleware JWT Token & Claims Context
│   │   └── cors.go                 # Konfigurasi Cross-Origin Resource Sharing (CORS)
│   └── utils/
│       └── response.go             # Standar struktur respon JSON (Success, Error, OK)
├── docs/
│   ├── docs.go                     # Hasil kompilasi Swaggo
│   ├── swagger.json                # Spesifikasi OpenAPI JSON
│   └── swagger.yaml                # Spesifikasi OpenAPI YAML
├── app/
│   └── modules/
│       ├── auth/                   # Modul Akun, JWT Token, Google OAuth, Profil
│       ├── target/                 # Modul Negara Berbintang, Batch Sync, Catatan Motivasi
│       ├── customtable/            # Modul Studio Aksara Kustom (Grid Susunan Huruf JSONB)
│       ├── progress/               # Modul Huruf Hafal, Streak Count, Riwayat Skor Kuis
│       ├── phrase/                 # Modul Frasa Kustom & Kategori Kartu Frasa
│       └── preference/             # Modul Setelan Audio (Speed, Toggle) & Tema
├── .env                            # Variabel lingkungan lokal
├── .env.example                    # Template variabel lingkungan
├── .gitignore                      # Git ignore file (binary, log, .env diabaikan)
├── go.mod                          # Go module dependencies
└── go.sum                          # Checksum dependencies
```

---

## 🗄️ Skema Database & Model

Backend ini mengelola **7 tabel relasional** di PostgreSQL dengan tipe data yang presisi menyesuaikan schema Supabase:

| Nama Tabel | Primary Key | Foreign Key | Kolom Kunci Lainnya |
| :--- | :--- | :--- | :--- |
| `users` | `id` (uuid) | - | `email`, `password_hash`, `display_name`, `avatar_url`, `created_at` |
| `user_targets` | `id` (uuid) | `user_id` (uuid) | `country_id`, `notes`, `created_at` |
| `user_custom_tables` | `id` (text) | `user_id` (uuid) | `name`, `script`, `description`, `rows` (jsonb), `created_at` |
| `user_study_progress` | `id` (uuid) | `user_id` (uuid) | `table_id`, `memorized_chars` (jsonb), `score`, `streak_count`, `last_practiced_at`, `metadata` (jsonb) |
| `user_custom_phrases` | `id` (text) | `user_id` (uuid) | `language_id`, `categories` (jsonb), `updated_at` |
| `user_phrase_categories` | `id` (uuid) | `user_id` (uuid) | `name`, `icon`, `created_at` |
| `user_preferences` | `user_id` (uuid) | `user_id` (uuid) | `theme`, `audio_speed` (numeric), `audio_enabled` (bool), `updated_at` |

> [!NOTE]
> **Auto-Migration Aktif:** Anda **tidak perlu membuat tabel secara manual**. Saat server backend pertama kali dijalankan dan terhubung ke PostgreSQL, seluruh tabel, kolom, index, dan primary key akan dibuatkan secara otomatis oleh GORM.

---

## 🚀 Panduan Memulai (Lokal)

### 1. Prasyarat
- **Go:** Versi 1.25 atau lebih baru
- **PostgreSQL:** Berjalan secara lokal (port 5432) atau via container Docker

### 2. Buat Database Kosong di PostgreSQL
Jalankan di `psql`, pgAdmin, atau DBeaver:
```sql
CREATE DATABASE learning_bahasa;
```

### 3. Konfigurasi File `.env`
Salin template atau sesuaikan file `.env`:
```env
PORT=8080
DATABASE_URL=postgres://postgres:password_anda@localhost:5432/learning_bahasa?sslmode=disable
JWT_SECRET=kunci-rahasia-jwt-panjang-dan-aman
GOOGLE_CLIENT_ID=
FRONTEND_URL=http://localhost:4321
```

### 4. Unduh Dependensi
```bash
go mod tidy
```

### 5. Jalankan Server
```bash
# Menggunakan go run:
go run cmd/api/main.go

# Atau build binary lokal Windows:
go build -o server.exe cmd/api/main.go
.\server.exe
```

Server akan aktif di `http://localhost:8080`.

---

## 📖 Dokumentasi Interaktif (Swagger UI)

Semua handler API telah dilengkapi anotasi Swaggo standar OpenAPI. Anda dapat mencoba dan menguji setiap endpoint secara visual:

👉 **URL Swagger UI:** [http://localhost:8080/swagger](http://localhost:8080/swagger)

### Cara Generate Ulang Swagger Docs (Jika Ada Penambahan Endpoint):
```bash
swag init --parseDependency -g cmd/api/main.go -o docs
```

---

## 📡 Daftar Endpoint REST API

### 1. Autentikasi (`/api/auth`)
- `POST /api/auth/register` - Pendaftaran akun baru (Email, Password, Display Name)
- `POST /api/auth/login` - Masuk dengan email & password (Kembalikan JWT Token)
- `POST /api/auth/google` - Masuk menggunakan Google OAuth
- `GET  /api/auth/me` - Ambil profil pengguna login saat ini `[BearerAuth]`
- `PUT  /api/auth/profile` - Perbarui Display Name & Avatar `[BearerAuth]`

### 2. Target Negara (`/api/targets`)
- `GET    /api/targets` - Ambil daftar negara berbintang & catatan `[BearerAuth]`
- `POST   /api/targets` - Beri bintang pada satu negara baru `[BearerAuth]`
- `DELETE /api/targets/{countryId}` - Hapus bintang dari negara target `[BearerAuth]`
- `POST   /api/targets/batch` - Sinkronisasi massal seluruh target dari localStorage `[BearerAuth]`
- `PUT    /api/targets/{countryId}/notes` - Simpan catatan motivasi / rencana belajar negara `[BearerAuth]`

### 3. Studio Aksara Kustom (`/api/custom-tables`)
- `GET    /api/custom-tables` - Ambil semua tabel aksara buatan user `[BearerAuth]`
- `POST   /api/custom-tables` - Simpan / buat tabel aksara baru (termasuk grid baris huruf JSONB) `[BearerAuth]`
- `DELETE /api/custom-tables/{id}` - Hapus tabel aksara kustom `[BearerAuth]`

### 4. Progres Belajar & Kuis (`/api/progress` & `/api/quiz-scores`)
- `GET  /api/progress` - Ambil riwayat karakter yang sudah ditandai hafal `[BearerAuth]`
- `POST /api/progress` - Simpan data karakter yang baru dihafal & streak count `[BearerAuth]`
- `POST /api/quiz-scores` - Simpan skor kuis latihan & streak `[BearerAuth]`

### 5. Frasa Kustom (`/api/phrases` & `/api/phrase-categories`)
- `GET    /api/phrases` - Ambil daftar frasa kustom per bahasa `[BearerAuth]`
- `POST   /api/phrases` - Simpan kumpulan frasa kustom baru `[BearerAuth]`
- `DELETE /api/phrases/{languageId}` - Hapus frasa kustom bahasa tertentu `[BearerAuth]`
- `GET    /api/phrase-categories` - Ambil daftar kategori kartu frasa `[BearerAuth]`
- `POST   /api/phrase-categories` - Tambah kategori kartu frasa baru `[BearerAuth]`
- `DELETE /api/phrase-categories/{name}` - Hapus kategori kartu frasa `[BearerAuth]`

### 6. Pengaturan & Preferensi (`/api/preferences`)
- `GET /api/preferences` - Ambil preferensi pengguna (kecepatan audio, status audio, tema) `[BearerAuth]`
- `PUT /api/preferences` - Simpan pembaruan preferensi `[BearerAuth]`

---

## 🐧 Panduan Deploy ke VPS (Linux Production)

Keuntungan utama menggunakan Go adalah binary mandiri yang **sangat mudah dideploy ke VPS Ubuntu/Debian tanpa perlu install runtime Go, Node.js, atau npm**.

### 1. Kompilasi Cross-Compile Binary untuk Linux 64-bit
Jalankan perintah ini di PowerShell Windows Anda:
```powershell
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -ldflags="-w -s" -o server cmd/api/main.go
```
*Akan menghasilkan satu file binary Linux mandiri bernama `server` (~15 MB).*

### 2. Upload ke VPS
Upload file `server` dan file `.env` ke folder di VPS Anda (misalnya `/var/www/learning-bahasa-be`):
```bash
scp server .env user@ip-vps-anda:/var/www/learning-bahasa-be/
```

### 3. Beri Izin Eksekusi di VPS
```bash
ssh user@ip-vps-anda
cd /var/www/learning-bahasa-be
chmod +x server
```

### 4. Pasang Systemd Service (Agar Otomatis Jalan di Background)
Buat file service systemd:
```bash
sudo nano /etc/systemd/system/learning-bahasa.service
```

Isi dengan konfigurasi berikut:
```ini
[Unit]
Description=Learning Bahasa Backend Service
After=network.target postgresql.service

[Service]
Type=simple
User=root
WorkingDirectory=/var/www/learning-bahasa-be
ExecStart=/var/www/learning-bahasa-be/server
Restart=always
RestartSec=5
EnvironmentFile=/var/www/learning-bahasa-be/.env

[Install]
WantedBy=multi-user.target
```

Aktifkan dan jalankan servicenya:
```bash
sudo systemctl daemon-reload
sudo systemctl enable learning-bahasa
sudo systemctl start learning-bahasa
sudo systemctl status learning-bahasa
```

---

## ⚙️ Variabel Environment (.env)

| Variabel | Deskripsi | Default / Contoh |
| :--- | :--- | :--- |
| `PORT` | Port HTTP yang digunakan server | `8080` |
| `DATABASE_URL` | String koneksi PostgreSQL | `postgres://user:pass@host:5432/dbname?sslmode=disable` |
| `JWT_SECRET` | Kunci enkripsi token autentikasi JWT | `string-acak-panjang-dan-rahasia` |
| `GOOGLE_CLIENT_ID` | Client ID Google OAuth (opsional) | `123456...apps.googleusercontent.com` |
| `FRONTEND_URL` | URL Frontend untuk izin CORS | `https://learning-bahasa-news.vercel.app` |

---

## 📄 Lisensi
Didistribusikan di bawah lisensi MIT.
