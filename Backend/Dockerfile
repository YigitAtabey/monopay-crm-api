# 1) Build aşaması
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Modul tanımları ve bağımlılıklar
COPY go.mod go.sum ./
RUN go mod download

# Kodları kopyala ve derle
COPY . .
RUN go build -o main .

# 2) Runtime aşaması
FROM alpine:latest
WORKDIR /app

# Derleme sonucu ikili dosya ve config klasörü
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config

# Uygulamanın dinleyeceği port
EXPOSE 8083

# Direkt olarak main'i çalıştır
CMD ["./main"]
