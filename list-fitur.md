
# 🧠 **Habit Tracker — Fitur & Endpoint API**

---

## 🧩 **1. Authentication & User Management**

Fitur dasar agar setiap user memiliki data habit masing-masing secara terpisah.

|  ✅  | Fitur            | Endpoint (contoh)              | Deskripsi                                            |
| :-: | ---------------- | ------------------------------ | ---------------------------------------------------- |
|  ✅  | Register user    | `POST /api/v1/register`        | Mendaftarkan akun baru dengan nama, email, password. |
|  ✅  | Login user       | `POST /api/v1/login`           | Login dan mendapatkan token JWT.                     |
|  ⬜  | Logout user      | `POST /api/v1/logout`          | Menghapus sesi/token aktif.                          |
|  ⬜  | Get user profile | `GET /api/v1/user`             | Menampilkan profil user yang sedang login.           |
|  ⬜  | Update profile   | `PUT /api/v1/user`             | Mengubah data profil user.                           |
|  ⬜  | Forgot Password  | `POST /api/v1/forgot-password` | Mengirim tautan atau token untuk reset password.     |

---

## 🗓️ **2. Habit Management**

Mengelola daftar kegiatan (habit) yang ingin dilakukan setiap hari.

|  ✅  | Fitur                        | Endpoint (contoh)                            | Deskripsi                                                       |
| :-: | ---------------------------- | -------------------------------------------- | --------------------------------------------------------------- |
|  ✅  | Buat habit baru              | `POST /api/v1/habits`                        | Membuat habit baru (contoh: “Olahraga pagi”, “Membaca buku”).   |
|  ✅  | Lihat semua habit            | `GET /api/v1/habits`                         | Menampilkan semua habit milik user.                             |
|  ✅  | Lihat detail habit           | `GET /api/v1/habits/{id}`                    | Menampilkan detail habit tertentu.                              |
|  ✅  | Update habit                 | `PUT /api/v1/habits/{id}`                    | Mengubah nama, deskripsi, atau kategori habit.                  |
|  ✅  | Hapus habit                  | `DELETE /api/v1/habits/{id}`                 | Menghapus habit tertentu.                                       |
|  ⬜  | Tandai habit sebagai selesai | `POST /api/v1/habits/{id}/complete`          | Menandai habit telah dilakukan pada hari tertentu.              |
|  ⬜  | Riwayat habit harian         | `GET /api/v1/habits/history?date=YYYY-MM-DD` | Melihat habit apa saja yang diselesaikan pada tanggal tertentu. |

---

## 😴 **3. Sleep Tracker**

Mencatat jam tidur dan durasi tidur pengguna.

|  ✅  | Fitur                | Endpoint (contoh)                   | Deskripsi                                        |
| :-: | -------------------- | ----------------------------------- | ------------------------------------------------ |
|  ⬜  | Tambah catatan tidur | `POST /api/v1/sleep`                | Menyimpan waktu mulai dan bangun tidur.          |
|  ⬜  | Lihat catatan tidur  | `GET /api/v1/sleep?date=YYYY-MM-DD` | Menampilkan durasi tidur di hari tertentu.       |
|  ⬜  | Update catatan tidur | `PUT /api/v1/sleep/{id}`            | Mengubah data tidur jika salah input.            |
|  ⬜  | Hapus catatan tidur  | `DELETE /api/v1/sleep/{id}`         | Menghapus catatan tidur tertentu.                |
|  ⬜  | Statistik tidur      | `GET /api/v1/sleep/statistics`      | Melihat rata-rata durasi tidur per minggu/bulan. |

---

## 📝 **4. Activity Log (Catatan Kegiatan Harian)**

Mencatat aktivitas sehari-hari seperti bekerja, berolahraga, membaca, dsb.

|  ✅  | Fitur                        | Endpoint (contoh)                        | Deskripsi                                                           |
| :-: | ---------------------------- | ---------------------------------------- | ------------------------------------------------------------------- |
|  ⬜  | Tambah catatan kegiatan      | `POST /api/v1/activities`                | Menyimpan aktivitas (contoh: “Bekerja 6 jam”, “Berjalan 30 menit”). |
|  ⬜  | Lihat semua catatan kegiatan | `GET /api/v1/activities?date=YYYY-MM-DD` | Menampilkan aktivitas per hari.                                     |
|  ⬜  | Update catatan kegiatan      | `PUT /api/v1/activities/{id}`            | Mengubah nama atau durasi aktivitas.                                |
|  ⬜  | Hapus catatan kegiatan       | `DELETE /api/v1/activities/{id}`         | Menghapus catatan tertentu.                                         |

---

## 💬 **5. Daily Reflection (Cerita & Mood Harian)**

Menulis cerita singkat dan suasana hati setiap hari.

|  ✅  | Fitur                    | Endpoint (contoh)                         | Deskripsi                                                             |
| :-: | ------------------------ | ----------------------------------------- | --------------------------------------------------------------------- |
|  ⬜  | Tambah cerita harian     | `POST /api/v1/daily-story`                | Menyimpan satu kalimat cerita (contoh: “Hari ini produktif banget!”). |
|  ⬜  | Lihat cerita per tanggal | `GET /api/v1/daily-story?date=YYYY-MM-DD` | Menampilkan cerita pada tanggal tertentu.                             |
|  ⬜  | Update cerita harian     | `PUT /api/v1/daily-story/{id}`            | Mengubah cerita jika perlu.                                           |
|  ⬜  | Hapus cerita harian      | `DELETE /api/v1/daily-story/{id}`         | Menghapus cerita.                                                     |
|  ⬜  | Statistik mood           | `GET /api/v1/daily-story/statistics`      | Menampilkan tren mood mingguan/bulanan dalam bentuk grafik.           |

---

## 📊 **6. Dashboard & Analytics**

Menyajikan ringkasan aktivitas dan kebiasaan pengguna.

|  ✅  | Fitur                      | Endpoint (contoh)                             | Deskripsi                                                                       |
| :-: | -------------------------- | --------------------------------------------- | ------------------------------------------------------------------------------- |
|  ⬜  | Ringkasan harian           | `GET /api/v1/dashboard/daily?date=YYYY-MM-DD` | Menampilkan habit yang diselesaikan, jam tidur, aktivitas, dan cerita hari itu. |
|  ⬜  | Statistik mingguan/bulanan | `GET /api/v1/dashboard/summary?range=week`    | Menampilkan total habit done, rata-rata tidur, jumlah kegiatan, dan tren mood.  |

---

## 🔔 **7. Notifications / Reminders**

Mengirim pengingat otomatis untuk menjalankan habit tertentu.

|  ✅  | Fitur                    | Endpoint (contoh)                                             | Deskripsi                                                |
| :-: | ------------------------ | ------------------------------------------------------------- | -------------------------------------------------------- |
|  ⬜  | Tambah pengingat habit   | `POST /api/v1/reminders`                                      | Menetapkan waktu pengingat (contoh: jam 07:00 olahraga). |
|  ⬜  | Lihat pengingat          | `GET /api/v1/reminders`                                       | Melihat semua pengingat aktif.                           |
|  ⬜  | Update / Hapus pengingat | `PUT /api/v1/reminders/{id}`, `DELETE /api/v1/reminders/{id}` | Mengelola status pengingat (aktif/nonaktif).             |

---

## 🧾 **Checklist Legend**

| Simbol | Status                                        |
| :----: | :-------------------------------------------- |
|    ✅   | Sudah diimplementasikan                       |
|    ⬜   | Belum diimplementasikan / dalam tahap rencana |

