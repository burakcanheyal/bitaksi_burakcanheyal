FROM golang:1.25.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bitaksi

# Container başlatıldığında çalışacak komut
CMD ["/bitaksi"]
