package util

import (
	"time"
	"net/http"
	"crypto/tls"
	"net"
)

var Timeout = 800 * time.Millisecond
var roundtrip = &http.Transport{

		Proxy:					http.ProxyFromEnvironment,
		MaxIdleConns: 			300,
		MaxIdleConnsPerHost:	50,
		IdleConnTimeout: 		90 * time.Second,
		TLSHandshakeTimeout:	10 * time.Second,
		ExpectContinueTimeout: 	1 * time.Second,
		DialContext: 			(&net.Dialer{
									Timeout: 30*time.Second,
									KeepAlive: 30 * time.Second,
									DualStack:true,
								}).DialContext,
		TLSClientConfig:		&tls.Config{
									InsecureSkipVerify:true,
								},
}

var Client = &http.Client{Timeout: Timeout, Transport: roundtrip}