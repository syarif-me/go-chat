FROM golang:1.26.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-chat-app ./src/cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-chat-app .

EXPOSE 8080

CMD ["./go-chat-app"]