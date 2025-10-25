# Go REST API - Habit Tracker Backend 🚀

Ini adalah proyek backend API yang dibangun menggunakan **Go (Golang)** untuk aplikasi **Habit Tracker**. Proyek ini dirancang dengan struktur **Clean Architecture** untuk memastikan kode yang bersih, mudah dikelola, dan teruji. Selain itu, proyek ini juga memanfaatkan fitur **Generics** (Go 1.18+) untuk mengurangi duplikasi kode pada *service layer*.

-----

## \#\# Arsitektur & Teknologi 🛠️

  * **Bahasa**: Go (Golang)
  * **Framework**: Gin Web Framework
  * **Database**: MySQL
  * **Arsitektur**: Clean Architecture
  * **Prinsip Desain**:
      * **Separation of Concerns**: Setiap lapisan (handler, service, repository) memiliki tanggung jawab yang jelas.
      * **Dependency Injection**: Ketergantungan antar lapisan disuntikkan dari file `main.go` untuk mempermudah *testing* dan fleksibilitas.
      * **DRY (Don't Repeat Yourself)**: Menggunakan *Generic Service* untuk logika bisnis CRUD yang berulang.

-----

## \#\# Fitur Habit Tracker 📋

Berikut adalah daftar fitur yang direncanakan untuk aplikasi Habit Tracker. Fitur yang sudah diimplementasikan di API ini ditandai dengan centang.

  * [x] **Manajemen User** (Dasar)
      * [x] Membuat User Baru
      * [x] Melihat Daftar User
      * [x] Melihat Detail User
      * [x] Mengubah Data User
      * [x] Menghapus User
      * [ ] Autentikasi (Login/Logout & Token JWT)
  * [x] **Manajemen Aktivitas/Habit**
      * [x] Membuat Aktivitas Baru
      * [x] Melihat Daftar Aktivitas
      * [x] Melihat Detail Aktivitas
      * [x] Mengubah Data Aktivitas
      * [x] Menghapus Aktivitas
  * [ ] **Pelacakan (Tracking)**
      * [ ] Menandai habit selesai untuk hari ini
      * [ ] Melihat riwayat penyelesaian habit
  * [ ] **Analitik & Progres**
      * [ ] Melihat *streak* (jumlah hari berturut-turut)
      * [ ] Kalkulasi tingkat keberhasilan per minggu/bulan
      * [ ] Dashboard visualisasi progres

-----

## \#\# Setup & Instalasi

1.  **Clone Repositori**

    ```bash
    git clone [URL_REPOSITORI_ANDA]
    cd go-rest
    ```

2.  **Konfigurasi Environment**
    Buat file bernama `.env` di root direktori proyek dan salin konten dari `.env.example`. Sesuaikan nilainya dengan konfigurasi database MySQL Anda.

    **.env.example**

    ```env
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=
    DB_NAME=crud_go
    ```

3.  **Install Dependencies**
    Pastikan semua *dependency* yang dibutuhkan sudah ter-install.

    ```bash
    go mod tidy
    ```

4.  **Jalankan Server**

    ```bash
    go run cmd/api/main.go
    ```

    Server akan berjalan di `http://localhost:8080`.

-----

## \#\# Daftar API Endpoint

Semua *endpoint* berada di bawah prefix `/api/v1`.

### \#\#\# User Endpoints

| Method | Path               | Deskripsi                 | Contoh Body (Payload)                                                              |
| :----- | :----------------- | :------------------------ | :--------------------------------------------------------------------------------- |
| `GET`  | `/users`           | Mengambil semua user      | -                                                                                  |
| `GET`  | `/users/{id}`      | Mengambil user by ID      | -                                                                                  |
| `POST` | `/users`           | Membuat user baru         | `{ "name": "Budi", "email": "budi@email.com", "password": "password123" }`         |
| `PUT`  | `/users/{id}`      | Memperbarui data user     | `{ "name": "Budi Santoso", "email": "budi.santoso@email.com" }`                    |
| `DELETE`| `/users/{id}`      | Menghapus user by ID      | -                                                                                 |

### \#\#\# Activity Endpoints

| Method | Path                 | Deskripsi                     | Contoh Body (Payload)                                        |
| :----- | :------------------- | :---------------------------- | :----------------------------------------------------------- |
| `GET`  | `/activities`        | Mengambil semua aktivitas     | -                                                            |
| `GET`  | `/activities/{id}`   | Mengambil aktivitas by ID     | -                                                            |
| `POST` | `/activities`        | Membuat aktivitas baru        | `{ "user_id": 1, "activity": "Belajar Golang 1 Jam" }`       |
| `PUT`  | `/activities/{id}`   | Memperbarui data aktivitas    | `{ "user_id": 1, "activity": "Olahraga Pagi 30 Menit" }`       |
| `DELETE`| `/activities/{id}`   | Menghapus aktivitas by ID     | -                                                            |