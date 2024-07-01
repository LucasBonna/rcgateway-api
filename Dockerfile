FROM golang:1.22.1-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

EXPOSE 8080

CMD [ "air", "-c", ".air.toml" ]