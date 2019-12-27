package main

import (
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"time"
)

type datamodel struct {
	Name  string
	Age   int
	Sex   string
	Money float64
}

func main() {
	var datas []datamodel
	for i := 0; i < 10000; i++ {
		datas = append(datas, datamodel{Name: "小可爱", Age: 19, Sex: "wuman", Money: 29461.251})
	}
	jsonstart := time.Now()
	json.Marshal(datas)
	fmt.Println("json 耗时：", time.Now().Sub(jsonstart).Seconds())
	msgpackstart := time.Now()
	msgpack.Marshal(datas)
	fmt.Println("msgpack 耗时：", time.Now().Sub(msgpackstart).Seconds())
}
