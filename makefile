# bin/sh
.PHONY:all build run gotool clean help

SystemName = "Happy"

all:gotool build

build:
	go build -o ${SystemName}

run:
	go run ./
	
gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${SystemName} ] ; then rm ${SystemName} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"