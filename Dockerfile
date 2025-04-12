FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git && \
    go install github.com/swaggo/swag/cmd/swag@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init -g api/api.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/api/docs /app/docs
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
EXPOSE 5555
ENV GIN_MODE=release \
    PORT=5555

CMD ["/app/main"]