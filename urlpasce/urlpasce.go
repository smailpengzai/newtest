package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {

	var inst interface{}
	compelePath := `https:\/\/platform-lookaside.fbsbx.com\/platform\/profilepic\/?psid=3025523350856345\u0026width=1024\u0026ext=1574402870\u0026hash=AeT7pTu3-OCc_sLC`
	//compelePath := "http://lcoalhost:8080/user?id=1"
	//subPath := "/user?id=1"
	fmt.Println(gjson.Unmarshal([]byte(compelePath), &inst))
	fmt.Println(inst)

	// 双斜杠
	//cP, _ := url.Parse(compelePath)
	//fmt.Println(cP.Host)
	// >>> lcoalhost:8080

	// 非双斜杠
	//sP, _ := url.Parse(subPath)
	//_log(sP.Host)
	//_log(sP.RawQuery)
	// >>>   （空值）
	// >>> id=1

	// 两者可获取的参数不同

}
