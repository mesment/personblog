# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.1.26-MariaDB)
# Database: myblog
# Generation Time: 2019-04-09 08:12:46 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table admins
# ------------------------------------------------------------

DROP TABLE IF EXISTS `admins`;

CREATE TABLE `admins` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table article
# ------------------------------------------------------------

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint(20) unsigned NOT NULL COMMENT '分类id',
  `content` longtext NOT NULL COMMENT '文章内容',
  `title` varchar(1024) NOT NULL COMMENT '文章标题',
  `view_count` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '阅读次数',
  `comment_count` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '评论次数',
  `username` varchar(128) NOT NULL DEFAULT '""' COMMENT '作者',
  `status` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `summary` varchar(256) NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_view_count` (`view_count`) USING BTREE COMMENT '阅读次数索引',
  KEY `idx_comment_count` (`comment_count`) USING BTREE COMMENT '评论数索引',
  KEY `idx_category_id` (`category_id`) USING BTREE COMMENT '分类id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;

INSERT INTO `article` (`id`, `category_id`, `content`, `title`, `view_count`, `comment_count`, `username`, `status`, `summary`, `create_time`, `update_time`)
VALUES
	(1,2,'$docker-compose up -d   启动容器\r\n$docker-compose down 停止容器\r\n$docker-compose rm  将停止的容器删除\r\n$docker-compose build  重新创建镜像','docker 常用命令',0,0,'\"\"',1,'$docker-compose up -d   启动容器\r\n$docker-compose down 停止容器\r\n$docker-compose rm  将停止的容器删除\r\n$docker-compose build  重新创建镜像','2019-01-07 21:44:27','2019-01-07 13:19:20'),
	(2,2,'新建一个目录learngit\r\n1、初始化仓库\r\n在learngit目录下执行\r\ngit init 初始化一个空的仓库\r\n\r\n2、 新建一个文件readme.txt\r\n3、把文件添加到git仓库暂存区\r\ngit add readme.txt\r\n4、把文件提交到仓库当前分支 git commit命令，-m后面输入的是本次提交的说明\r\ngit commit -m “wrote a readme file\"\r\n5、查看提交过的git版本\r\ngit log         //git log —graph 命令可以看到分支合并图\r\n6、退回到上一个版本(HEAD是当前版本)\r\ngit reset  - - hard HEAD^\r\ngit reset HEAD readme.txt  撤销readme.txt提交到暂存区的修改，重新放回工作区\r\n7、退回到上上个版本\r\ngit reset  HEAD^^   ,退回到前10个版本git reset HEAD~10\r\n8  退回到某个指定的版本\r\ngit reset - - hard  版本号\r\n9、显示执行过的历史命令\r\ngit reflog\r\n10、将git 工作区的文件撤销修改，恢复到暂存区(添加到暂存区后有新的修改)或版本库的状态\r\ngit checkout  - - readme.txt  把这个文件恢复到最近一次git add 或者git commit 时的状态\r\n11、从版本库里删除一个文件test.txt\r\ngit rm test.txt\r\ngit commit -m “delete test.txt\"\r\n12、 本地错误的删除了一个文件test.txt,需要进行恢复\r\ngit checkout  - - test.txt  实际是用版本库的版本替换工作区的版本\r\n13、把本地仓库的内容推送到远程\r\n     （1）关联远程仓库\r\n     git remote add origin git@github.com:username/learngit.git        //origin 是远程仓库的名字\r\n     （2）推送master分支的内容，我们第一次推送master分支时，参数-u会把本地的分支master内容推送的远程新的master分支，还会把本地的master分支和远程的master分支关联起来，后续的推送可以不用再加-u 参数\r\n     git push -u origin master 把当前master分支推送到远程并跟远程master分支关联\r\n     （3）后续本非提交后就可以将本地的最新修改推送到远程 仓库\r\n      git push origin master\r\n14、 从远程仓库克隆一个项目\r\n          git clone git@github.com:username/projectname.git\r\n15、创建一个dev分支，并切换到该分支\r\n     git checkout -b dev   //创建并切换分支\r\n或者使用下面两条命令完成同样的功能\r\n     git branch dev  //创建分支dev\r\n     git checkout dev  //切换到dev分支\r\n16、 查看当前分支, git branch命令会列出所有分支，当前分支前面会标一个*号\r\n     git branch\r\n17、在分支上修改和提交后可以切换回主分支\r\n     git checkout master   //注意在分支dev上做的修改在主分支上是看不到的（dev的修改还未合并到主分支）\r\n18、将分支dev的修改合并到当前分支(master)\r\n     git merge dev\r\n19、删除分支, 分支的内容合并到主分支后可以将分支删除\r\n     git branch -d dev\r\n\r\n20、 当master分支和dev同时做了修改并分别都进行了提交，合并时可能会出现冲突\r\n需要手工修改冲突内容后，重新git add, git commit 完成合并\r\n\r\n21、合并分支时，如果可能，Git会用Fast forward模式，但这种模式下，删除分支后，会丢掉分支信息。\r\n如果要强制禁用Fast forward模式，Git就会在merge时生成一个新的commit，--no-ff参数，表示禁用Fast forward：git merge --no-ff -m \"merge with no-ff\" dev\r\n\r\n22、分支策略\r\n在实际开发中，我们应该按照几个基本原则进行分支管理：\r\n首先，master分支应该是非常稳定的，也就是仅用来发布新版本，平时不能在上面干活；\r\n那在哪干活呢？干活都在dev分支上，也就是说，dev分支是不稳定的，到某个时候，比如1.0版本发布时，再把dev分支合并到master上，在master分支发布1.0版本。\r\n你和你的小伙伴们每个人都在dev分支上干活，每个人都有自己的分支，时不时地往dev分支上合并就可以了。\r\nGit分支十分强大，在团队开发中应该充分应用。\r\n合并分支时，加上--no-ff参数就可以用普通模式合并，合并后的历史有分支，能看出来曾经做过合并，而fast forward合并就看不出来曾经做过合并。\r\n\r\n23、Bug修复 有了bug就需要修复，在Git中，由于分支是如此的强大，所以，每个bug都可以通过一个新的临时分支来修复，修复后，合并分支，然后将临时分支删除。加入bug修复任务紧急，当前工作区dev有新功能开发到一半无法合并可以使用git stash把当前工作现场储存起来，等以后恢复现场后继续工作\r\n     git stash\r\n\r\n24 如果要在主分支master上修复bug，先从master上创建分支\r\n     (1)\r\n     git branch master\r\n     git checkout -b issue001\r\n     (2)完成bug修改并提交\r\n      git add readme.txt\r\n      git commit -m “bug fix\"\r\n     (3)修复完 切回主分支，完成合并并删除issue001分支\r\n     git branch master\r\n     git merge - - no-ff -m “merge bug fix 001”  issue001\r\n     git branch -d issue001\r\n     (4)切换回dev工作区继续之前的开发工作\r\n     git checkout dev\r\n     (5)查看刚才的工作现场存到哪去了？\r\n     git stash list\r\n     (6)恢复工作现场\r\n     方法一：\r\n     git stash apply 回复后，stash内容并不删除，需要用下个命令删除\r\n     git stash drop\r\n     方法二：\r\n     git stash pop 恢复的同时删除stash内容\r\n修复bug时，我们会通过创建新的bug分支进行修复，然后合并，最后删除；\r\n当手头工作没有完成时，先把工作现场git stash一下，然后去修复bug，修复后，再git stash pop，回到工作现场。\r\nGit鼓励大量使用分支：\r\n查看分支：git branch\r\n创建分支：git branch <name>\r\n切换分支：git checkout <name>\r\n创建+切换分支：git checkout -b <name>\r\n合并某分支到当前分支：git merge <name>\r\n\r\n25、添加一个新功能时，你肯定不希望因为一些实验性质的代码，把主分支搞乱了，所以，每添加一个新功能，最好新建一个feature分支，在上面开发，完成后，合并，最后，删除该feature分支。例如需要新增一个转账功能。\r\n     git checkout -b feature-transfer  //创建分支\r\n     git add     transfer.go      //t添加修改\r\n     git status     //查看状态\r\n     git commit -m “add transfer function\"\r\n     git checkout dev      //切换回dev分支\r\n     git merge feature-transfer //合并feature-transfer\r\n     git branch -d feature-transfer  //删除分支\r\n26、如果新分支的内容开发到一半，提交后还没有合并，某种原因该功能又不需要了，需要将该分支销毁\r\n     git  branch -d feature-transfer\r\n会提示销毁失败，feature-transfer还没有合并，删除将丢失修改，如果要强行删除需要添加-D参数\r\n git  branch -D feature-transfer\r\n开发一个新feature，最好新建一个分支；\r\n如果要丢弃一个没有被合并过的分支，可以通过git branch -D <name>强行删除。\r\n\r\n27、当从远程仓库克隆时，实际上Git自动把本地的master分支和远程的master分支对应起来了，并且，远程仓库的默认名称是origin\r\n查看远程仓库的信息用git remote\r\n     git remote          //git remote -v 显示更多详细信息\r\n\r\n28、推送分支。把该分支上的所有本地提交推送到远程库。推送时，要指定本地分支，这样，Git就会把该分支推送到远程库对应的远程分支上：\r\n         git  push origin master\r\n如果要推送其他分支如dev:\r\n          git push origin dev\r\n但是，并不是一定要把本地分支往远程推送，那么，哪些分支需要推送，哪些不需要呢？\r\n\r\n- master分支是主分支，因此要时刻与远程同步；\r\n- dev分支是开发分支，团队所有成员都需要在上面工作，所以也需要与远程同步；\r\n- bug分支只用于在本地修复bug，就没必要推到远程了，除非老板要看看你每周到底修复了几个bug；\r\n- feature分支是否推到远程，取决于你是否和你的小伙伴合作在上面开发。\r\n总之，就是在Git中，分支完全可以在本地自己藏着玩，是否推送，视你的心情而定！\r\n\r\n29、抓取分支。多人协作时，大家都会往master和dev分支上推送各自的修改。\r\n     (1)从远程仓库克隆后，默认本地只能看到master分支，要在dev分支上开发就必须创建远程origin的分支dev到本地：\r\n     git  clone git@github.com:username/project.git\r\n     git branch                                   //查看当前分支\r\n     git checkout -b dev origin/dev   //创建本地dev分支\r\n     (2)在dev分支上做了开发完后，提交修改并push到远程\r\n          git add xxxx.go\r\n          git commit -m “new function”     //提交修改\r\n          git push origin dev                     //push dev 分支到远程仓库origin\r\n     (3)假如你的小伙伴的最新提交到dev仓库的修改和你试图推送的提交有冲突，需要先将小伙伴已提交到远程仓库的内容\r\n拉取到本地，在本地合并解决冲突后再重新推送。\r\n          git pull  //如果拉取失败，原因是没有指定本地dev分支与远程origin/dev分支的链接，根据提示，设置dev和origin/dev的链接：\r\n          git branch - - set-upstream-to=origin/dev  dev\r\n          git pull               //再次拉取\r\n\r\n因此，多人协作的工作模式通常是这样：\r\n- 首先，可以试图用git push origin <branch-name>推送自己的修改；\r\n- 如果推送失败，则因为远程分支比你的本地更新，需要先用git pull试图合并；\r\n- 如果合并有冲突，则解决冲突，并在本地提交；\r\n- 没有冲突或者解决掉冲突后，再用git push origin <branch-name>推送就能成功！\r\n如果git pull提示no tracking information，则说明本地分支和远程分支的链接关系没有创建，用命令git branch --set-upstream-to <branch-name> origin/<branch-name>。\r\n- 查看远程库信息，使用git remote -v；\r\n- 本地新建的分支如果不推送到远程，对其他人就是不可见的；\r\n- 从本地推送分支，使用git push origin branch-name，如果推送失败，先用git pull抓取远程的新提交；\r\n- 在本地创建和远程分支对应的分支，使用git checkout -b branch-name origin/branch-name，本地和远程分支的名称最好一致；\r\n- 建立本地分支和远程分支的关联，使用git branch --set-upstream branch-name origin/branch-name；\r\n- 从远程抓取分支，使用git pull，如果有冲突，要先处理冲突。\r\n- 这就是多人协作的工作模式，一旦熟悉了，就非常简单。\r\n\r\n30、标签管理\r\n\r\n发布一个版本时，我们通常先在版本库中打一个标签（tag），这样，就唯一确定了打标签时刻的版本。将来无论什么时候，取某个标签的版本，就是把那个打标签的时刻的历史版本取出来。所以，标签也是版本库的一个快照。\r\n\r\nGit的标签虽然是版本库的快照，但其实它就是指向某个commit的指针（跟分支很像对不对？但是分支可以移动，标签不能移动），所以，创建和删除标签都是瞬间完成的。\r\n\r\nGit有commit，为什么还要引入tag？\r\n“请把上周一的那个版本打包发布，commit号是6a5819e...”\r\n“一串乱七八糟的数字不好找！”\r\n如果换一个办法：\r\n“请把上周一的那个版本打包发布，版本号是v1.2”\r\n“好的，按照tag v1.2查找commit就行！”\r\n所以，tag就是一个让人容易记住的有意义的名字，它跟某个commit绑在一起。\r\n\r\n31、git打标签非常简单，先切换到需要打标签的分支上\r\n     git branch                    //查看分支\r\n     git checkout master     //切换到master分支\r\n     git tag v1.0                    //打标签V1.0\r\n\r\n32、查看所有标签\r\n     git tag\r\n\r\n33、默认抱歉是打在最新提交的commit上的。有时候，如果忘了打标签，现在已经是周五了，但应该在周一打的标签没有打\r\n只要找到历史提交的commit id，然后打上就可以了：\r\n\r\ngit log --pretty=oneline --abbrev-commit    //显示历史命令\r\n\r\n比方说要对add merge这次提交打标签，它对应的commit id是f52c633，敲入命令：\r\n     git tag  v0.9 f52c633\r\n\r\n34、标签是按照字母顺序排列。可以用git show <tagname>查看标签信息：\r\n     git show v1.0\r\n\r\n35、可以 创建带有说明的标签，用-a指定标签名，-m指定说明文字：\r\n     git tag -a v1.0 -m “version 1.0 released”   1094adb\r\n- 命令git tag <tagname>用于新建一个标签，默认为HEAD，也可以指定一个commit id；\r\n- 命令git tag -a <tagname> -m \"blablabla...\"可以指定标签信息；\r\n- 命令git tag可以查看所有标签。\r\n\r\n36、操作标签。创建的标签都只存储在本地，不会自动推送到远程\r\n     git tag -d v0.1            //删除标签\r\n\r\n37、如果要推送某个标签到远程，使用     git push origin <tagname>\r\n     git push origin v1.0\r\n\r\n38、一次性推送全部尚未推送到远程的本地标签\r\n     git push origin  - - tags\r\n\r\n39、如果标签已推送到远程，要删除远程的标签，需要先从本地删除，然后从远程删除\r\n     git tag -d v0.9                               //删除本地tag v0.9\r\n     git push origin  :refs/tags/v0.9     //删除远程的tag v0.9\r\n\r\n- 命令git push origin <tagname>可以推送一个本地标签；\r\n- 命令git push origin --tags可以推送全部未推送过的本地标签；\r\n- 命令git tag -d <tagname>可以删除一个本地标签；\r\n- 命令git push origin :refs/tags/<tagname>可以删除一个远程标签。\r\n\r\n\r\n','Git 相关的操作',1,0,'\"\"',1,'新建一个目录learngit\r\n1、初始化仓库 在learngit目录下执行\r\ngit init 初始化一个空的仓库\r\n2、 新建一个文件readme.txt\r\n3、把文件添加到git仓库暂存区\r\ngit add readme.txt','2019-01-08 00:48:29','2019-04-09 13:26:35'),
	(3,2,'go get -u github.com/golang/protobuf/protoc-gen-go\r\ngo install github.com/golang/protobuf/protoc-gen-go\r\n\r\ngo get -u github.com/golang/protobuf/{proto,protoc-gen-go}\r\n\r\nERROR:package google.golang.org/genproto/googleapis/rpc/status: unrecognized import path \"google.golang.org/genproto/googleapis/rpc/status\" (https fetch: Get https://google.golang.org/genproto/googleapis/rpc/status?go-get=1: unexpected EOF)\r\n解决方案：\r\n1、cd $GOPATH/src/google.golang.org\r\n2、git clone https://github.com/google/go-genproto\r\n3、mv -f go-genproto  genproto','Protobuf &grpc',0,0,'\"\"',1,'go get -u github.com/golang/protobuf/protoc-gen-go\r\ngo install github.com/golang/protobuf/protoc-gen-go\r\ngo get -u github.com/','2019-01-21 11:25:46','2019-04-09 13:21:25'),
	(4,2,'1. 安装docker 建议上官网下载安装包安装。\r\n$ brew install docker\r\n2. 下载mysql镜像\r\n$ docker pull mysql\r\n3. 启动mysql实例\r\n$ docker run --name mingxie-mysql -p 32xxx:3306 -e MYSQL_ROOT_PASSWORD=1234 -d mysql:latest\r\n--name 后面的是docker容器名\r\n-p 32xxx:3306 这里需要注意 `32xxx` 是你**链接mysql的时候的`Port`。**\r\n-e MYSQL_ROOT_PASSWORD 是设置mysql的root账号密码\r\n-d mysql 是你的镜像标签\r\n4. 在shell中访问mysql\r\ndocker exec -it mingxie-mysql bash\r\nroot@7c289aa0ca95:/#\r\nmysql -uroot -p -h localhost\r\nEnter password:\r\n输入密码即可。\r\n5. 在shell中访问mysql日志\r\n$ docker logs mingxie-mysql\r\n\r\n\r\n\r\n\r\n','使用docker运行mysql',7,0,'\"\"',1,'1. 安装docker 建议上官网下载安装包安装。\r\n$ brew install docker\r\n2. 下载mysql镜像\r\n$ docker pull mysql\r\n3. 启动mysql实例\r\n$ docker run --name ming','2019-03-09 11:29:37','2019-04-09 13:22:24'),
	(42,2,'$docker images  显示容器中所有镜像文件\r\nfabric-samples/bin:\r\n     cryptogen 用来生成组织结构以及相应的证书秘钥\r\n                    在联盟链中有哪些组织，以及对应组织下有哪些节点\r\n                         Orderer\r\n                         Org\r\n                                   org1\r\n                                        peer0.org1.example.com\r\n                                        peer1.org1.example.com\r\n                                   org2\r\n                                        peer0.org2.example.com\r\n                                        peer1.org2.example.com\r\n                         什么组织可以访问通道中的数据依赖于生成的证书和秘钥\r\nconfigtxgen\r\n           1：用来生成Orderer 服务的初始区块(创世区块)\r\n            2：可以生成对应的通道交易配置文件（包含了通道中的成员及访问策略）\r\n            3： 生成Anchor (锚节点更新文件)\r\n                  （1）用来跨组织的数据交换\r\n                   （2）发现通道内新加入的组织/节点\r\n          每个组织都会有一个Anchor 节点\r\nconfigtxlator: 用来添加新组织（有新组织想要加入通道的时候使用）\r\nbin： 相应的工具目录\r\nchain code：链码示例\r\nchain code -docker-devmode:在开发模式下测试的环境\r\nfabcar  :提供node.js的示例\r\nfabric-ca : 简单的证书\r\nfirst-network: 搭建fabric网络的目录\r\n\r\nconfig/\r\n 关于order和peer配置信息的文件存放目录\r\nconfigtx.ymal ：生成初始区块以及应用通道配置文件的参考\r\ncore.ymal :peer配置信息的参考\r\norderer.ymal:orderer配置信息的参考','fabric目录简介',0,0,'\"\"',1,'$docker images  显示容器中所有镜像文件\r\nfabric-samples/bin:\r\n     cryptogen 用来生成组织结构以及相应的证书秘钥\r\n                    在联盟链中有哪些组织，以及对应组织下有哪些节点\r','2019-04-09 15:53:05','2019-04-09 15:53:05'),
	(43,2,'手动模式启动网络\r\n在frist-network 目录下执行\r\n../bin/cryptogen --help  查看命令帮助文档\r\n\r\n生成的组织结构\r\n   在当前目录下生成目录 crypto-config\r\n                                             orderer..\r\n                                             peer...\r\n                                                  org1\r\n                                                       peer0.org1...\r\n                                                       peer1.org1…\r\n                                                  org2\r\n                                                       peer0.org2...\r\n                                                       peer1.org2…\r\n\r\n../bin/cryptogen  show template  查看组织结构的模板配置文件\r\n\r\n模板文件位于first-network 目录下的crypto-config.yaml 文件中\r\n\r\n  - Name: Orderer\r\n    Domain: example.com\r\n    Specs:\r\n      - Hostname: orderer\r\n\r\nPeerOrgs:\r\n  - Name: Org1\r\n    Domain: org1.example.com\r\n    EnableNodeOUs: true\r\n    Template:\r\n      Count: 2\r\n    Users:\r\n      Count: 1\r\n  - Name: Org2\r\n    Domain: org2.example.com\r\n    EnableNodeOUs: true\r\n    Template:\r\n      Count: 2   （Org2组织下的节点的个数）\r\n    Users:\r\n      Count: 1 (节点的用户数)\r\n\r\n使用tree命令查看执行目录的树形结构\r\n$ tree -L num  指定目录  —num表示要显示几级目录\r\ntree -L 4 ./\r\n\r\n启动网络手动实现\r\n实现步骤\r\n1、生成组织关系和身份证书\r\n\r\n确定在fabric-sample/first-network 目录下\r\n$ cd fabric-sample/first-network/\r\n为fabric网络生成指定拓扑结构的组织关系和身份证书\r\n$ ../bin/cryptogen generate  — config=./crypto-config.yaml\r\n以上命令依赖crypto-config.yaml配置文件\r\n输出如下\r\norg1.example.com\r\norg2.example.com\r\n证书和秘钥（即MSP材料）被输出到目录first-network/crypto-config的目录中ordererOrganizations子目录下包括\r\n构成Orderer组织（1个Orderer节点）的身份信息，peerOrganizations子目录下为所有Peer节点组织（两个组织，四个节点）的\r\n相关身份信息。其中最关键的是MSP目录，代表了实体的身份信息。\r\n\r\n在修改yaml配置文件的时候需要注意每一层前面要有两个空格\r\n\r\n2、生成Orderer初始区块配置文件:使用configtxgen\r\n$../bin/configtxgen --help\r\n$ export CHANNEL_NAME=mychannel 定义通道名称环境变量\r\n$../bin/configtxgen -profile TwoOrgsOrdererGenesis  -channelID $CHANNEL_NAM -outputBlock ./channel-artifacts/genisis.block通道名称  E\r\n\r\n其中 -TwoOrgsOrdererGenesis 是模板名称位于first-network/configtx.yaml文件里\r\n\r\n生成应用通道配置文件时，必须指定通道名，通道中包含的组织信息。\r\n                    --组织信息属于某个联盟\r\n锚节点是在生成应用通道时通过配置文件指定\r\n\r\n3、生成应用通道交易配置文件\r\n$ ../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAM\r\n\r\n4、生成锚节点更新配置文件\r\n$ ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAM -asOrg Org1MSP\r\n\r\n$ ../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAM -asOrg Org2MSP\r\n\r\n5、启动网络\r\n$ docker-compose -f\r\n                         -f 指定启动网络时所使用的配置文件（该配置文件中描述了启动网络时有哪些容器被启动，指定容器中所挂载的内容）','手动启动网络生成组织结构和身份证书',0,0,'\"\"',1,'手动模式启动网络\r\n在frist-network 目录下执行\r\n../bin/cryptogen --help  查看命令帮助文档\r\n\r\n生成的组织结构\r\n   在当前目录下生成目录 crypto-config\r\n                     ','2019-04-09 16:00:10','2019-04-09 16:00:10');

/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL COMMENT '分类名字',
  `category_no` int(10) unsigned NOT NULL COMMENT '分类排序',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;

INSERT INTO `category` (`id`, `category_name`, `category_no`, `create_time`, `update_time`)
VALUES
	(1,'CSS/HTML',1,'2018-08-12 10:55:45','2019-04-09 13:16:05'),
	(2,'服务端开发',2,'2018-08-12 10:56:07','2019-04-09 13:16:27'),
	(3,'Linux/Unix',3,'2018-08-12 10:56:16','2019-04-09 13:15:55'),
	(4,'C++开发',4,'2018-08-12 10:56:24','2018-08-12 10:59:08'),
	(5,'架构剖析',5,'2018-08-12 10:56:36','2018-08-12 10:59:10'),
	(6,'Golang开发',6,'2018-08-12 10:56:45','2018-08-12 10:59:14');

/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table comment
# ------------------------------------------------------------

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text NOT NULL COMMENT '评论内容',
  `username` varchar(64) NOT NULL COMMENT '评论作者',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) unsigned NOT NULL DEFAULT '1' COMMENT '评论状态: 0, 删除；1， 正常',
  `article_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;



# Dump of table message
# ------------------------------------------------------------

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;

INSERT INTO `message` (`id`, `username`, `content`, `create_time`, `update_time`)
VALUES
	(24,'mesment','nihao','2019-04-09 16:02:41','2019-04-09 16:02:41');

/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `name`, `password`)
VALUES
	(1,'mesment','123'),
	(2,'a','a');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
