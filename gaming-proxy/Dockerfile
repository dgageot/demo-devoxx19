FROM golang:1.11.5-alpine3.9 as builder
WORKDIR /usr/local/go/src/github.com/dgageot/hexagons-proxy/
COPY main.go .
RUN go build -o /proxy main.go

FROM alpine:3.9
EXPOSE 8888
ENTRYPOINT ["/proxy"]
COPY --from=builder /proxy /
