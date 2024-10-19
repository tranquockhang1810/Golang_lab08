FROM golang:1.23.2-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o book-management

EXPOSE 8080

CMD ["./book-management"]