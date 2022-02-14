package main

import (
	"Project/Dao/DBAccessor"
	"Project/Routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Routers.RegisterRouter(r)
	_, _ = DBAccessor.MySqlInit()
	r.Run(":80")
}
