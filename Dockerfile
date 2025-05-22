# Build stage
FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /bin/quote-api ./cmd/main.go

# Run stage
FROM alpine:3.18
RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /
COPY --from=builder /bin/quote-api /quote-api
RUN chmod +x /quote-api

EXPOSE 8080
CMD ["/quote-api"]