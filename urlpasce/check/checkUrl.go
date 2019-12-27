package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	getCheckUrl()
}

//检测某个带端口的服务是否可用
func checkService() {
	timeout := time.Duration(1 * time.Second)
	conn, err := net.DialTimeout("tcp", "mysql:port", timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
	}
	fmt.Println(conn)
}

func getCheckUrl() {
	url := "https://www.baidu.com"
	clenit := http.Client{}
	responce, err := clenit.Get(url)
	if err != nil {
		fmt.Println(fmt.Sprintf("该网站访问失败：%s", err.Error()))
	}
	fmt.Println(fmt.Sprintf("status: %s", responce.Status))
}
