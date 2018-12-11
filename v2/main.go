package main

import (
	"api_server/v2/config"
	"api_server/v2/router"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {

	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

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

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i <= viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		fmt.Printf("正在ping URL : %v\n", viper.GetString("url"))
		log.Print("等待路由,一秒后重试")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接路由")

}
