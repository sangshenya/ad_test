package util

import (
	"net"
	//"github.com/prometheus/common/log"
	//"github.com/ww156/go-eureka"
	//"ad_test/config"
	//"strconv"
	//"time"
)

func RegistServer() {
	//ip := GetInternalIp()
	//if ip == "" {
	//	//log.Fatal("get ip error")
	//}
	//
	//serverUrl := []string{
	//
	//}
	//
	//e, err := eureka.NewEureka(serverUrl, Client)
	//if err != nil {
	//	panic(err)
	//}
	//
	//app := config.EUREKAAPP
	//port := config.HTTPPORT
	//ins := eureka.Instance{
	//	HostName: ip,
	//	App:app,
	//	Port: &eureka.Port{
	//		Port:port,
	//		Enable:true,
	//	},
	//	SecurePort: &eureka.Port{
	//		Port:443,
	//		Enable:false,
	//	},
	//	IPAddr:ip,
	//	HealthCheckUrl:"http://" + ip + ":" + strconv.Itoa(port) + "/health",
	//	StatusPageUrl: "http://" + ip + ":" + strconv.Itoa(port) + "/status",
	//	HomePageUrl:"http://" + ip + ":" + strconv.Itoa(port),
	//	Status:"UP",
	//	DataCenterInfo: &eureka.DataCenterInfo{
	//		Name:"MyOwn",
	//		Class:"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
	//	},
	//}
	//
	//e.RegisterInstane(&ins)
	//e.SendHeartBeat(&ins, time.Second * 20)

}


func GetInternalIp() string {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		 //log.Fatal(err)
	}
	for _,item := range ips {
		if ipnet, ok := item.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}