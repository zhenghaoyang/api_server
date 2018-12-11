package main

import (
	"api_server/v1/router"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {

	g := gin.New()
	//g的中间件
	middlewares := []gin.HandlerFunc{}

	//调用 router.Load 函数来加载路由
	router.Load(
		g,
		middlewares...
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("路由没有响应")
		}
		log.Print("路由部署成功")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}

func pingServer() error {
	for i := 0; i <= 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("等待路由,一秒后重试")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接路由")

}
