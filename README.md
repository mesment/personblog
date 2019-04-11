# personblog
### 编译前请修改数据库用户名、密码.
1. 将项目克隆到本地.
1. 执行sqldb文件夹里的数据库myblog.sql脚本建立博客相关的表.
2. 修改main.go里的数据库用户名、密码.
3. 安装docker
4. 安装依赖编译并执行
####安装步骤
----
1. 安装mysql  

	$docker pull mysql:5.7  //拉取镜像
	$docker images		//查看镜像
----
2. 为了使所有容器在同一个网络内能够相互访问，首先创建一个网桥   

	docker network create -d bridge blog_network
----
3. 运行mysql,根据自己的情况删减 参数   

 docker container run -it --detach --name mysql  --network blog_network -p 3307:3306  -v $PWD/data/myscript/:/docker-entrypoint-initdb.d/ \
 --env MYSQL_RANDOM_ROOT_PASSWORD=yes mysql:5.7
----
4. $ docker container logs mysql| grep 'GENERATED ROOT PASSWORD: ' | awk -F': ' '{print $2}'
----
5. docker container exec -it mysql  bash
----
6. 登录mysql 修改root密码  

  mysql -u root -p
----
7. 下面的这条是mysql5.7的版本，不同版本的字段不一样。   

	mysql>update mysql.user set authentication_string=password('123456') where user='root' ;
	mysql> flush privileges;
	mysql> exit;
----
8. 创建应用的Docker镜像    

	make build  
----
9. 运行应用容器    

	make run
----
10. 清理应用容器和镜像    

	./clearblogdocker.sh


