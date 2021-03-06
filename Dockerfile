FROM golang:1.11.5-alpine3.9 as builder
COPY main.go .
RUN go build -o /app

FROM alpine:3.9
CMD ["./app"]
COPY --from=builder /app .
