FROM golang:latest

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . .

RUN go build -o /run-gin-api

EXPOSE 8080

CMD ["/run-gin-api"]
