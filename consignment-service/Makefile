build:
    # 告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统
    GOOS=linux GOARCH=amd64 go build    
    # 根据当前目录下的 Dockerfile 生成名为 consignment-service 的镜像
    docker build -t consignment-service .

run:
    docker run -p 50051:50051 \
     -e MICRO_SERVER_ADDRESS=:50051 \
     -e MICRO_REGISTRY=mdns \
     consignment-service