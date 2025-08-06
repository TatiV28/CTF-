FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o mod-ctf cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mod-ctf .
EXPOSE 8080
CMD ["./mod-ctf"]
