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

### Найти PID процесса, использующего порт 5432
```bash
sudo lsof -i :5432
```
### Убить процесс
```bash
sudo kill -9 <PID>
```
### Попробовать запустить контейнеры снова
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

### Cваггер
```bash
curl http://localhost:8080/swagger/index.html
```
# GOBrokers


# Практическая работа №4

1) Установлен Nginx
  - Был создан отдельный сервис в docker-compose.yml для nginx
2) Установлен балансировщик нагрузки
  - В docker-compose.yml создано 3 одинаковых сервера Go, в конфигурации nginx указана балансровка между ними в зависимоти то кол-ва соединений к серверам. 
3) Установлен самоподписывающийся SSL сертификат
  - При переходе на домен localhost получаем предупереждение
4) Меры безопасности:
  - Установлено ограничение для подключений
  - Защита от брутфорса
  - Защита от крупных запросов 
  - Защита от крупных запросов 
  - Ограничение доступа по IP (Если попробовать перейти по пути admin с любого ip, кроме 192.168.1.100, то в ответ получим ошибку)
  - Скрыты версии сервеного ПО
5) Логирование и мониторинг
  - Создан отдельный сервис в docker-compose.yml для fail2ban, установлена свзяь с nginx, настроена конфигурация согласно примеру


# Практическая работа №5

1) Конфигурация
- Релализована конфигурация при помощи .env и yaml. 
- Для выбора типа конфигурационного файла используются флаги --config (путь до файла конфигурации) и --config-type (тип файла конфигурации). 
- Если не указан тип или путь, то стандартно используется .env. Пример запуска:
```bash
go run main.go --config-type env --config ./.env 
```

```bash
go run main.go --config-type yaml --config ./config.yaml 
```
Примерный файл для yaml и .env:
```yaml
    env: 
    app_port: 
    db_host: 
    db_user: 
    db_password: 
    db_name: 
    db_port: 
```

```env
  ENV=
  APP_ENV=
  APP_PORT=
  DB_HOST=
  DB_USER=
  DB_PASSWORD=
  DB_NAME=
  DB_PORT=
```

2) Логирование
- Реализовано логирование при помощи logrus. 
- Оно ведется и в консоль и в файл logs/app.log. 
- Настроена связь с docker-compose при помощи тома. 
- Для добавление логгера в http-клиент gin используется middleware. 
- Также логирование зависит от переменной Env в конфигуарционном файле. 
- При значении dev логирование начинается с уровня debug, при значении prod с уровня info.

# Практическая работа №6
1) Создать базовую модель пользователя
 - У нас уже была создана модель пользователя, дополним ее полями Username и Password:
```go 
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Email     string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Username  string         `gorm:"size:255;uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	Lendings  []Lending      `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```
- Добавили методы для создания пользователя, получения пользователя по username, авторизации по паролю

2) Реализовать хеширование пароля
 - Для метода хэширования и проверки пароля добавили ограничения по составу и длине пароля

3) Генерация и проверка JWT
- Реализовали кастомные claims для токена
- В методе для проверки токена реализовали проверку на валидность (соответствие секрету, метод подписи, срок действия)

4) Реализация эндпоинтов для аутентификации
 - Добавили эндпоинт для авторизации

5) Middleware для проверки JWT
- Усовершенствовали middleware (проверка на формат токена, передача данных о токене и пользователе в контексте)

6) Добавить защищенные маршруты 
- Добавили защищенный маршрут /api/protected.

## Тестирование:

<image src="./images/practice 6/passcheck1.png" alt="Описание изображения">

[Ошибка при создании пользователя](/images/passcheck1.png)

<image src="./images/practice 6/createUser.png" alt="Описание изображения">

[Создание пользователя](/images/createUser.png)

<image src="./images/practice 6/login.png" alt="Описание изображения">

[Авторизация](/images/login.png)

<image src="./images/practice 6/prot2.png" alt="Описание изображения">

[Ошибка в токене](/images/prot2.png)


<image src="./images/practice 6/prot1.png" alt="Описание изображения">

[Доступ до закрытого эндпоинта](/images/prot1.png)