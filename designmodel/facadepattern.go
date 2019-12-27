package main

import (
	"fmt"
)

type DataHandler struct {
	decoder *Decoder     //解码子系统
	handler *MainHandler //处理子系统
	encoder *Encoder     //编码子系统
}

func (pD *DataHandler) Working() {
	pD.decoder.Working()
	pD.handler.Working()
	pD.encoder.Working()
}

type Decoder struct {
}

func (pD *Decoder) Working() {
	fmt.Println("解码子系统先XXX格式数据,并转换为xxxx结构数据")
}

type MainHandler struct {
}

func (pM *MainHandler) Working() {
	fmt.Println("数据处理子系统，处理数据")
}

type Encoder struct {
}

func (pE *Encoder) Working() {
	fmt.Println("编码子系统将xxx数据格式转换为Json格式")
}

func main() {

	worker := &DataHandler{decoder: &Decoder{}, handler: &MainHandler{}, encoder: &Encoder{}}
	worker.Working()
	return
}
