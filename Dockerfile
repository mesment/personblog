FROM alpine:latest
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone 
#在容器根目录下创建app目录
RUN mkdir /app
#将工作目录切换到/app下
WORKDIR /app
#将当前目录下的文件拷贝到/app下
COPY .  /app
#声明端口
EXPOSE 8080:8080
#运行服务
CMD ["./personblog"]
