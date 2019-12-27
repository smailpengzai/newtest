package main

import "fmt"

func main() {
	// Initialize Jenkins API
	//
	// For example:
	jenkinsApi := Init(&Connection{
		Username:    "dazhu",
		AccessToken: "110f0746378f35de05124e37045d5611ba",
		BaseUrl:     "http://192.168.31.200:8080",
	})
	//http://10.94.10.18:8080/job/test/api/json
	fmt.Println(jenkinsApi.GetJob("ag_dev_apv-app"))

}
