# Stage 1: Build Stage
# Updated from 1.24 to 1.25 to match your go.mod requirement
FROM golang:1.25-alpine AS builder
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
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]