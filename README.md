# 2025_REST_API

#golang#postgres#fier#gorm#test

- [+] docker rmi $(docker images -f "dangling=true" -q)
- [+] docker images
- [+] docker volume ls 
- [+] docker volume rm ...
- [+] docker ps -a  
- [+] docekr rm ...
- [+] пересобрать с очисткой volume
        docker-compose down -v  # Удалить контейнеры и volume
        docker-compose up --build  # Пересобрать образы и запустить

- [+] go mod init mod/2025_REST_API

- [+] postgres env:
POSTGRES_IN_PORT=5432
POSTGRES_OUT_PORT=5432
POSTGRES_HOST=pgdb_container
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=pass
POSTGRES_DB=shop_db
POSTGRES_SSLMODE=disable

- [+] postgres docker compose:

services:
  pgdb_container:
    container_name: pgdb_container
    image: postgres:16
    environment:
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_DB: ${POSTGRES_DB} 
    ports:
      - ${POSTGRES_IN_PORT}:${POSTGRES_OUT_PORT}  
    volumes:
      - pgdata:/var/lib/postgres/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres -d shop_db"]
      interval: 10s
      retries: 3
      start_period: 30s
      timeout: 10s

volumes:
  pgdata: {}

- [+] docker exec -it pg_container psql -U postgres -d shop_db
- [+] \dt

- [+] api env
API_IN_PORT=8000
API_OUT_PORT=8080
- [+] api docker compose 
  rest_api:
    image: iusmanof/golang_restapi:v1.0
    container_name: rest_api
    build: "."
    ports:
      - "${API_IN_PORT}:${API_OUT_PORT}"
    environment:
      - POSTGRES_USER=${POSTGRES_USER}
      - POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
      - POSTGRES_DB=${POSTGRES_DB}
      - POSTGRES_HOST=${POSTGRES_HOST}
      - POSTGRES_PORT=${POSTGRES_PORT}
      - POSTGRES_SSLMODE=${POSTGRES_SSLMODE}
    depends_on:
      pg_container:
        condition: service_healthy

- [+] bind mounts
  docker compose up pg_container -d
  docker compose up rest_api --build
  volumes:
    - .:/app  # Привязка текущей папки с кодом внутрь контейнера

- [] docker run test command
docker compose exec rest_api go test -v ./...

go mod tidy (на остановленных контейнерах)
docker exec -it rest_api sh  # или bash
go test ./handlers

- [] test GET
- [] test POST
- [] test PUT
- [] test DELETE
- [] test AUTH
- [] AUTH

 
ERRORS:

- [+] ERROR
  [error] failed to initialize database, got error cannot parse `host=localhost ... sslmode=false`: failed to configure TLS (sslmode is invalid)
  DB_SSLMODE=disable  

- [+] ERROR
  Error: socket hang up
%%%%%%%%%%%
  API_IN_PORT=8000
  API_OUT_PORT=8080   Fiber v2.52.6 http://127.0.0.1:8080 (bound on host 0.0.0.0 and port 8080)
%%%%%%%%%%%

- [+] ERROR
Conflict. The container name "/pgdb_container" is already in use
%%%%%%%%%%%
pg_container:
  container_name: pg_container
%%%%%%%%%%%

- [] ERROR
 => CACHED [rest_api 2/5] WORKDIR /
 => [rest_api 3/5] COPY . .
 => ERROR [rest_api 4/5] RUN go get -d   

%%%%%%%%%%%
go.mod:
module github.com/2025_REST_API

go mod tidy

main.go:
import (
  ...
	"github.com/2025_REST_API/models"
	"github.com/2025_REST_API/storage"
  ...
)

go mod tidy
%%%%%%%%%%%

- [+] ERROR 
rest_api exited with code 0
POSTMAN:Error: connect ECONNREFUSED 127.0.0.1:8000

%%%%%%%%%%%
Rebuild container:

docker-compose down -v
docker-compose build --no-cache
docker-compose up
%%%%%%%%%%%

- [] test FAIL
handlers/genre_test.go:18:1: expected declaration, found Host
FAIL    github.com/2025_REST_API/handlers [setup failed]
?       github.com/2025_REST_API        [no test files]
?       github.com/2025_REST_API/models [no test files]
?       github.com/2025_REST_API/storage        [no test files]
FAIL