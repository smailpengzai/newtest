package main

import (
	"container/ring"
	"fmt"
)

func printRing(r *ring.Ring) {
	r.Do(func(v interface{}) {
		fmt.Print(v.(int), " ")
	})
	fmt.Println()
}

func main() {
	//创建环形链表
	r := ring.New(5)
	//循环赋值
	for i := 0; i < 5; i++ {
		r.Value = i
		//取得下一个元素
		r = r.Next()
	}
	fmt.Println("第一次打印环的数据")
	printRing(r)
	//环的长度
	fmt.Println("第一次打印环的长度")
	fmt.Println(r.Len())

	//移动环的指针
	r.Move(2)
	fmt.Println("移动环的指针到2 打印数据")
	printRing(r)
	//从当前指针删除n个元素
	fmt.Println("从指针2删除两个元素打印数据")
	r.Unlink(2)
	printRing(r)

	//连接两个环
	r2 := ring.New(3)
	for i := 0; i < 3; i++ {
		r2.Value = i + 10
		//取得下一个元素
		r2 = r2.Next()
	}
	fmt.Println("打印第二个环的数据")
	printRing(r2)
	fmt.Println("链接俩个环并打印数据")
	r.Link(r2)
	printRing(r)
}
