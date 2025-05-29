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

# Найти PID процесса, использующего порт 5432
```bash
sudo lsof -i :5432
```
# Убить процесс
```bash
sudo kill -9 <PID>
```
# Попробовать запустить контейнеры снова
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
- GET /api/books - получить список всех книг
- GET /api/books/:id - получить книгу по ID
- POST /api/books - создать новую книгу
- PUT /api/books/:id - обновить книгу
- DELETE /api/books/:id - удалить книгу

### Создание новой книги
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Война и мир","author":"Лев Толстой","isbn":"1234567890"}'
```
### Получение списка книг
```bash
curl http://localhost:8080/api/books
```

### Получение списка книг
```bash
curl http://localhost:8080/swagger/index.html
```
