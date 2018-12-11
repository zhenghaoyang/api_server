package main

import (
	"api_server/v6/config"
	"api_server/v6/model"
	"api_server/v6/router"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	pflag.Parse()
	//init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init db
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("路由没有响应", err)
		}
		log.Info("路由部署成功")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())

}
func pingServer() error {
	for i := 0; i <= viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		fmt.Printf("正在ping URL : %v\n", viper.GetString("url"))
		log.Info("等待路由,一秒后重试")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接路由")

}
