package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type WebWeChat struct {
	token   string
	cookies []*http.Cookie
}

func NewWebWeChat() *WebWeChat {
	w := new(WebWeChat)
	return w
}

func (w *WebWeChat) login() bool {
	login_url := "https://mp.weixin.qq.com/cgi-bin/login?lang=zh_CN"
	email := "songbohr@163.com"
	password := "xxx"
	h := md5.New()
	h.Write([]byte(password))
	password = hex.EncodeToString(h.Sum(nil))
	fmt.Println(password)
	post_arg := url.Values{"username": {email}, "pwd": {password}, "imgcode": {""}, "f": {"json"}}

	fmt.Println(strings.NewReader(post_arg.Encode()))
	req, err := http.NewRequest("POST", login_url, strings.NewReader(post_arg.Encode()))
	req.Header.Set("Referer", "https://mp.weixin.qq.com/")

	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	resp, _ := client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)

	s := string(data)
	fmt.Printf("%s", s)

	doc := json.NewDecoder(strings.NewReader(s))

	type Msg struct {
		Ret                     int
		ErrMsg                  string
		ShowVerifyCode, ErrCode int
	}

	var m Msg
	if err := doc.Decode(&m); err == io.EOF {
		fmt.Println(err)
	} else if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println(m)

	if m.ErrCode == 0 || m.ErrCode == 65201 || m.ErrCode == 65202 {

		w.token = strings.Split(m.ErrMsg, "=")[3]

		fmt.Printf("token:%v\n", w.token)

		w.cookies = resp.Cookies()

		fmt.Println(w.cookies)
		return true
	}

	switch m.ErrCode {
	case -1:
		fmt.Println("系统错误，请稍候再试。")
	case -2:
		fmt.Println("帐号或密码错误。")
	case -3:
		fmt.Println("您输入的帐号或者密码不正确，请重新输入。")
	case -4:
		fmt.Println("不存在该帐户。")
	case -5:
		fmt.Println("您目前处于访问受限状态。")
	case -6:
		fmt.Println("请输入图中的验证码")
	case -7:
		fmt.Println("此帐号已绑定私人微信号，不可用于公众平台登录。")
	case -8:
		fmt.Println("邮箱已存在。")
	case -32:
		fmt.Println("您输入的验证码不正确，请重新输入。")
	case -200:
		fmt.Println("因频繁提交虚假资料，该帐号被拒绝登录。")
	case -94:
		fmt.Println("请使用邮箱登陆。")
	case 10:
		fmt.Println("该公众会议号已经过期，无法再登录使用。")
	case -100:
		fmt.Println("海外帐号请在公众平台海外版登录,<a href=\"http://admin.wechat.com/\">点击登录</a>")
	default:
		fmt.Println("未知的返回。")
	}

	return false
}
func (w *WebWeChat) loginSonar() bool {
	login_url := "http://10.94.10.49:9000/api/authentication/login"
	usernaem := "admin"
	password := "admin"
	//h := md5.New()
	//h.Write([]byte(password))
	//password = hex.EncodeToString(h.Sum(nil))
	fmt.Println(password)
	post_arg := url.Values{"login": {usernaem}, "password": {password}}

	fmt.Println(strings.NewReader(post_arg.Encode()))
	req, err := http.NewRequest("POST", login_url, strings.NewReader(post_arg.Encode()))
	request2, _ := http.NewRequest("GET", "http://10.94.10.49:9000/api/system/health", strings.NewReader(""))
	req.Header.Add("Host", "10.94.10.49:9000")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:67.0) Gecko/20100101 Firefox/67.0")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Referer", "http://10.94.10.49:9000/sessions/new")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Origin", "http://10.94.10.49:9000")
	req.Header.Add("Content-Length", "26")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "_gitlab_session=7c88ba8c4efe13f121b0aa83133c769e; event_filter=all; sidebar_collapsed=false")
	//req.Header.Set("Referer", "https://mp.weixin.qq.com/")

	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	resp, _ := client.Do(req)

	fmt.Printf("++++  %s", resp.Status)
	for i := range resp.Cookies() {
		request2.AddCookie(resp.Cookies()[i])
	}

	resp, err = client.Do(request2)
	if err != nil {
		beego.Error(err)

	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		beego.Error(err)

	}
	beego.Notice("2222222222222", string(body))
	//doc := json.NewDecoder(strings.NewReader(s))
	//
	//type Msg struct {
	//	Ret                     int
	//	ErrMsg                  string
	//	ShowVerifyCode, ErrCode int
	//}
	//
	//var m Msg
	//if err := doc.Decode(&m); err == io.EOF {
	//	fmt.Println(err)
	//} else if err != nil {
	//	log.Println(err)
	//	return false
	//}
	//
	//fmt.Println(m)
	//
	//if m.ErrCode == 0 || m.ErrCode == 65201 || m.ErrCode == 65202 {
	//
	//	w.token = strings.Split(m.ErrMsg, "=")[3]
	//
	//	fmt.Printf("token:%v\n", w.token)
	//
	//	w.cookies = resp.Cookies()
	//
	//	fmt.Println(w.cookies)
	//	return true
	//}
	//
	//switch m.ErrCode {
	//case -1:
	//	fmt.Println("系统错误，请稍候再试。")
	//case -2:
	//	fmt.Println("帐号或密码错误。")
	//case -3:
	//	fmt.Println("您输入的帐号或者密码不正确，请重新输入。")
	//case -4:
	//	fmt.Println("不存在该帐户。")
	//case -5:
	//	fmt.Println("您目前处于访问受限状态。")
	//case -6:
	//	fmt.Println("请输入图中的验证码")
	//case -7:
	//	fmt.Println("此帐号已绑定私人微信号，不可用于公众平台登录。")
	//case -8:
	//	fmt.Println("邮箱已存在。")
	//case -32:
	//	fmt.Println("您输入的验证码不正确，请重新输入。")
	//case -200:
	//	fmt.Println("因频繁提交虚假资料，该帐号被拒绝登录。")
	//case -94:
	//	fmt.Println("请使用邮箱登陆。")
	//case 10:
	//	fmt.Println("该公众会议号已经过期，无法再登录使用。")
	//case -100:
	//	fmt.Println("海外帐号请在公众平台海外版登录,<a href=\"http://admin.wechat.com/\">点击登录</a>")
	//default:
	//	fmt.Println("未知的返回。")
	//}

	return false
}

