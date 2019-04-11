# personblog
### 编译前请修改数据库用户名、密码.
1. 将项目克隆到本地.
1. 执行sqldb文件夹里的数据库myblog.sql脚本建立博客相关的表.
2. 修改main.go里的数据库用户名、密码.
3. 安装docker
4. 安装依赖编译并执行

#### 安装步骤
----
1. 安装mysql <br>
$docker pull mysql:5.7  //拉取镜像 <br>
$docker images		//查看镜像 <br>
----
2. 为了使所有容器在同一个网络内能够相互访问，首先创建一个网桥 <br>
docker network create -d bridge blog_network <br>
----
3. 运行mysql,根据自己的情况删减 参数 <br>
 docker container run -it --detach --name mysql  --network blog_network -p 3307:3306  -v $PWD/data/myscript/:/docker-entrypoint-initdb.d/ --env MYSQL_RANDOM_ROOT_PASSWORD=yes mysql:5.7 <br>
----
4. $ docker container logs mysql| grep 'GENERATED ROOT PASSWORD: ' | awk -F': ' '{print $2}' <br>
----
5. docker container exec -it mysql  bash  <br>
----
6. 登录mysql 修改root密码  <br>
  mysql -u root -p    <br>
----
7. 下面的这条是mysql5.7的版本，不同版本的字段不一样。 <br>
mysql>update mysql.user set authentication_string=password('123456') where user='root' ;   <br>
mysql> flush privileges;  <br>
mysql> exit;  <br>
----
8. 创建应用的Docker镜像  <br>
make build    <br>
----
9. 运行应用容器 <br>
make run  <br>
----
10. 清理应用容器和镜像 <br>
./clearblogdocker.sh <br>


