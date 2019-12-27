package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func main() {
	//创建一个链表
	l := list.New()

	//链表最后插入元素
	a1 := l.PushBack(1)
	b2 := l.PushBack(2)

	//链表头部插入元素
	l.PushFront(3)
	l.PushFront(4)

	printList(l)

	//取第一个元素
	f := l.Front()
	fmt.Println(f.Value)

	//取最后一个元素
	b := l.Back()
	fmt.Println(b.Value)

	//获取链表长度
	fmt.Println(l.Len())

	//在某元素之后插入
	l.InsertAfter(66, a1)

	//在某元素之前插入
	l.InsertBefore(88, a1)

	printList(l)

	l2 := list.New()
	l2.PushBack(11)
	l2.PushBack(22)
	//链表最后插入新链表
	l.PushBackList(l2)
	printList(l)

	//链表头部插入新链表
	l.PushFrontList(l2)
	printList(l)

	//移动元素到最后
	l.MoveToBack(a1)
	printList(l)

	//移动元素到头部
	l.MoveToFront(a1)
	printList(l)

	//移动元素在某元素之后
	l.MoveAfter(b2, a1)
	printList(l)

	//移动元素在某元素之前
	l.MoveBefore(b2, a1)
	printList(l)

	//删除某元素
	l.Remove(a1)
	printList(l)
}
