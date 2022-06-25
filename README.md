<h2 align="center"><a href="https://friend-service-api.herokuapp.com" >REST API для платформы ДРУГ</a></h2>

Протестировать сервис можно по ссылке: https://friend-service-api.herokuapp.com
### База данных
Работает с PostgreSQL  
Настройки подключения хранятся в /configs/config.yml  
Пароль пользователя в системной переменной FRIEND_DB_PASS  

Для запроса GET /v1/user добавлена простая проверка на авторизацию пользователя, 
флаг авторизации хранится в БД в поле login_status
Причина: предотвращения утечки персональных данных без авторизации

### Запуск:  
docker build -t friend-api:latest -f Dockerfile .
docker run -d -p 8080:8080 -p 5432:5432 friend-api

### Деплой на Heroku
- зарегистрировать аккаунт Heroku
- Подключить репозиторий к проекту
- подключить PostgreSQL к проекту
- Настроить переменные среды, для доступа к БД
- Можно деплоить: Deploy branch


