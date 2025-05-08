# go-gin-crud

Простой CRUD на Golang. 
В основном, это повзрослевшая версия моего древнего проекта "DreamCatalogue",
но в формате веб-сервиса.

В проекте реализовано следующее:

- Управление SQLite при помощи GORM
- Добавление и удаление книг
- Интеграционные тесты

---

## Технологии

- Golang 1.24.2
- Gin Web Framework
- SQLite
- GORM
- Testing
---

## Запуск

### Инструкция
1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/diemensa/go-gin-crud
   cd go-gin-crud
2. Установить зависимости
   ```bash
   go mod download
3. Запустить тесты
   ```bash
   go test -v .\tests\
4. Запустить приложение
   ```bash
   go run main.go