
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

- [+] docker run test command
docker compose exec rest_api go test -cover -v ./...

go mod tidy (на остановленных контейнерах)
docker exec -it rest_api sh  # или bash
go test ./handlers


- [+] test POST genre
func TestCreateGenre(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()
	repo := Repository{DB: db}
	repo.SetupRoutes(app)

	name := "Horror"
	genre := models.Genre{Name: &name }
	body, _ := json.Marshal(genre)

	req := httptest.NewRequest("POST", "/api/genre", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
