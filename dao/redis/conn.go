package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mesment/personblog/pkg/setting"
	"log"
	"time"
)

var (
	pool *redis.Pool

	Host string
	Port string
	redisHost  string
	redisPassWd string
)

func Setup()  {
	Host = setting.RedisCfg.Host
	Port= setting.RedisCfg.Port
	redisHost = Host + ":" +Port
	redisPassWd = setting.RedisCfg.PassWord

}




func newRedisPool() *redis.Pool {


	return &redis.Pool{
		MaxIdle:50,						//池中最多可用连接数
		MaxActive:50,					//同时最大连接数
		IdleTimeout: 300 *time.Second,  //连接超过5分钟没有使用就回收
		Dial:func()(redis.Conn,error){
			//打开连接
			c,err :=redis.Dial("tcp", redisHost)
			if err != nil {
				log.Println("redis连接失败：",err)
				return nil,err
			}

			//2、访问认证
			if _, err := c.Do("AUTH", redisPassWd); err != nil {
				if err != nil {
					log.Println("redis连接失败：",err)
				}
				c.Close()
			}

			return c, nil
		},
		TestOnBorrow:func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_,err := c.Do("PING")
			return err
		},
	}
}

func init()  {
	pool = newRedisPool()

}

func RedisPool() *redis.Pool  {
	return pool
}
