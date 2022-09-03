#1.编译程序源码
FROM golang:1.17.5 AS builder
#设置gomod代理
ENV GOPROXY https://goproxy.cn
#启动gomod
ENV GO111MODULE on
ENV PROJECTNAME generalapp
#将编译所需必要文件添加过来
ADD . /home/generalapp
WORKDIR /home/generalapp/
#编译 使用纯go编译
ENV CGO_ENABLE 0
RUN go build -tags netgo -o generalapp .

#2.构建运行镜像
FROM alpine
RUN sed -i 's!dl-cdn.alpinelinux.org/!mirrors.ustc.edu.cn/!g' /etc/apk/repositories \
# && sed -i 's!https://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories \
&&  apk add --no-cache ca-certificates \
&&  update-ca-certificates \
&&  apk --no-cache add tzdata  \
&&  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&&  echo "Asia/Shanghai" > /etc/timezone
ENV TZ=Asia/Shanghai
WORKDIR /home/generalapp/
#将上一阶段的产物复制过来 解决镜像过大的问题
COPY generalapp .
RUN chmod +x generalapp
#启动命令
EXPOSE 8001
CMD ["./generalapp"]

