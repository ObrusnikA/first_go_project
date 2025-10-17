# Build stage
FROM golang AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM debian:bullseye-slim

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]