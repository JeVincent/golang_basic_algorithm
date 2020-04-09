package linkedList

import (
	"errors"
	"fmt"
)

type DoublyNode struct {
	Val       interface{}
	Pre, Next *DoublyNode
}

type DoublyLinkedList struct {
	head, tail *DoublyNode
}

func (this *DoublyLinkedList) IsEmpty() bool {
	return this.head == nil
}

func (this *DoublyLinkedList) Length() int {
	if this.head == nil {
		return 0
	}
	l := 1
	for h := this.head; h != this.tail; {
		l++
		h = h.Next
	}
	return l
}

func (this *DoublyLinkedList) OfferLast(val interface{}) {
	node := &DoublyNode{
		Val: val,
	}
	if this.head == nil {
		this.head = node
		this.tail = this.head
		return
	}
	this.tail.Next = node
	node.Pre = this.tail
	this.tail = node
}

func (this *DoublyLinkedList) OfferFirst(val interface{}) {
	node := &DoublyNode{
		Val: val,
	}
	if this.head == nil {
		this.head = node
		this.tail = this.head
		return
	}
	this.head.Pre = node
	node.Next = this.head
	this.head = node
}

func (this *DoublyLinkedList) PeekFirst() interface{} {
	if this.head == nil {
		return nil
	}
	return this.head.Val
}

func (this *DoublyLinkedList) PeekLast() interface{} {
	if this.tail == nil {
		return nil
	}
	return this.tail.Val
}

func (this *DoublyLinkedList) PollFirst() interface{} {
	if this.head == nil {
		return nil
	}
	v := this.head.Val
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.Next
		this.head.Pre = nil
	}
	return v
}

func (this *DoublyLinkedList) PollLast() interface{} {
	if this.tail == nil {
		return nil
	}
	v := this.tail.Val
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.tail = this.tail.Pre
		this.tail.Next = nil
	}
	return v
}

func (this *DoublyLinkedList) Insert(index int, val interface{}) error {
	if index < 0 || index > this.Length() {
		return errors.New("out of range")
	}
	if index == 0 {
		this.OfferFirst(val)
	} else if index == this.Length() {
		this.OfferLast(val)
	} else {
		node := &DoublyNode{
			Val: val,
		}
		h := this.head
		for index > 1 {
			h = h.Next
			index--
		}
		node.Next = h.Next
		node.Pre = h
		h.Next = node
		node.Next.Pre = node
	}
	return nil
}

func (this *DoublyLinkedList) Remove(val interface{}) {
	if this.head.Val == val {
		this.PollFirst()
	}
	h := this.head
	for h != nil {
		if h == this.tail {
			this.PollLast()
			break
		} else if h.Val == val {
			h.Pre.Next = h.Next
			h.Next.Pre = h.Pre
			break
		} else {
			h = h.Next
		}
	}
}

func (this *DoublyLinkedList) RemoveAtIndex(index int) {
	if index < 0 || index >= this.Length() {
		return
	}
	if index == 0 {
		this.PollFirst()
	} else if index == this.Length()-1 {
		this.PollLast()
	} else {
		h := this.head
		for index > 0 {
			h = h.Next
			index--
		}
		h.Pre.Next = h.Next
		h.Next.Pre = h.Pre
	}
}

func (this *DoublyLinkedList)Contains(val interface{}) bool{
	if this.head == nil {
		return false
	}
	h := this.head
	for h != nil {
		if h.Val == val {
			return true
		}
		h = h.Next
	}
	return false
}

func (this *DoublyLinkedList) Iterates() {
	fmt.Print("[")
	h := this.head
	for h != nil {
		fmt.Print(h.Val)
		if h.Next != nil {
			fmt.Print(", ")
		}
		h = h.Next
	}
	fmt.Println("]")
}
