# GoLibrary

## Перед запуском:

### Войти в консоль PostgreSQL
```bash
sudo -u postgres psql
```
### В консоли PostgreSQL выполнить:
```sql
ALTER USER postgres WITH PASSWORD 'новый_пароль';
```
### Выйти из консоли
```sql
\q
```

---
## Сам запуск:
###  Сборка и запуск контейнеров

```bash
docker-compose up -d
```

### Проверка, что контейнеры запущены

```bash
docker ps
```

### Проверка логов
```bash
docker-compose logs -f
```



## API
### Эндпоинты
GET /api/books - получить список всех книг
GET /api/books/:id - получить книгу по ID
POST /api/books - создать новую книгу
PUT /api/books/:id - обновить книгу
DELETE /api/books/:id - удалить книгу

### Создание новой книги
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Война и мир","author":"Лев Толстой","isbn":"1234567890"}'

### Получение списка книг
curl http://localhost:8080/api/books