func (w *WebWeChat) SendTextMsg(fakeid string, content string) bool {
	send_url := "http://mp.weixin.qq.com/cgi-bin/singlesend"
	referer_url := "https://mp.weixin.qq.com/cgi-bin/singlesendpage?t=message/send&action=index&tofakeid=%s&token=%s&lang=zh_CN"

	post_arg := url.Values{
		"tofakeid": {fakeid},
		"type":     {"1"},
		"content":  {content},
		"ajax":     {"1"},
		"token":    {w.token},
		"t":        {"ajax-response"},
	}

	req, _ := http.NewRequest("POST", send_url, strings.NewReader(post_arg.Encode()))

	req.Header.Set("Referer", fmt.Sprintf(referer_url, fakeid, w.token))

	for i := range w.cookies {
		req.AddCookie(w.cookies[i])
	}

	client := new(http.Client)
	resp, _ := client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)

	doc := json.NewDecoder(strings.NewReader(string(data)))

	type Msg struct {
		Ret string
		Msg string
	}

	var m Msg
	if err := doc.Decode(&m); err == io.EOF {
		fmt.Println(err)
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.Msg)

	if m.Msg == "ok" {
		return true
	}

	return false
}

func (w *WebWeChat) GetFakeId() bool {
	msg_url := "https://mp.weixin.qq.com/cgi-bin/contactmanage?t=user/index&pagesize=10&pageidx=0&type=0&groupid=0&token=%s&lang=zh_CN"
	referer_url := "https://mp.weixin.qq.com/cgi-bin/home?t=home/index&lang=zh_CN&token=%s"

	req, _ := http.NewRequest("GET", fmt.Sprintf(msg_url, w.token), nil)

	req.Header.Set("Referer", fmt.Sprintf(referer_url, w.token))

	for i := range w.cookies {
		req.AddCookie(w.cookies[i])
	}

	client := new(http.Client)
	resp, _ := client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(data))
	fmt.Println(string(data))
	re := regexp.MustCompile(`(?s)(?U)contacts.+contacts`)
	list := re.FindString(string(data))
	list = strings.Replace(list, `contacts`, "", -1)
	list = strings.Replace(list, `contacts`, "", -1)
	list = strings.Replace(list, ` `, " ", -1)
	fmt.Println(list)

	list = strings.TrimLeft(list, "\":")
	list = strings.TrimRight(list, "}).")

	fmt.Println(list)

	return true
}

