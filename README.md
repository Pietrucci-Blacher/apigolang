# Api Golang

## Clone

```
git clone git@github.com:Pietrucci-Blacher/apigolang
```

## Init the database

run this command
```bash
docker compose up -d
```
and import this file `go_api.sql` in [adminer](http://localhost:8080)

## Dotenv

create a `.env` file and this example
```
HOST=localhost
PORT=8081
DB_DRIVER=mysql
DB_USER=root
DB_PASSWORD=azerty
DB_HOST=localhost
DB_PORT=3306
DB_NAME=go_api
JWT_SECRET=mysecret
```

## run Golang api

```
go run main.go
```

## Link

* [Api](http://localhost:8081/)
* [Swagger](http://localhost:8081/api/swagger/index.html)
