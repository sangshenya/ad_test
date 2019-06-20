package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ad_test/config"
	"time"
	"ad_test/util"
)

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":"success",
		"code":0,
	})
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"description":"fargo",
		"status":"UP",
	})
}

func data(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":"success",
		"data":"数据12",
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// eureka状态监测接口
	router.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// 状态监测
	router.Any("/status", status)
	router.Any("/health", health)

	// 获取数据
	router.GET("/ad/data", data)
	router.POST("/ad/data", data)

	s:= &http.Server{
		Addr: 			config.HTTPADDR,
		Handler:		router,
		ReadTimeout: 	10 * time.Second,
		WriteTimeout: 	10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 服务注册
	if config.ENV == "PRO" {
		util.RegistServer()
	}
	s.ListenAndServe()
}
