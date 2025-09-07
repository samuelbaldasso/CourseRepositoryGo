FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o plataforma-cursos ./cmd/main.go

# Imagem final, menor e mais segura
FROM alpine:3.19

WORKDIR /app

RUN apk --no-cache upgrade && apk --no-cache add ca-certificates

COPY --from=builder /app/plataforma-cursos .

EXPOSE 8080

CMD ["./plataforma-cursos"]
