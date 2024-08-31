FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod download

COPY . .

RUN swag init -g ./cmd/api/main.go -o ./cmd/docs

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]
