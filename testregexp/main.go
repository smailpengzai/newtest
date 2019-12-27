package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := " th-uat_rfs-app_郑强_SQL变更"
	reg := regexp.MustCompile(".*th-uat_rfs-app.*SQL变更.*")
	fmt.Println(reg.MatchString(str))
}
