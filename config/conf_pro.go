//+build !dev

package config

const (
	// 是否开启debug模式
	DEBUG = false
	ENV = "PRO"


	// 服务监控地址
	HTTPADDR = ":80"

	// eureka
	EUREKAAPP = "ad_test"
	HTTPPORT = 80

	// redis
	REDISURL = ""
	REDISPASSWORD = ""

)