# 1) The builder
FROM golang:1.25-alpine AS builder

WORKDIR /app

# 2. Only copy the dependency managers
COPY go.mod go.sum ./

# 3. Download the images from internet
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api-qr .

# -----------------------

# 2) Final container
FROM alpine:latest

WORKDIR /app

# Takes the "nanoqr-api" file created in the 1st step
COPY --from=builder /app/api-qr .

EXPOSE 8080

# When run, do:
CMD ["./api-qr"]