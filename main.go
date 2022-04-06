package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Say string `json:"say"`
}

func main() {
	bindFlagPtr := flag.String("bind", ":8080", "bind address")
	disableGinModeFlagPtr := flag.Bool("disable-gin-debug-mode", false, "gin.ReleaseMode")
	flag.Parse()
	if *disableGinModeFlagPtr {
		gin.SetMode(gin.ReleaseMode)
	}
	log.Println("Welcome 🐱‍🏍")
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.0.0/24", "172.16.0.0/12", "10.0.0.0/8"})
	api := r.Group("/api/v1")
	api.GET("/ping", Ping)
	r.Run(*bindFlagPtr)
}

func Ping(c *gin.Context) {
	res := gin.H{
		"status":    http.StatusOK,
		"message":   "pong",
		"timestamp": GetUnixMillisTimestamp(),
	}

	var req Req
	if c.Request.Body != nil {
		err := c.ShouldBindJSON(&req)
		if err != nil && err != io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid JSON",
				"error":   err.Error(),
			})
			return
		}
		res["say"] = req.Say
	}

	c.JSON(http.StatusOK, res)
}

func GetUnixMillisTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
