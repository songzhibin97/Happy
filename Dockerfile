FROM golang:alpine

# 设置环境变量
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

# 设置工作目录
WORKDIR /go/happy

# Copy文件进行打包
COPY . .

# 下载依赖 并打包
RUN go env
RUN  go build -o server .

FROM alpine:latest
WORKDIR /go
RUN mkdir log
COPY ./config.ini ./
COPY --from=0 /go/happy/server ./

ENTRYPOINT ./server

