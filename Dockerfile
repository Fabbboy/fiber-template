FROM golang:1.23-alpine AS builder

WORKDIR /app
RUN apk add --no-cache make curl git bash libstdc++ gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/server cmd/main.go


FROM alpine:3.21
WORKDIR /app

COPY --from=builder /app/bin/server /app/server
ENV HOST=0.0.0.0:3000

CMD ["/app/server"]
