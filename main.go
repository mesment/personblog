package main

import (
	"github.com/mesment/personblog/dao/db"
	"github.com/mesment/personblog/dao/redis"
	"github.com/mesment/personblog/models"
	"github.com/mesment/personblog/pkg/setting"
	"github.com/mesment/personblog/routers"
	"github.com/mesment/personblog/pkg/logger"
)

func setup()  {
	setting.Setup()
	db.SetUp()
	models.SetUp()
	logger.Setup()
	redis.Setup()
}

func main() {

	setup()
	router := routers.InitRouter()
	db := db.Conn()
	logger.Setup()
	defer db.DB.Close()
	defer logger.Close()

	router.Run() // listen and serve on 0.0.0.0:8080
}
