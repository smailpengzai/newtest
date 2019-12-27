package main

import (
	"fmt"
	"strings"
)

func main() {
	testCount()

}
func testContains() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("in failure", "s g"))
	fmt.Println(strings.ContainsAny("infailure", "s g"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

func testCount() {
	str := "谷歌中国"
	strdujge := ""
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(len(str))

	for i := 0; i < len(str); i++ {
		fmt.Println(i, "======", str[i])
	}
	for index, each := range str {
		fmt.Println(index, "======", each)
		//fmt.Println( index,"======",str[index])
	}
	for i := 0; i < len(strdujge); i++ {
		fmt.Println(i, "======", strdujge[i])
	}
	fmt.Println(strings.Count(str, strdujge))
}
