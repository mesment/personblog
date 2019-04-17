package setting

import (
	"gopkg.in/ini.v1"
	"log"
)


//对应ini配置文件中app配置
type App struct {
	PageSize string	//每页显示文章条数
	PagJWTSecret  string	//JWT 加密字符串
}
var AppCfg = &App{}

//对应ini配置文件中log配置
type Log struct {
	LogType string
	Level string
	LogPath string
	LogName string
}
var LogCfg  = &Log{}

//对应ini配置文件中server配置
type Server struct {
	RunModel  	string
	Port	int
	ReadTimeout  int
	WriteTimeout int
}
var ServerCfg  = &Server{}

//对应ini配置文件中database配置
type DataBase struct {
	DBType 	string  			//数据库类型mysql
	User 	string				//用户名
	PassWord	string			//密码
	Host 	string				// 主机地址 127.0.0.1:3306
	DBName 	string				//数据库名
}
var DBCfg  = &DataBase{}

//对应ini配置文件中redis配置
type Redis struct {
	Host 	string				// 主机地址 127.0.0.1
	Port 	string				    //端口号
	PassWord string
}
var RedisCfg  = &Redis{}


func Setup()  {
	var err error
	Cfg,err :=ini.LooseLoad("config/app.ini","config/server.ini")
	if err != nil {
		log.Fatalf("Failed to parse config app.ini.%s",err.Error())
		return
	}

	err = Cfg.Section("database").MapTo(DBCfg)
	if err != nil {
		log.Printf("setting 映射配置文件database失败：%s",err)
		return
	}

	err = Cfg.Section("server").MapTo(ServerCfg)
	if err != nil {
		log.Printf("setting 映射配置文件server失败：%s",err)
		return
	}

	err = Cfg.Section("app").MapTo(AppCfg)
	if err != nil {
		log.Printf("setting 映射配置文件app失败：%s",err)
		return
	}
	err = Cfg.Section("log").MapTo(LogCfg)
	if err != nil {
		log.Printf("setting 映射配置文件log失败：%s",err)
		return
	}

	err = Cfg.Section("redis").MapTo(RedisCfg)
	if err != nil {
		log.Printf("setting 映射配置文件redis失败：%s",err)
		return
	}

}

