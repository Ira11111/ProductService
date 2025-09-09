# ProductService API

RESTful API для управления инвентарем, реализованный на Go с использованием фреймворка Gin. API предоставляет полный CRUD для продуктов, продавцов, складов и категорий.

## 🚀 Возможности

- **Управление продуктами**: Создание, чтение, обновление, удаление продуктов
- **Управление категориями**: Иерархия категорий товаров
- **Управление продавцами**: Регистрация и управление продавцами
- **Управление складами**: Учет товаров на различных складах
- **JWT аутентификация**: Защищенные endpoints с bearer токенами
- **Загрузка файлов**: Поддержка загрузки фотографий товаров и логотипов

## 📋 API Endpoints

### Продукты
- `GET /products` - Список всех продуктов (с пагинацией)
- `POST /products` - Создать новый продукт
- `GET /products/{id}` - Получить продукт по ID
- `PUT /products/{id}` - Обновить продукт
- `DELETE /products/{id}` - Удалить продукт

### Категории
- `GET /categories` - Получить дерево категорий
- `POST /categories` - Создать новую категорию
- `DELETE /categories/{id}` - Удалить категорию
- `GET /categories/{id}/products` - Продукты в категории

### Продавцы
- `POST /sellers` - Зарегистрировать нового продавца
- `GET /sellers/{id}` - Найти продавца по ID
- `PUT /sellers/{id}` - Редактировать продавца
- `DELETE /sellers/{id}` - Удалить продавца

### Склады
- `GET /warehouses` - Список всех складов
- `POST /warehouses` - Создать новый склад
- `GET /warehouses/{id}` - Получить склад по ID
- `GET /warehouses/{id}/products` - Продукты на складе

## 🛠 Технологии

- **Go 1.21+** - Язык программирования
- **Gin** - Веб-фреймворк
- **JWT** - Аутентификация
- **PostgreSQL** - База данных
- **Docker** - Контейнеризация

## 📦 Установка и запуск

### Предварительные требования

- Go 1.21 или выше
- PostgreSQL 12+
- Docker (опционально)

## ⚙️ Конфигурация

Основные переменные окружения:

```env
DB_PASS
CONFIG_PATH
GOOSE_DRIVER
GOOSE_DBSTRING
GOOSE_MIGRATION_DIR=./migrations
```

```yaml
# Окружение (local/dev/prod)
env: 

# Настройки базы данных
db:
  host:       # Хост БД
  port:       # Порт БД
  user:       # Пользователь БД
  database:   # Имя БД
  ssl_mode:   # Режим SSL

http:
  port: # Порт сервера
  handler: # тип обработчика
  read_timeout: # время на чтение
  write_timeout: # время на запись
  idle_timeout: # время простоя соединения
```

## 🔐 Аутентификация

API использует JWT токены для аутентификации. Для доступа к защищенным endpoints необходимо:

1. Получить JWT токен через вашу систему аутентификации
2. Добавить заголовок в запросы:
```
Authorization: Bearer <your-jwt-token>
```

## 📝 Примеры запросов

### Создание продукта

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <jwt-token>" \
  -d '{
    "name": "Smartphone X10",
    "description": "Flagship smartphone with AI camera",
    "price": 999.99,
    "categories_id": [1, 2],
    "photos": []
  }'
```

### Получение списка продуктов

```bash
curl -X GET "http://localhost:8080/products?limit=10&offset=0"
```

### Получение продуктов в категории

```bash
curl -X GET "http://localhost:8080/categories/1/products?limit=10&offset=0"
```

## 🗄 Структура проекта

```
├── cmd/
│   ├── products/          # Основное приложение
├── internal/
│   ├── app/             # инициализация приложения
│   ├── handlers/        # HTTP обработчики
│   ├── services/        # Бизнес-логика
│   └──storage/         # Работа с базой данных
├── migrations/         # SQL миграции
└── configs/           # Конфигурационные файлы
```

middleware для аутентификации, модели и интерфейсы обработчиков находятся в проекте https://github.com/Ira11111/protos
