//+build dev
package config

const (
	// 是否开启debug模式
	DEBUG = true
	ENV = "DEV"


	// 服务监控地址
	HTTPADDR = ":8090"

	// eureka
	EUREKAAPP = "ad_test"
	HTTPPORT = 80

	// redis
	REDISURL = "localhost:6379"
	REDISPASSWORD = ""
)