package main

import "fmt"

func main() {
	secret := NewGoogleAuth().GetSecret()
	fmt.Println(secret)
	//oldcode := ""
	//codecount := 0
	//
	//for {
	//	secret = "FZEHUTIF5J3RJSFDVAMU6K7K475OAFYD"
	//	code, err := NewGoogleAuth().GetCode(secret)
	//
	//	fmt.Println(secret, code, err)
	//	time.Sleep(1 * time.Second)
	//	codecount++
	//	if codecount == 1 {
	//		oldcode = code
	//	} else {
	//		if oldcode != code {
	//			fmt.Println(codecount, secret, oldcode, code, err)
	//			codecount =0
	//		}
	//	}
	//}
}
