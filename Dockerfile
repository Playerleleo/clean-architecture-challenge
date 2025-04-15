FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go clean -modcache
RUN go mod download
RUN go mod tidy

COPY . .

RUN go mod tidy && go build -o main cmd/ordersystem/main.go

CMD ["./main"] 