# Tugas Pertemuan 4: Worker Assignment Perhitungan Nilai Mahasiswa

Aplikasi ini adalah sebuah simulasi sistem untuk memberikan tugas kepada mahasiswa dan menilainya secara bersamaan (concurrently) menggunakan Goroutine, Channel, dan WaitGroup di Go.

## Deskripsi

Proyek ini dibuat untuk mendemonstrasikan konsep konkurensi dalam Go. Aplikasi akan melakukan hal-hal berikut:
1.  Terhubung ke database PostgreSQL.
2.  Menghapus data tugas dan hasil sebelumnya untuk memastikan setiap eksekusi dimulai dari keadaan bersih.
3.  Melakukan seeding data mahasiswa jika tabel mahasiswa masih kosong.
4.  Menggunakan **goroutine** untuk memberikan tugas kepada setiap mahasiswa yang ada di database.
5.  Menggunakan **goroutine** lain untuk menilai tugas yang telah diberikan.
6.  Menggunakan **channel** untuk komunikasi antara goroutine pemberi tugas dan goroutine penilai.
7.  Menggunakan **WaitGroup** untuk sinkronisasi dan memastikan semua proses selesai sebelum program berakhir.
8.  Menyimpan semua tugas dan hasil penilaian ke dalam database.
9.  Menampilkan hasil akhir di konsol.

## Fitur

-   **Koneksi Database**: Menggunakan GORM untuk terhubung ke PostgreSQL.
-   **Migrasi Database**: Skema database untuk `mahasiswa`, `tugas`, dan `hasil` dapat dibuat secara otomatis.
-   **Seeder Data**: Mengisi data awal untuk tabel `mahasiswa`.
-   **Proses Konkuren**: Implementasi `worker` untuk proses pemberian tugas dan penilaian secara bersamaan.
-   **Manajemen State**: Membersihkan data lama pada setiap eksekusi untuk menjaga konsistensi.

## Struktur Proyek

```
.
├── cmd
│   └── main.go         # Entry point utama aplikasi
├── config
│   └── db.go           # Konfigurasi koneksi database, migrasi, dan seeder
├── go.mod
├── go.sum
├── main.go             # Entry point untuk migrasi dan seeder
├── models
│   ├── hasil.go        # Model GORM untuk tabel hasil
│   ├── mahasiswa.go    # Model GORM untuk tabel mahasiswa
│   └── tugas.go        # Model GORM untuk tabel tugas
├── worker
│   ├── assignment_worker.go # Worker untuk memberikan tugas ke mahasiswa
│   └── grading_worker.go    # Worker untuk menilai tugas
└── .env                # File untuk menyimpan kredensial database (tidak termasuk di repo)
```

## Cara Menjalankan

1.  **Prasyarat**
    -   Pastikan Anda sudah menginstall **Go**.
    -   Pastikan Anda memiliki **PostgreSQL** yang sedang berjalan.

2.  **Clone Repositori**
    ```bash
    git clone <url-repositori-anda>
    cd tugas-pertemuan-4
    ```

3.  **Install Dependensi**
    ```bash
    go mod tidy
    ```

4.  **Konfigurasi Lingkungan**
    -   Buat file bernama `.env` di root direktori proyek.
    -   Isi file tersebut dengan informasi koneksi database Anda. Contoh:
        ```env
        DB_HOST=localhost
        DB_USER=postgres
        DB_PASS=password_anda
        DB_NAME=nama_database_anda
        DB_PORT=5432
        ```

5.  **Jalankan Migrasi**
    Buat tabel-tabel yang diperlukan di database Anda dengan menjalankan perintah:
    ```bash
    go run main.go migrate
    ```

6.  **Jalankan Seeder (Opsional)**
    Jika Anda ingin mengisi tabel `mahasiswas` dengan data awal, jalankan:
    ```bash
    go run main.go seed
    ```

7.  **Jalankan Aplikasi Utama**
    Untuk memulai proses pemberian tugas dan penilaian, jalankan:
    ```bash
    go run cmd/main.go
    ```
    Aplikasi akan mencetak proses pemberian tugas dan hasil penilaian di konsol.
