# Stage 1: Build Stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
# Build the binary (CGO_ENABLED=0 makes it a static binary that runs anywhere)
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Stage 2: Run Stage (The production image)
FROM alpine:3.21
WORKDIR /app
# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .
# Copy your migration files so the app can run them if needed
COPY db/migration ./db/migration

EXPOSE 8080
CMD ["/app/main"]