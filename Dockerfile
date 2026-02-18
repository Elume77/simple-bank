# Stage 1: Build Stage
FROM golang:1.25-alpine AS builder

# Install build dependencies (alpine is very minimal and lacks curl/tar)
RUN apk add --no-cache curl tar

WORKDIR /app
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

RUN apk add --no-cache curl tar

# Download and extract migrate (Fixed the tar extraction logic)
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-amd64.tar.gz | tar xvz
# Note: The binary extracted is usually named 'migrate'

# Stage 2: Run Stage
FROM alpine:3.21
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/main .

COPY --from=builder /app/migrate /usr/local/bin/migrate

# Copy configuration and migrations
COPY db/migration ./db/migration
COPY app.env .

COPY start.sh .

COPY wait-for.sh .

RUN chmod +x start.sh 


ENTRYPOINT [ "/app/start.sh" ]

EXPOSE 8080


CMD ["./main"]