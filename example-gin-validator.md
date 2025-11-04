
## 📘 Dasar Penggunaan

Di Gin, validasi bekerja seperti ini:

```go
type UserInput struct {
    Name  string `json:"name" binding:"required,min=3,max=50"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=18,lte=65"`
}
```

---

## ✅ Daftar Lengkap Validasi Umum (Gin + go-playground/validator)

| Validator     | Keterangan                                           | Contoh                           |
| ------------- | ---------------------------------------------------- | -------------------------------- |
| **required**  | Field wajib diisi                                    | `binding:"required"`             |
| **omitempty** | Abaikan validasi jika kosong                         | `binding:"omitempty,email"`      |
| **len**       | Panjang harus sama dengan nilai                      | `binding:"len=10"`               |
| **min**       | Nilai minimal (untuk string, slice, map, atau angka) | `binding:"min=3"`                |
| **max**       | Nilai maksimal                                       | `binding:"max=100"`              |
| **eq**        | Harus sama dengan nilai tertentu                     | `binding:"eq=admin"`             |
| **ne**        | Tidak boleh sama dengan nilai tertentu               | `binding:"ne=guest"`             |
| **oneof**     | Harus salah satu dari nilai tertentu                 | `binding:"oneof=red green blue"` |
| **gt**        | Lebih besar dari (angka, waktu, panjang)             | `binding:"gt=18"`                |
| **gte**       | Lebih besar atau sama dengan                         | `binding:"gte=0"`                |
| **lt**        | Lebih kecil dari                                     | `binding:"lt=10"`                |
| **lte**       | Lebih kecil atau sama dengan                         | `binding:"lte=100"`              |

---

## 🧩 Validasi Format Data Umum

| Validator                 | Keterangan                    | Contoh                     |
| ------------------------- | ----------------------------- | -------------------------- |
| **email**                 | Validasi format email         | `binding:"required,email"` |
| **url**                   | Validasi URL                  | `binding:"url"`            |
| **uri**                   | Validasi URI                  | `binding:"uri"`            |
| **ip**                    | Validasi alamat IP            | `binding:"ip"`             |
| **ipv4**                  | Hanya IPv4                    | `binding:"ipv4"`           |
| **ipv6**                  | Hanya IPv6                    | `binding:"ipv6"`           |
| **mac**                   | Alamat MAC                    | `binding:"mac"`            |
| **hostname**              | Nama host valid               | `binding:"hostname"`       |
| **alpha**                 | Hanya huruf                   | `binding:"alpha"`          |
| **alphanum**              | Huruf dan angka               | `binding:"alphanum"`       |
| **numeric**               | Hanya angka                   | `binding:"numeric"`        |
| **hexadecimal**           | Nilai hex                     | `binding:"hexadecimal"`    |
| **uuid**                  | Valid UUID (versi berapa pun) | `binding:"uuid"`           |
| **uuid3 / uuid4 / uuid5** | UUID versi tertentu           | `binding:"uuid4"`          |

---

## 🕒 Validasi Waktu

| Validator    | Keterangan                                  | Contoh                          |
| ------------ | ------------------------------------------- | ------------------------------- |
| **datetime** | Validasi format waktu (layout `2006-01-02`) | `binding:"datetime=2006-01-02"` |
| **ltfield**  | Lebih kecil dari field lain                 | `binding:"ltfield=EndDate"`     |
| **ltefield** | Lebih kecil atau sama dengan field lain     | `binding:"ltefield=EndDate"`    |
| **gtfield**  | Lebih besar dari field lain                 | `binding:"gtfield=StartDate"`   |
| **gtefield** | Lebih besar atau sama dengan field lain     | `binding:"gtefield=StartDate"`  |

Contoh:

```go
type SleepInput struct {
    SleepStart time.Time `json:"sleep_start" binding:"required"`
    SleepEnd   time.Time `json:"sleep_end" binding:"required,gtfield=SleepStart"`
}
```

---

## 🧮 Validasi Angka

| Validator                                   | Keterangan                    | Contoh                        |
| ------------------------------------------- | ----------------------------- | ----------------------------- |
| **number**                                  | Memastikan field berisi angka | `binding:"number"`            |
| **gtefield / gtfield / ltefield / ltfield** | Bandingkan antar field        | `binding:"gtefield=MinValue"` |

---

## 🧰 Validasi Collection (array, slice, map)

| Validator  | Keterangan                                | Contoh                                   |
| ---------- | ----------------------------------------- | ---------------------------------------- |
| **dive**   | Validasi setiap elemen di dalam slice/map | `binding:"required,dive,required,min=3"` |
| **unique** | Semua elemen harus unik                   | `binding:"unique"`                       |

Contoh:

```go
type TagsInput struct {
    Tags []string `json:"tags" binding:"required,dive,min=2"`
}
```

---

## 💼 Validasi Khusus Tipe Data

| Validator     | Keterangan                                 | Contoh                          |
| ------------- | ------------------------------------------ | ------------------------------- |
| **boolean**   | Hanya `true` / `false`                     | `binding:"boolean"`             |
| **isdefault** | Harus memiliki nilai default (nol value)   | `binding:"isdefault"`           |
| **contains**  | String harus mengandung substring tertentu | `binding:"contains=@gmail.com"` |
| **excludes**  | Tidak boleh mengandung substring tertentu  | `binding:"excludes=@"`          |

---

## 💬 Contoh Lengkap Gabungan

```go
type RegisterInput struct {
    Username string `json:"username" binding:"required,alphanum,min=3,max=20"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
    Age      int    `json:"age" binding:"required,gte=18,lte=65"`
    Website  string `json:"website" binding:"omitempty,url"`
    Mood     string `json:"mood" binding:"required,min=3"`
}
```

---

## ⚙️ Bonus: Custom Validator

Kamu juga bisa bikin validator sendiri, contoh:

```go
import "github.com/go-playground/validator/v10"

var validate = validator.New()

func init() {
    validate.RegisterValidation("is-even", func(fl validator.FieldLevel) bool {
        value := fl.Field().Int()
        return value%2 == 0
    })
}
```

Pemakaiannya:

```go
type NumberInput struct {
    Value int `json:"value" binding:"required,is-even"`
}
```
