FROM golang:1.22.5 AS builder

WORKDIR /app

COPY . .

RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go build -o GoToKube

FROM alpine:latest

LABEL authors="horonlee"

COPY --from=builder /app/GoToKube /usr/local/bin/GoToKube

RUN chmod +x /usr/local/bin/GoToKube

ENTRYPOINT ["/usr/local/bin/GoToKube"]
