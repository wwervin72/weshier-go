FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /build

# 将代码复制到容器中
COPY . .

# RUN apk add --update iproute2 && \
# 	go build -o app .
RUN go build -o app .


EXPOSE 3333

# CMD ["ip -4 route list match 0/0 | awk ‘{print $3 “host.docker.internal”}’ >> /etc/hosts; /build/app"]
CMD ["/build/app"]
