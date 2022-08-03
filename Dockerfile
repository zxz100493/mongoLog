FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/app

ADD go.mod .
ADD go.sum .
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod download

RUN export GO111MODULE=on && \
export GOPROXY=https://goproxy.cn && \
go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/apps ./cmd
ADD app.yaml /app/
ADD resource /app/

FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/apps /app/apps
COPY --from=builder /app/app.yaml /app/app.yaml
COPY --from=builder /app/service.sh /app/service.sh
COPY --from=builder /app/resource /app/resource
RUN chmod +x /app/service.sh

EXPOSE 8080     
 
EXPOSE 3000     

CMD ["/app/service.sh"]

