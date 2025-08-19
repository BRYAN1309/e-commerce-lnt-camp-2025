# Backend Management System

Sebuah sistem backend untuk manajemen produk dan karyawan dengan autentikasi **JWT**, dibangun menggunakan **Go** dan **Gin framework**.

---
[Link Postman = https://api.postman.com/collections/45400437-7d2ac694-51c9-4518-9bf1-1450a42cb07e?access_key=
](https://api.postman.com/collections/45400437-7d2ac694-51c9-4518-9bf1-1450a42cb07e?access_key=PMAT-01K31HVJJ7JCY381210MPHHW4G)
## 📋 Fitur

- **Autentikasi & Otorisasi**
  - Register & Login dengan JWT
- **Manajemen Produk**
  - CRUD Produk
  - Upload gambar
  - Ekspor data ke Excel
- **Manajemen Karyawan**
  - CRUD karyawan
  - Status aktif / tidak aktif
- **Dashboard Admin**
  - Statistik & ringkasan data
- **Landing Page API**
  - Produk terbaru
  - Produk tersedia
- **Ekspor Data**
  - Ekspor ke format Excel

---

## 🛠 Teknologi

- **Bahasa:** Go (Golang)  
- **Framework:** Gin  
- **Database:** PostgreSQL / MySQL (via GORM)  
- **Autentikasi:** JWT (JSON Web Token)  
- **File Processing:** Excelize (ekspor Excel)  
- **Keamanan:** bcrypt (hashing password)  

---

## 📁 Struktur Proyek

backend/
├── configs/ # Konfigurasi database
├── handlers/ # Handler untuk route API
├── models/ # Struct model data
├── repository/ # Layer akses database
├── routes/ # Definisi route API
└── uploads/ # Penyimpanan file upload

---

## 🔌 API Endpoints

### Autentikasi
- `POST /admin/register` – Registrasi user baru  
- `POST /admin/login` – Login user  

### Produk (Admin)
- `POST /admin/products` – Buat produk baru  
- `GET /admin/products` – Dapatkan semua produk  
- `PUT /admin/products/:id` – Update produk  
- `GET /admin/products/export` – Ekspor produk ke Excel  

### Karyawan (Admin)
- `POST /admin/users` – Buat karyawan baru  
- `GET /admin/users` – Dapatkan semua karyawan  
- `PUT /admin/users/:id` – Update karyawan  

### Dashboard
- `GET /admin/dashboard` – Dapatkan statistik dashboard  

### Landing Page
- `GET /products/latest` – Dapatkan produk terbaru  
- `GET /products/available` – Dapatkan produk yang tersedia  

---

## 🚀 Instalasi & Menjalankan

### Prasyarat
- Go 1.16+  
- Database (PostgreSQL/MySQL)  

### Langkah-langkah

1. **Clone repository**
   ```bash
   git clone <repository-url>
   cd backend
Install dependencies
go mod download
Setup environment variables
export JWT_SECRET=your_jwt_secret_key
export DB_DSN=your_database_connection_string
Setup database

Buat database sesuai konfigurasi

Jalankan migrasi (jika ada)

Jalankan aplikasi

go run main.go


Akses aplikasi

Server akan berjalan di: http://localhost:8080
