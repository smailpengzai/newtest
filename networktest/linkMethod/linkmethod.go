package linkMethod

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"pingcap/log"
	"strings"
	"time"
)

//post 请求获取数据
//返回 []byte 请求到的数据
func PostUrl(requstUrl string) []byte {
	// AllAlarmConfigBool 这个为false表示没有获取到需要继续获取
	cententType := "application/x-www-form-urlencoded"
	resp, err := http.Post(requstUrl,
		cententType,
		strings.NewReader(""))
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(fmt.Sprintf("status: %s", resp.Status))
	defer resp.Body.Close()
	//拿出body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
	}
	return body

}

//检测某个带端口的服务是否可用
func checkService() {
	timeout := time.Duration(1 * time.Second)
	conn, err := net.DialTimeout("tcp", "mysql:port", timeout)
	if err != nil {
		fmt.Println("Site unreachable, error: ", err)
	}
	fmt.Println(conn)
}

//get检测
func getCheckUrl() {
	url := "https://www.baidu.com"
	clenit := http.Client{}
	responce, err := clenit.Get(url)
	defer responce.Body.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("该网站访问失败：%s", err.Error()))
	}
	fmt.Println(fmt.Sprintf("status: %s", responce.Status))
	body, _ := ioutil.ReadAll(responce.Body)
	fmt.Println(fmt.Sprintf("返回页面：\n %s", body))
}
