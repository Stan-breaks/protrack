# 🚀 NimbleStack

_A Modern Go + Templ + Tailwind CSS Starter Template with HTMX, Alpine.js & SQLC_

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)  
![Templ](https://img.shields.io/badge/Templ-0.2+-blue)  
![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-3.3+-06B6D4?logo=tailwind-css)  
![HTMX](https://img.shields.io/badge/HTMX-1.9+-red)  
![Alpine.js](https://img.shields.io/badge/Alpine.js-3.13+-8BC0D0)
![sqlc](https://img.shields.io/badge/sqlc-1.25+-brightgreen)  
![SQLite](https://img.shields.io/badge/SQLite-3+-003B57?logo=sqlite)

**NimbleStack** is a lightning-fast, full-stack starter template designed for developers who want to build modern web apps with minimal boilerplate. Features **SQLite** + **SQLC** for embedded database magic! ✨

---

## 🌟 Features

- **Go Backend**: Blazing-fast API and server logic with Go
- **SQLite + SQLC**: Type-safe database access with single-file storage
- **Templ Templates**: Clean, type-safe HTML templating
- **Tailwind CSS**: JIT-compiled utility-first CSS
- **HTMX + Alpine.js**: Dynamic UI without JavaScript fatigue

---

## 🛠️ Why NimbleStack?

- **Zero Deployment Hassle**: Single binary with embedded SQLite database
- **Full-Stack Type Safety**: SQLC → Go → Templ pipeline
- **Local Development Bliss**: No database servers required
- **HTMX+Alpine.js Simplicity**: Partial reloads without React complexity
- **Modern UI/UX**: Tailwind + Alpine.js for polished, responsive interfaces.

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+ & pnpm
- Tailwind
- SQLC

### Installation

1. Clone the repo:

   ```bash
   git clone https://github.com/Stan-breaks/nimblestack.git
   cd nimblestack
   ```

2. Install dependencies:

   ```bash
   pnpm install
   go mod tidy
   go get modernc.org/sqlite  # SQLite driver
   ```

3. Generate code:

   ```bash
   templ generate ./views/
   sqlc generate
   ```

4. Start server:
   ```bash
   go run main.go
   ```

---

## 📂 Project Structure

```
nimblestack/
├── database/         # Generated Go models
├── sqlc/             # sqlc schema and queries
├── public/           # Static assets
├── views/            # Templ components
├── handlers/         # HTTP handlers
├── sqlc.yaml         # SQLC configuration
└── main.go           # Server entry
```

---

## 🔌 Database Workflow (SQLite + SQLC)

### 1. Create Migration

`db/migrations/001_users.up.sql`:

```sql
CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### 2. Write Queries

`db/query/users.sql`:

```sql
-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES (?, ?)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ?;
```

### 3. Generate Code

```bash
sqlc generate
```

### 4. Use in Handler

`handlers/users.go`:

```go
   func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    // Get DB connection
    db, _ := sql.Open("sqlite3", ".sqlite.db")
    defer db.Close()

    queries := db.NewQueries()

    // Type-safe database operation
    user, err := queries.CreateUser(r.Context(),
        db.CreateUserParams{
            Name: r.FormValue("name"),
            Email: r.FormValue("email"),
        })

    if err != nil {
        http.Error(w, "Database error", 500)
        return
    }

    // Render response with Templ
    components.UserCard(user).Render(r.Context(), w)
   }
```

---

## 📈 Roadmap

- [ ] Add SQLite migration tool
- [ ] HTMX CRUD example with optimistic UI
- [ ] SQLite connection pool benchmarks
- [ ] ARM64 build support

---

## 📚 Learning Resources

- [SQLC SQLite Guide](https://docs.sqlc.dev/en/latest/howto/sqlite.html)
- [Modern SQLite Driver Docs](https://pkg.go.dev/modernc.org/sqlite)
- [HTMX Patterns](https://htmx.org/examples/)

---

## License

MIT © [Stan-breaks] | Made with ❤️ for fast web apps
