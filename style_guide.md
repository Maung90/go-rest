````markdown
# 🧭 Golang Style Guide

> Panduan penulisan kode Go agar konsisten, mudah dibaca, dan sesuai dengan idiom Go.  
> Berdasarkan praktik terbaik komunitas Go dan rekomendasi dari [Effective Go](https://go.dev/doc/effective_go).

---

## 🧱 1. Package dan Folder
- Gunakan huruf kecil semua, tanpa underscore (`_`) atau huruf besar.
- Gunakan bentuk tunggal (`user`, bukan `users`).
- Nama folder mengikuti nama package.

✅ Contoh:
```go
package user
````

📁 Struktur:

```
internal/user/
pkg/auth/
```

---

## 🧩 2. Variable

* Gunakan **camelCase**.
* Hindari underscore.
* Gunakan nama singkat tapi jelas (`userID`, bukan `u` atau `user_identifier`).

✅ Contoh:

```go
userID := 42
createdAt := time.Now()
```

---

## 🧱 3. Struct

* Gunakan **PascalCase**.
* Nama berupa **noun (kata benda)**.
* Hindari awalan `Data` atau `Struct`.

✅ Contoh:

```go
type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}
```

---

## 🧩 4. Field Struct

* Gunakan **PascalCase** untuk field yang diekspor.
* Gunakan **camelCase** untuk field internal.
* Gunakan **akronim kapital penuh** (`UserID`, bukan `UserId`).

✅ Contoh:

```go
type DailyStory struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    StoryText string    `json:"story_text"`
    CreatedAt time.Time `json:"created_at"`
}
```

---

## ⚙️ 5. Function & Method

* Gunakan **camelCase** untuk private.
* Gunakan **PascalCase** untuk public/exported.
* Gunakan kata kerja aktif yang menjelaskan tindakan.

✅ Contoh:

```go
func NewRepository(db *sql.DB) Repository
func (r *Repository) FindByID(id int) (User, error)
```

---

## 🧩 6. Interface

* Nama interface mewakili perilaku, berakhiran “-er”.
* Hindari awalan `I`.

✅ Contoh:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type UserRepository interface {
    FindByID(id int) (User, error)
}
```

---

## 🔠 7. Constant

* Gunakan **PascalCase** jika diekspor, **camelCase** jika internal.
* Gunakan `const (...)` untuk grup constant.

✅ Contoh:

```go
const (
    MoodHappy   = "happy"
    MoodSad     = "sad"
    MoodNeutral = "neutral"
)
```

---

## 📁 8. File Name

* Gunakan huruf kecil semua dengan underscore.
* Hindari huruf besar.

✅ Contoh:

```
handler.go
repository.go
daily_story.go
auth_middleware.go
```

---

## ⚙️ 9. Receiver Naming

* Gunakan 1 huruf kecil yang mewakili struct-nya.

✅ Contoh:

```go
func (s *Service) Save(...) {...}
func (r *Repository) FindAll(...) {...}
```

---

## ⚠️ 10. Error Handling

* Gunakan awalan `Err` untuk error global.
* Gunakan pesan error yang jelas dan deskriptif.

✅ Contoh:

```go
var (
    ErrNotFound     = errors.New("record not found")
    ErrUnauthorized = errors.New("unauthorized access")
)
```

---

## 🔡 11. Acronym

* Gunakan **ALL CAPS** untuk singkatan.

✅ Contoh:

```go
UserID
HTTPServer
JSONData
URL
```

---

## 🧩 12. Comment & Documentation

* Gunakan **GoDoc style** — komentar dimulai dengan nama fungsi/struct.
* Komentar menjelaskan fungsi secara singkat.

✅ Contoh:

```go
// FindByID mengembalikan user berdasarkan ID.
func (r *Repository) FindByID(id int) (User, error) { ... }

// DailyStory merepresentasikan entitas cerita harian.
type DailyStory struct { ... }
```

---

## 🧩 13. Import & Package Order

Urutan import:

1. Package standar Go
2. Third-party package
3. Internal package

Pisahkan tiap kelompok dengan satu baris kosong.

✅ Contoh:

```go
import (
    "database/sql"
    "time"

    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"

    "myapp/internal/user"
)
```

---

## 🧹 14. Return Value & Error

* Kembalikan error terakhir.
* Gunakan `if err != nil` pattern.
* Hindari panic kecuali benar-benar fatal.

✅ Contoh:

```go
func (r *Repository) FindByID(id int) (User, error) {
    row := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)

    var u User
    if err := row.Scan(&u.ID, &u.Name); err != nil {
        return User{}, err
    }

    return u, nil
}
```

---

## ⚡ 15. Naming Summary

| Elemen       | Style                     | Contoh                      |
| ------------ | ------------------------- | --------------------------- |
| Package      | lowercase                 | `package user`              |
| File         | lowercase_with_underscore | `daily_story.go`            |
| Struct       | PascalCase                | `type DailyStory struct`    |
| Field Struct | PascalCase                | `UserID int`                |
| Variable     | camelCase                 | `userID := 10`              |
| Function     | camelCase / PascalCase    | `FindByID()`                |
| Interface    | PascalCase + “-er”        | `type Reader interface {}`  |
| Constant     | PascalCase                | `const MoodHappy = "happy"` |
| Receiver     | 1 huruf singkat           | `(s *Service)`              |
| Acronym      | ALLCAPS                   | `UserID`, `HTTPServer`      |

---

## 🧭 16. Tools Disarankan

Untuk memastikan gaya konsisten secara otomatis:

* `gofmt` → Format kode.
* `goimports` → Format dan urutkan import.
* `golangci-lint` → Jalankan linting otomatis.
* `revive` → Linter tambahan dengan aturan Go style.

---

## ✨ Penutup

> "Code is read more often than it is written."
>
> * Robert C. Martin

Jagalah konsistensi, gunakan nama yang jelas, dan biarkan kode berbicara sendiri.

---

