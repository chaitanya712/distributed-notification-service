FROM golang:1.24.1

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o main ./cmd/server

EXPOSE 50051
EXPOSE 2112
EXPOSE 8080

CMD ["./main"]
