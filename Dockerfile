FROM golang:1.23.2-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

RUN apk add --no-cache gcc musl-dev

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

#Stage 2 : Running the application

FROM alpine:latest
RUN apk add --no-cache sqlite

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Copy the SQLite database into the container
COPY ./anime.db ./anime.db
EXPOSE 3000
CMD ["./main"]
