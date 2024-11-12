# Cart Services API

API ini adalah aplikasi manajemen pengguna menggunakan Golang dan PostgreSQL.

## Persiapan Sebelum Memulai

1. **Pastikan Database Sudah Tersedia**
   - Pastikan Anda sudah membuat database PostgreSQL dengan nama **`db_order`**.

2. **Konfigurasi Aplikasi**
   - Buka file `config.yaml` pada aplikasi ini.
   - Sesuaikan pengaturan koneksi database dan konfigurasi lainnya sesuai dengan lingkungan Anda.

## Langkah-Langkah Menjalankan Aplikasi

### 1. Menjalankan Migrations
   - Migration digunakan untuk mengatur struktur database yang diperlukan oleh aplikasi ini.
   - Masuk ke direktori `migrations/` dengan perintah berikut:
     ```bash
     cd migrations/
     ```
   - Jalankan perintah berikut untuk menjalankan migration:
     ```bash
     go run migration.go ./sql "host=localhost port=5432 user=root dbname=db_order password=password sslmode=disable" up
     ```
   - Pastikan detail koneksi (seperti host, port, user, dan password) sesuai dengan konfigurasi database PostgreSQL Anda.

### 2. Menjalankan Aplikasi
   - Setelah konfigurasi selesai, Anda dapat menjalankan aplikasi dengan perintah:
     ```bash
     go run .
     ```
   - Aplikasi sekarang akan berjalan dan terhubung ke database `db_order`.

## Teknologi yang Digunakan

- **Golang**: Backend aplikasi utama.
- **PostgreSQL**: Database untuk menyimpan data pengguna.

## AFter created proto file:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## Generated
protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative proto/cart/cart.proto
protoc --proto_path=proto --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/cart/cart.proto