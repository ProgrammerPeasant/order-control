# 📊 Приложение для управления сметами

##  О проекте

Данное приложение предназначено для автоматизации процесса создания, согласования и управления сметами в компаниях, предоставляющих услуги.  Оно упрощает взаимодействие между менеджерами и клиентами, снижает количество ошибок и ускоряет процесс согласования. 

###   Функциональные возможности

Приложение предоставляет следующие возможности:

* Регистрация компаний
* Добавление услуг
* Формирование детализированных смет
* Внесение изменений и комментариев в сметы
* Предоставление клиентам доступа к сметам через ссылки
* Экспорт данных в различные форматы 

###   Целевая аудитория

Программа предназначена для использования через веб-браузеры. 

* **Администраторы:** используют систему для настройки услуг и управления правами доступа. 
* **Менеджеры:** создают и редактируют сметы. 
* **Клиенты:** просматривают, комментируют и согласовывают документы. 

##  🚀 Запуск проекта

###   Требования к окружению

Для работы приложения необходимы:

* Любое клиентское устройство с операционной системой Windows 7 или выше, либо Ubuntu 22 или выше. 
* Сервер с операционной системой Ubuntu версии 22 или выше, 64-разрядным (x64) процессором, доступом в Интернет, 4 ГБ ОЗУ и 8 ГБ свободного места на внутреннем накопителе для программы и её зависимостей, а также 60 ГБ свободного места для хранения файлов клиентов.
* Установленный Docker

###   Установка

1.  Клонируйте репозиторий:

    ```bash
    git clone https://github.com/ProgrammerPeasant/order-control.git
    cd order-control
    ```
3.  Установите и запустите Docker (приложение на Windows или через sudo apt install docker-compose-plugin на Линукс)
4.  Создайте файл переменных окружения `.env` в директории `/backend` и заполните его необходимыми значениями:

    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=<пароль>
    DB_NAME=order_control
    JWT_SECRET=<секретный_ключ_для_jwt>
    ```

###   Запуск
    
Запуск программы осуществляется автоматически через Docker.

    docker-compose build
    docker-compos up

##  🛠️ Используемые технологии

* Golang версии 1.23.2 
* Docker (на базе образов golang:1.23-alpine и alpine:latest) 
* HTTP-сервер на порту 8080 
* Gin (веб-фреймворк) версии 1.10.0 
* JWT (библиотека авторизации) версии 3.2.0 
* GORM (ORM) версии 1.9.16 
* Swagger/Swag (документирование API) версий 1.16.4/1.6.0 
* Excelize (работа с Excel) версии 2.9.0 
* Godotenv (работа с переменными окружения) версии 1.5.1 
* PostgreSQL 
