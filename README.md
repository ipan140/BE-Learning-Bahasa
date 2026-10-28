# Learning Bahasa - Backend API (Go + Echo)

Backend REST API mandiri untuk platform **Learning Bahasa**, dibangun menggunakan arsitektur **Modular Monolith** dengan pendekatan **Clean Layered Architecture (N-Tier)** per modul fitur.

---

## 🏗️ Pola Arsitektur

### 1. Pola Utama: Modular Monolith
* **Single Deployable Unit:** Seluruh backend dikompilasi menjadi satu file binary executable mandiri (`server` / `server.exe`), sangat hemat memori (RAM < 30 MB pada VPS).
* **Domain Encapsulation:** Kode dipisahkan per modul bisnis di `app/modules/` (`auth`, `target`, `customtable`, `progress`, `phrase`, `preference`).

### 2. Pola Internal Modul: 5 Lapisan (Clean Layered)
Setiap folder modul menerapkan 5 lapisan pemisahan tanggung jawab:
```
[ HTTP Request dari Frontend ]
            │
            ▼
   route.go & handler.go   ──► Presentation Layer (Routing Echo, Binding, JSON Response)
            │
            ▼
        service.go          ──► Business Logic Layer (Validasi Aturan & Alur Data)
            │
            ▼
      repository.go         ──► Data Access Layer (GORM DB Queries & Clauses)
            │
            ▼
         model.go           ──► Entity Layer (GORM Struct, JSON Tags, Schema Mapping)
```

---

## 📁 Struktur Direktori

```text
D:\BE learning-bahasa\
├── cmd/
│   └── api/
│       └── main.go              # Entry point: inisialisasi Echo, DB, router, & start server
├── app/
│   ├── config/                  # Load file .env
│   ├── database/                # Koneksi GORM PostgreSQL & AutoMigrate 6 tabel
│   ├── middleware/              # JWT Auth Middleware & klaim context
│   └── modules/
│       ├── auth/                # Login, Register, Google OAuth, Profile
│       ├── target/              # Target negara berbintang, batch sync, catatan
│       ├── customtable/         # CRUD tabel aksara kustom (grid huruf JSONB)
│       ├── progress/            # Huruf hafal & skor kuis (flashcard & listening)
│       ├── phrase/              # Frasa kustom harian & kategori frasa
│       └── preference/          # Setelan kecepatan suara, toggle audio, tema
├── pkg/
│   └── response/                # Standar response JSON
├── .env                         # Konfigurasi database & JWT
├── go.mod
└── go.sum
```

---

## 🚀 Panduan Memulai (Lokal)

### 1. Unduh Dependensi
```bash
go mod tidy
```

### 2. Konfigurasi Database
Buka file `.env` dan pastikan URL PostgreSQL sudah sesuai:
```env
DATABASE_URL=postgres://postgres:password@localhost:5432/learning_bahasa?sslmode=disable
```

### 3. Jalankan Server
```bash
go run cmd/api/main.go
```
Server akan otomatis membuat 6 tabel database dan siap melayani request di port `http://localhost:8080`.

---

## 🐧 Kompilasi untuk Deploy ke VPS (Linux 64-bit)

Salah satu keunggulan terbesar Go adalah kemudahan deploy ke VPS tanpa perlu install Node.js/npm:

```bash
# Jalankan di terminal PowerShell Anda:
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -ldflags="-w -s" -o server cmd/api/main.go
```
* Perintah ini menghasilkan 1 file binary Linux bernama `server` (ukuran ~15 MB).
* Cukup upload file `server` dan `.env` ke VPS Anda, lalu jalankan:
```bash
chmod +x server
./server
```

---

## 📡 Daftar Endpoint API

### 1. Autentikasi (`/api/auth`)
* `POST /api/auth/register` - Daftar akun baru (Email, Password, Display Name)
* `POST /api/auth/login` - Masuk dengan email & sandi (Mengembalikan JWT token)
* `POST /api/auth/google` - Masuk lewat akun Google OAuth
* `GET  /api/auth/me` - Ambil profil pengguna yang sedang login `[Wajib Token]`
* `PUT  /api/auth/profile` - Perbarui nama tampilan & foto profil avatar `[Wajib Token]`

### 2. Target Negara (`/api/targets`) `[Wajib Token]`
* `GET    /api/targets` - Ambil daftar negara berbintang & catatan impian
* `POST   /api/targets` - Beri bintang pada negara target baru
* `DELETE /api/targets/:countryId` - Hapus bintang dari negara target
* `POST   /api/targets/batch` - Sinkronisasi massal seluruh target lokal
* `PUT    /api/targets/:countryId/notes` - Simpan catatan rencana / motivasi belajar negara

### 3. Tabel Aksara Kustom (`/api/custom-tables`) `[Wajib Token]`
* `GET    /api/custom-tables` - Ambil seluruh tabel aksara buatan user
* `POST   /api/custom-tables` - Simpan / buat tabel aksara baru (termasuk susunan kotak huruf)
* `PUT    /api/custom-tables/:id` - Perbarui tabel aksara saat Mode Edit aktif
* `DELETE /api/custom-tables/:id` - Hapus tabel aksara kustom

### 4. Progres Belajar & Kuis (`/api/progress` & `/api/quiz-scores`) `[Wajib Token]`
* `GET  /api/progress` - Ambil riwayat huruf yang sudah ditandai hafal (bergaris hijau)
* `POST /api/progress` - Simpan progres huruf yang baru dihafal
* `POST /api/quiz-scores` - Simpan skor kuis & streak latihan

### 5. Frasa Kustom (`/api/phrases` & `/api/phrase-categories`) `[Wajib Token]`
* `GET    /api/phrases` - Ambil seluruh frasa kustom per bahasa
* `POST   /api/phrases` - Simpan kumpulan frasa kustom baru
* `DELETE /api/phrases/:languageId` - Hapus frasa untuk suatu bahasa
* `GET    /api/phrase-categories` - Ambil daftar kategori frasa kustom
* `POST   /api/phrase-categories` - Tambah satu kategori frasa kustom baru
* `DELETE /api/phrase-categories/:name` - Hapus kategori frasa kustom

### 6. Pengaturan & Preferensi (`/api/preferences`) `[Wajib Token]`
* `GET /api/preferences` - Ambil setelan kecepatan audio, toggle audio, dan tema
* `PUT /api/preferences` - Simpan perubahan preferensi pengguna
