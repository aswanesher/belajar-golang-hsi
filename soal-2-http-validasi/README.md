# HTTP Validasi dengan Golang
- HTTP Validasi adalah sebuah aplikasi web yang digunakan untuk memvalidasi data input dari pengguna melalui HTTP request. Aplikasi ini ditulis dalam bahasa pemrograman Go dan menggunakan package `github.com/gin-gonic/gin` untuk menangani routing dan middleware.
## Fitur
- Validasi input berupa angka, string, dan email.
- Menampilkan pesan kesalahan jika input tidak valid.
- Menampilkan pesan sukses jika input valid.
- Mendukung format JSON untuk input dan output.
- Mendukung pengujian dengan Postman atau alat serupa.
- Penanganan error yang baik dengan kode status HTTP yang sesuai.
## Instalasi
Untuk menginstal aplikasi ini, pastikan Anda telah menginstal Go di sistem Anda. Kemudian, jalankan perintah berikut untuk mengunduh dan menginstal dependensi yang diperlukan:
```bash
go mod tidy
```
## Penggunaan
Untuk menjalankan aplikasi, gunakan perintah berikut:
```bash
go run main.go
```
### Endpoint
- `GET /validate?email=&age=`: Validasi input berupa angka.
  - Body: JSON dengan format `{"number": <angka>}`
  - Response:
    - 200 OK: Jika input valid, akan mengembalikan pesan sukses.
    - 400 Bad Request: Jika input tidak valid, akan mengembalikan pesan kesalahan.
