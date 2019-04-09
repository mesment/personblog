build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t personblog .
run:
	# -p在容器的8080端口上运行,映射到主机的80端口, -d服务放到后台运行
	docker run -p 8080:8080 personblog
