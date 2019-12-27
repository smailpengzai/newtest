/*
go 解析未知的json数据使用go-simplejson
*/
package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	body := `{
  "version": {
      "max": 3,
      "last": "2016-03-11",
      "detail": [
          {
              "time": "2016-03-12",
              "ops": "add my email"
          }
         ]
      }
  }
  `
	js, err := simplejson.NewJson([]byte(body)) //反序列化
	if err != nil {
		panic(err.Error())
	}

	body = `{
  "version": {
      "max": 3,
      "last": "2016-03-11",
      "detail": [
          {
              "time": "2016-03-12",
              "ops": "add my email"
          }
         ]
      },
  "version": {
      "max": 3,
      "last": "2016-03-11",
      "detail": [
          {
              "time": "2016-03-12",
              "ops": "add my email"
          }
         ]
      }
  }
  `
	err = js.UnmarshalJSON([]byte(body))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(js)
}
