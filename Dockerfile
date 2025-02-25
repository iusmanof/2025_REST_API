FROM golang:1.23.6-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 8000

CMD ["go", "run", "."]

