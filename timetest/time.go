package main

import (
	"fmt"
	"time"
)

func main() {
	tmptime, err := time.ParseInLocation("2006-01-02 15:04:05.000", "2019-10-10 21:21:56.586", time.Local)
	fmt.Println(tmptime, err)
}
