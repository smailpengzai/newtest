package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/sys/menus", sayBye)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":18000", nil))
}

type test struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Id   string `json:"id"`
	Pid  string `json:"pId"`
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	str := []test{{
		Name: "首页",
		Url:  "shouye.html",
		Id:   "1",
		Pid:  "0",
	},
		{
			Name: "数据库",
			Url:  "#",
			Id:   "1",
			Pid:  "0",
		},
		{
			Name: "MYSQL",
			Url:  "mysql.html",
			Id:   "2",
			Pid:  "1",
		},
		{
			Name: "ORACLE",
			Url:  "oracle.html",
			Id:   "3",
			Pid:  "1",
		},
		{
			Name: "开发语言",
			Url:  "#",
			Id:   "4",
			Pid:  "0",
		},
		{
			Name: "JAVA",
			Url:  "java.html",
			Id:   "5",
			Pid:  "4",
		},
		{
			Name: "Python",
			Url:  "python.html",
			Id:   "6",
			Pid:  "4",
		},
	}

	jsonstr, _ := json.Marshal(&str)
	w.Write(jsonstr)
}
