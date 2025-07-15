# URL Shortener Backend

REST API для сервиса сокращения ссылок. Написан на Go с использованием chi и стандартной библиотеки `database/sql`. Позволяет создавать короткие ссылки и перенаправлять по ним пользователей.

## 🚀 Функциональность

- Генерация короткой ссылки по оригинальному URL
- Редирект по короткому коду
- Валидация и логирование запросов

## 🛠️ Стек технологий

- [Go 1.22+](https://go.dev)
- [Chi](https://github.com/go-chi/chi) — маршруты
- [database/sql](https://pkg.go.dev/database/sql) — SQL-интерфейс
- [godotenv](https://github.com/joho/godotenv)

## 📁 Структура проекта

```
.
├── cmd/                # Точка входа (main.go)
├── internal/
│   └── shortener/      # Основной модуль приложения
├── pkg/                # Утилиты
├── .env                # Конфигурация окружения
└── go.mod / go.sum     # Зависимости
```

## 🔧 Быстрый старт

### 1. Подготовка базы данных

Создай базу и таблицу:

```sql
CREATE TABLE `shorten_links` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `shorten_url` varchar(255) NOT NULL,
    `original_url` varchar(255) NOT NULL,
    `redirect_count` int(11) NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

### 2. Настрой `.env`

```env
DB_NAME=
DB_USERNAME=
DB_PASSWORD=
DB_HOST=localhost:3306

APP_URL=localhost:8080
```

### 3. Сборка и запуск

```bash
go mod tidy
go run cmd/main.go
```

API будет доступен по адресу: `http://localhost:8080`

## 📬 Примеры API

### POST /api/shorten

```http
POST /api/shorten
Content-Type: application/json

{
  "url": "https://example.com"
}
```

Ответ:

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### GET /{id}

```http
GET /abc123
→ 302 Found → https://example.com
```
