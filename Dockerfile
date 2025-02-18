# 定义第一步用来编译的环境(使用alpine镜像)
FROM golang:1.18-alpine AS go-alpine-builder
# 拷贝代码
COPY . /go/src/http-server
# 设置工作目录
WORKDIR /go/src/http-server
# 设置执行命令
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go build -o app

# 重制作镜像-通过dockerhub可查看golang1.18-alpine对应的alpine版本
FROM alpine:3.15
COPY --from=go-alpine-builder /go/src/http-server/app /bin/http-server

# 暴露端口
ENV PORT=9001
EXPOSE 9001

ENTRYPOINT [ "/bin/http-server" ]