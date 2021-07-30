# Build Stage 
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go build -o goss main.go

# Run Stage 
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/goss .


EXPOSE 1211
CMD ["/app/goss"]