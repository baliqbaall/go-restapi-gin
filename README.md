# Nama Proyek Anda (misalnya: Go Gin REST API)

A simple RESTful API built with Go and the Gin web framework. This project demonstrates basic CRUD operations for products.

## Fitur

- Endpoint untuk membuat, membaca, memperbarui, dan menghapus produk.
- Menggunakan GORM untuk ORM.
- Struktur proyek MVC (Model-View-Controller, meskipun di Go lebih ke Model-Handler).

## Teknologi yang Digunakan

- [Go](https://golang.org/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/) (Go ORM)
- **MySQL Database**

## Persyaratan Sistem

- Go 1.x.x (Ganti dengan versi Go yang Anda gunakan)
- Database (Sebutkan database yang Anda gunakan, dan pastikan sudah terinstal/berjalan)

## Cara Menjalankan Proyek

1.  **Kloning Repositori:**

    ```bash
    git clone [https://github.com/your-username/go-restapi-gin.git](https://github.com/your-username/go-restapi-gin.git)
    cd go-restapi-gin
    ```

    _(Ganti `your-username/go-restapi-gin` dengan URL repositori Anda)_

2.  **Instal Dependensi Go:**

    ```bash
    go mod tidy
    ```

3.  **Konfigurasi Database:**

    - Buat file `.env` di direktori root proyek.
    - Tambahkan variabel lingkungan untuk koneksi database Anda (sesuaikan dengan setup Anda):
      ```
      DB_HOST=127.0.0.1
      DB_PORT=3306
      DB_USER=root
      DB_PASSWORD=
      DB_NAME=your_database_name
      ```
      _(Sesuaikan ini dengan kredensial database Anda. Pastikan database `your_database_name` sudah dibuat.)_

4.  **Jalankan Aplikasi:**
    ```bash
    go run main.go
    ```
    Aplikasi akan berjalan di `http://localhost:8080` (atau port yang Anda konfigurasi).

## Endpoint API

Berikut adalah daftar endpoint yang tersedia:

- `GET /api/products` - Mendapatkan semua produk
- `GET /api/product/:id` - Mendapatkan produk berdasarkan ID
  - Contoh: `GET http://localhost:8080/api/product/1`
- `POST /api/product` - Membuat produk baru (contoh body JSON di Postman/Insomnia)
  ```json
  {
    "name": "Product Name",
    "description": "Product Description"
  }
  ```
- `PUT /api/product/:id` - Memperbarui produk berdasarkan ID
  - Contoh: `PUT http://localhost:8080/api/product/1`
  - **Body JSON Contoh (sama seperti POST):**
    ```json
    {
      "name": "Updated Product Name",
      "description": "Updated Product Description"
    }
    ```
- `DELETE /api/product/` - Menghapus produk
  - Contoh: `DELETE http://localhost:8080/api/product`
    ```json
    {
      "id": 1
    }
    ```

## Kontribusi

Kontribusi sangat dihargai! Jika Anda memiliki ide atau ingin melaporkan bug, silakan buat _issue_ atau kirim _pull request_.
