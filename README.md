# 🚀 Go API Auth

Mini API Auth berbasis Go dengan fitur:

* JWT


---

# 📁 Struktur Project

```
auth/
 ├─ go.mod
 ├─ main.go
 ├─ jwt.go
 ├─ repository.go
 ├─ entity.go
 ├─ users.sql
 └─ README.md
```

---

# 🧰 Prasyarat

* Go >= 1.22
* PostgreSQL >= 12
* Git

Cek versi:

```bash
go version
psql --version
```

---

# 🗄️ Database Schema

Jalankan SQL berikut di PostgreSQL:

- users.sql

---

# 🔐 Environment Variables

Auth membaca koneksi database dari:

```env
DATABASE_URL=postgres://user:password@host:5432/dbname?sslmode=disable
JWT_SECRET=supersecret
```

---

# 📦 Install Dependency

Masuk folder project:

```bash
cd auth
```

Download dependency:

```bash
go mod tidy
```

---

# ▶️ Menjalankan Aplikasi (Development)

```bash
go run .
```

Output:

```
Gateway running on :8080
```

---

# 🏗️ Build Binary

```bash
go build -o build/auth
```

Jalankan:

```bash
./build/auth
```

---

# 🌍 Contoh Request

Health endpoint:

```
GET http://localhost:8080/health
```

Login endpoint:

```bash
curl -X POST http://localhost:8080/api/login -H 'Content-Type: application/json' -d '{"email":"admin@admin.com","password":"admin"}'
```

Response:

```json
{
    "code": 200,
    "message": "Login Successful",
    "data": {
        "token": "token"
    }
}
```

---

# 📝 Lisensi

MIT / Bebas digunakan untuk keperluan pribadi dan edukasi.

---