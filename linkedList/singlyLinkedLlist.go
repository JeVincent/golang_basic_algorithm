package linkedList

import (
	"errors"
	"fmt"
)

type SinglyNode struct {
	Val  interface{}
	Next *SinglyNode
}

type SinglyList struct {
	head *SinglyNode
}

// Whether singly linked list is empty
func (this *SinglyList) IsEmpty() bool {
	if this.head == nil {
		return true
	}
	return false
}

// Return the length of the singly linked list
func (this *SinglyList) Length() int {
	l := 0
	for n := this.head; n != nil; n = n.Next {
		l++
	}
	return l
}

// Add elements from the head
func (this *SinglyList) OfferFirst(val interface{}) *SinglyNode {
	node := &SinglyNode{
		Val:  val,
		Next: this.head,
	}
	this.head = node
	return node
}

// Add elements from the tail
func (this *SinglyList) OfferLast(val interface{}) *SinglyNode {
	node := &SinglyNode{
		Val: val,
	}
	if this.head == nil {
		this.head = node
	} else {
		n := this.head
		for n.Next != nil {
			n = n.Next
		}
		n.Next = node
	}
	return this.head
}

// Add element at specified position
func (this *SinglyList) Insert(val interface{}, index int) (*SinglyNode, error) {
	l := this.Length()
	if (index > 0 && index > l) || (index < 0 && index < -(l+1)) {
		return nil, errors.New("out ouf the range")
	}
	if index < 0 {
		index += l + 1
	}
	if index == 0 {
		this.OfferFirst(val)
	} else if index == l {
		this.OfferLast(val)
	} else {
		tmp := this.head;
		for index > 1 {
			tmp = tmp.Next
			index--
		}
		n := &SinglyNode{
			Val: val,
		}
		n.Next = tmp.Next
		tmp.Next = n
	}
	return this.head, nil
}

// Delete the specified element
func (this *SinglyList) Remove(val int) *SinglyNode {
	tmp := this.head
	if tmp == nil {
		return tmp
	}
	if tmp.Val == val {
		this.head = tmp.Next
		tmp.Next = nil
		return this.head
	}
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return this.head
}

// Delete the element at the specified position
func (this *SinglyList) RemoveAtIndex(index int) (*SinglyNode, error) {
	l := this.Length()
	if (index > 0 && index >= l) || (index < 0 && index < -l) {
		return nil, errors.New("out ouf the range")
	}
	if index < 0 {
		index += l
	}
	if index == 0 {
		this.head = this.head.Next
	} else {
		tmp := this.head;
		for index > 1 {
			tmp = tmp.Next
			index--
		}
		tmp.Next = tmp.Next.Next
	}
	return this.head, nil
}

// See if it contains an element
func (this *SinglyList) Contains(val int) bool {
	tmp := this.head
	for tmp != nil {
		if tmp.Val == val {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

// Iterate over all elements
func (this *SinglyList) Iterates() {
	fmt.Print("[")
	tmp := this.head
	for tmp != nil {
		fmt.Print(tmp.Val)
		if tmp.Next != nil {
			fmt.Print(", ")
		}
		tmp = tmp.Next
	}
	fmt.Println("]")
}
