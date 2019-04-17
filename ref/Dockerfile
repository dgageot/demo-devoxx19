FROM golang:1.11.5-alpine3.9 as builder
WORKDIR /usr/local/go/src/github.com/dgageot/hexagons-ref/
COPY vendor vendor
RUN find vendor -maxdepth 2 -mindepth 2 -type d -exec sh -c 'go install -i github.com/dgageot/hexagons/{}/... || true' \;
COPY main.go types.go ./
RUN go build -o /ref .

FROM alpine:3.9
EXPOSE 8888
ENTRYPOINT ["/ref"]
COPY --from=builder /ref /
