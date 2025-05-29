# Сервер
FROM golang:1.24 AS builder
WORKDIR /app
COPY ../../go.mod ../../go.sum ./
COPY ../../cmd/server ./cmd 
COPY ../../internal/server/serverAPI ./internal/server/serverAPI
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/tasks ./cmd/main.go

FROM golang:1.24-alpine AS production
WORKDIR /app
COPY --from=builder /app/bin/tasks ./
ENTRYPOINT [ "/app/tasks" ]