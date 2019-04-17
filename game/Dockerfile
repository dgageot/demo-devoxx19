FROM golang:1.11.5-alpine3.9 as builder
WORKDIR /usr/local/go/src/github.com/dgageot/hexagons-game/
COPY main.go types.go ./
RUN go build -o /game .

FROM alpine:3.9
EXPOSE 8888
ENTRYPOINT ["/game"]
COPY --from=builder /game /
