FROM golang:1.17 as builder
RUN mkdir -p /
WORKDIR /
COPY . .
RUN make build


FROM alpine:3.14
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates
WORKDIR /
USER nobody

COPY --from=builder service .

ENTRYPOINT ["./service"]