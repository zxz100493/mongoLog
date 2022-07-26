FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/app

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.cn
RUN go build -ldflags="-s -w" -o /app/apps ./cmd

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/apps /app/apps

CMD ["./apps"]

