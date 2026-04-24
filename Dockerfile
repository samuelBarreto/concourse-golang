FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY *.go .
RUN go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8081
CMD ["./server"]
