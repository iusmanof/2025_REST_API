
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