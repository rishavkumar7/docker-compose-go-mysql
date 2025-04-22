# Stage 1: Build the Go app
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a mini image of the Go app
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

RUN apk --no-cache add ca-certificates tzdata

CMD ["./main"]