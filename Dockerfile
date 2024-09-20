FROM golang:1.23.1

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o ./bin/blockchain-supply-service ./cmd/http/main.go

CMD ["./bin/blockchain-supply-service"]
