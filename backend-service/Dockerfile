# Build
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o app .

# Run
FROM scratch
COPY --from=builder /app/app /app
EXPOSE 8080
CMD ["/app"]