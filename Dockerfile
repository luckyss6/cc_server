FROM golang:latest AS builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

# CGO_ENABLED禁用cgo 然后指定OS等，并go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

# 运行阶段指定scratch作为基础镜像
FROM scratch AS runner

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/main .

EXPOSE 8080

ENTRYPOINT ["./main"]