func sonarLogin() {
	conn, err := net.Dial("udp", "10.94.10.49:80")
	if err != nil {
		beego.Error("这台服务器挂掉了", err)
	}
	defer conn.Close()
	intranetIP := conn.LocalAddr().(*net.UDPAddr).IP.String()
	beego.Notice(intranetIP)

}
func PostURL(url string, data []byte) []byte {
	reqData := bytes.NewBuffer(data)
	request, _ := http.NewRequest("POST", url, reqData)
	request2, _ := http.NewRequest("POST", "http://10.94.10.49:9000/api/system/health", reqData)
	request.Header.Set("Content-type", "application/json")

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		beego.Error(err)
		return []byte("")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		beego.Error(err)
		return []byte("")
	}
	fmt.Println("11111111111", response.Status)

	for i := range response.Cookies() {
		beego.Notice(response.Cookies()[i])
		request2.AddCookie(response.Cookies()[i])
	}

	response, err = client.Do(request2)
	if err != nil {
		beego.Error(err)
		return []byte("")
	}
	body, err = ioutil.ReadAll(response.Body)

	if err != nil {
		beego.Error(err)
		return []byte("")
	}
	return body
}

func main() {

	//logininfo := map[string]string{"login":"admin","password":"admin"}
	//logininfoJson, _ := json.Marshal(logininfo)
	//fmt.Println(string(PostURL("http://10.94.10.49:9000/api/authentication/login", logininfoJson)))
	wechat := NewWebWeChat()
	wechat.loginSonar()
	//if wechat.login() == true {
	//	log.Println(wechat.GetFakeId())
	//	tofakeid := "333215495" //my fakeid for test
	//	wechat.SendTextMsg(tofakeid, "Hello Phil.")
	//} else {
	//	fmt.Println("wechat login failed.")
	//}
}

type login struct {
	Login, Password string
}

//func main() {
//	client :=http.DefaultClient
//	logininfo :=login{Login:"admin",Password:"admin"}
//	logininfoJson ,_ := json.Marshal(logininfo)
//	resp ,err :=client.Post("http://10.94.10.49:9000/api/authentication/login",
//		"application/x-www-form-urlencoded",
//		strings.NewReader(string(logininfoJson)))
//	if err!= nil {
//		log.Error(err.Error())
//	}
//	defer resp.Body.Close()
//	//拿出body
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Error(err.Error())
//	}
//	resp ,err = client.Post("http://10.94.10.49:9000/api/authentication/login",
//		"application/x-www-form-urlencoded",
//		strings.NewReader(""))
//	//拿出body
//	body, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Error(err.Error())
//	}
//	fmt.Println(string(body))
//}
//type login struct {
//	login,password string
//}
//func basicAuth(username, password string) string {
//auth := username + ":" + password
//return base64.StdEncoding.EncodeToString([]byte(auth))
//}
//
//func main() {
//c := colly.NewCollector()
//h := http.Header{}
//h.Set("Authorization", "Basic "+basicAuth("admin", "admin"))
//
//c.OnResponse(func(r *colly.Response) {
////fmt.Println(r)
//fmt.Println(string(r.Body))
//})
//
//err := c.Request("POST", "http://10.94.10.49:9000/api/authentication/login", nil, nil, h)
//fmt.Println(err)
//}
