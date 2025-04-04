# Go dasturini rasmga olish
FROM golang:1.19-alpine as builder

# Ish joyini belgilash
WORKDIR /app

# Go mod fayllarini nusxalash va yuklash
COPY go.mod go.sum ./
RUN go mod tidy

# Dastur kodlarini nusxalash
COPY . .

# Dasturini qurish
RUN go build -o main ./cmd/server

# Ishlatish uchun konteyner rasmiga boshqa rasmni qo'llash
FROM alpine:latest

# Kerakli paketlarni o'rnatish
RUN apk add --no-cache ca-certificates

# Qurilgan dastur faylini nusxalash
COPY --from=builder /app/main /main

# Swagger uchun portni ochish
EXPOSE 8080

# Dasturini ishga tushirish
CMD ["/main"]
