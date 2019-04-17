package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/mesment/personblog/models"
	"log"
)

func GetArticleViewCount(articleId string) (count int, err error) {
	var redisZsetKey = "articles:page.view"
	conn := RedisPool().Get()
	defer  conn.Close()
	count, err = redis.Int(conn.Do("ZINCRBY", redisZsetKey, "1",articleId))
	if err != nil {
		log.Printf("获取文章(id:%s)阅读次数失败：%s",articleId,err.Error())
	}
	return count, err
}

func AddArticleComment(articleId string, comment interface{}) error {
	//list key: articles:文章id:comments
	var commentListKey = fmt.Sprintf("articles:%s:comments",articleId)

	//将comment转换成json数据
	if data, err := json.Marshal(comment); err == nil {
		//从池中获取连接
		conn := RedisPool().Get()
		if conn == nil {
			err := fmt.Errorf("Redis 连接为nil")
			return err
		}
		//用完放回连接池
		defer conn.Close()
		_, err := conn.Do("LPUSH",commentListKey,string(data))
		if err != nil {
			log.Fatalf("redis存储文章（id：%s）评论失败:%s",articleId,err.Error())
		}
		return nil
	} else {
		return err
	}

	return nil
}

//取前100条
func GetArticleComments(articleId string) (comments []models.Comment,err error ){
	//get key
	var commentListKey = fmt.Sprintf("articles:%s:comments",articleId)
	//获取redis连接
	conn := RedisPool().Get()
	if conn == nil {
		err := fmt.Errorf("Redis 连接为nil")
		return nil,err
	}
	defer conn.Close()


	result, err := redis.Strings(conn.Do("LRANGE",commentListKey,"0","99"))
	if err != nil {
		log.Printf("获取文章(id:%s)评论列表失败%s",articleId,err.Error())
		return nil, err
	}

	for _, v := range result {
		comment := models.Comment{}
		err := json.Unmarshal([]byte(v),&comment)
		if err != nil {
			break
		}
		comments = append(comments, comment)
	}

	return
}
