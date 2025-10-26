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
    DB_NAME=habit_tracker
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
