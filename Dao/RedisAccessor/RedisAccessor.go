package RedisAccessor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type RedisAccessorTest struct {
}

func InitRedis() (error, *redis.Client) {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		return err, redisDB
	}
	return nil, redisDB
}
func (r RedisAccessorTest) Test(c *gin.Context) {
	err, _ := InitRedis()
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println("Yes")
}
