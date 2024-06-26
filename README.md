# API сервис для дилерского центра

api сервис для дилерского центра

## Содержание

- [Установка и запуск](#установка-и-запуск)
- [Документация](#документация)

## Установка и запуск

### Зависимости

- [Go](https://go.dev/doc/install)(версия 1.19 и выше)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Docker-compose](https://docs.docker.com/compose/install/)

### Клонирование репозитория

Скопировать все три репозитория в одну директорию и перейти в директорию back

```sh
git clone https://github.com/ShawaDev/back
git clone https://github.com/ShawaDev/auth
git clone https://github.com/ShawaDev/front
cd back
```

### Запуск контейнеров

```sh
docker-compose up --build
```

### Доступ к контейнерам

API сервис доступен на порту localhost:8001
Сервис авторизации доступен на порту localhost:8002
Frontend работает на localhost:3000

## Документация

Swagger документация лежит в файле swagger/ swagger_docs.yaml

[swagger документация(API)](https://github.com/ShawaDev/back/blob/main/swagger/swagger_docs.yaml)  

[swagger документация(Auth)](https://github.com/ShawaDev/auth/blob/main/swagger/swagger_docs.yaml)

## Ссылки на другие репозитории

React.js приложение [react.js](https://github.com/ShawaDev/front)  

Сервис авторизации [auth](https://github.com/ShawaDev/auth)