package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func main() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	memStatsbytes, _ := json.Marshal(memStats)
	fmt.Println(string(memStatsbytes))
}
