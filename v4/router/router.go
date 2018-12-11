package router

import (
	"api_server/v4/handler/sd"
	"api_server/v4/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	//设置中间件过滤器
	//router.Load 函数通过 g.Use() 来为每一个请求设置 Header
	// 在 router/router.go 文件中设置 Header：
	g.Use(gin.Recovery())
	g.Use(middleware.Print)
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "Incorret API route")
	})

	// 定义了一个叫 sd 的分组，在该分组下注册了 /health、/disk、/cpu、/ram HTTP 路径，
	// 分别路由到 sd.HealthCheck、sd.DiskCheck、sd.CPUCheck、sd.RAMCheck 函数
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}
	return g
}
