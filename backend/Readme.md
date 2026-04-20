# 📚 Bookstore API

A REST API for managing a bookstore built with **Go (Gin)**, **PostgreSQL**, and **Bob ORM**.

---

## 🚀 Setup & Run

### 1. Clone Repo

```bash
git clone https://github.com/your-username/book-store.git
cd book-store
```

---

## ⚙️ Environment Setup

Create a `.env` file in root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=bookstore
APP_PORT=8080
```

---

## 📦 Database Setup

### 1. Create Migration

```bash
./scripts/create_migration.sh <name>
```

Example:

```bash
./scripts/create_migration.sh create_users
```

---

### 2. Run Migrations

```bash
./scripts/migrate.sh
```

---

### 3. Generate Models (Bob ORM)

```bash
bobgen-psql
```

---

## 📄 Swagger Docs

Generate Swagger:

```bash
./scripts/swagger.sh
```

Access:

```
http://localhost:8080/swagger/index.html
```

---

## ▶️ Run Server

```bash
go run cmd/api/main.go
```

Server:

```
http://localhost:8080
```

---

## 🩺 Health Check

```http
GET /health
```

---

## 📁 Project Structure

```
internal/
├── config/
├── db/
├── routes/
cmd/
└── api/
scripts/
├── create_migration.sh
├── migrate.sh
└── swagger.sh
```

---

## 🛠️ Tech Stack

- Go (Golang)
- Gin
- PostgreSQL
- Bob ORM (`bobgen-psql`)
- Swagger (swaggo)

---

## ⚡ Quick Flow

```bash
# 1. create migration
./scripts/create_migration.sh create_books

# 2. write SQL in migration

# 3. run migrations
./scripts/migrate.sh

# 4. generate models
bobgen-psql

# 5. generate swagger
./scripts/swagger.sh

# 6. run server
go run cmd/api/main.go
```

---

## 👨‍💻 Author

Shailendra Pawar=>
