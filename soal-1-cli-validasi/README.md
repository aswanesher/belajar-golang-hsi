# CLI Validasi
CLI Validasi adalah sebuah aplikasi command line interface (CLI) yang digunakan untuk memvalidasi data input dari pengguna. Aplikasi ini ditulis dalam bahasa pemrograman Go dan menggunakan package `github.com/urfave/cli/v2` untuk menangani argumen dan opsi CLI.
## Fitur
- Validasi input berupa angka, string, dan email.
- Menampilkan pesan kesalahan jika input tidak valid.
- Menampilkan pesan sukses jika input valid.
## Instalasi
Untuk menginstal aplikasi ini, pastikan Anda telah menginstal Go di sistem Anda. Kemudian, jalankan perintah berikut untuk mengunduh dan menginstal dependensi yang diperlukan:
```bash
go mod tidy
```
## Penggunaan
Untuk menjalankan aplikasi, gunakan perintah berikut:
```bash
go run main.go [opsi] [argumen]
```
### Opsi
- `--number <angka>`: Validasi input berupa angka.
- `--string <teks>`: Validasi input berupa string.
- `--email <alamat_email>`: Validasi input berupa alamat email.