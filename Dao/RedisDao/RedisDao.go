package RedisDao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

type RedisDaoTest struct {
}

func initRedis() error {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func (r RedisDaoTest) Test(c *gin.Context) {
	err := initRedis()
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println("Yes")
}
