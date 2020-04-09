package main

import (
	"fmt"
	"golang_basic_algorithm/linkedList"
)

func main(){
	doublyLinkedListTest()
}
func doublyLinkedListTest()  {
	list := linkedList.DoublyLinkedList{}
	// 1
	list.OfferLast(1)
	list.Iterates()
	// 0, 1
	list.OfferFirst(0)
	list.Iterates()
	// 0, 1, 2
	list.OfferLast(2)
	list.Iterates()
	// 3
	fmt.Println(list.Length())
	// true, false
	fmt.Println(list.Contains(2))
	fmt.Println(list.Contains(3))
	// 0, 3, 1, 2
	err := list.Insert(1, 3)
	if err == nil {
		list.Iterates()
	}
	// 0, 3, 1 ,4, 2
	err = list.Insert(3, 4)
	if err == nil {
		list.Iterates()
	}
	// 0, 3, 1, 4, 2
	// error
	err = list.Insert(6, 6)
	if err != nil {
		fmt.Println(err)
	}
	list.Iterates()
	// 0, 1, 4, 2
	list.Remove(3)
	list.Iterates()
	fmt.Println(list.PeekLast())
	// 0, 1, 4
	list.RemoveAtIndex(3)
	list.Iterates()
	// 0, 1
	list.PollLast()
	list.Iterates()
	// 1
	list.PollFirst()
	fmt.Println(list.PeekFirst())
	list.Iterates()
	// []
	list.PollLast()
	list.Iterates()
	// true
	fmt.Println(list.IsEmpty())
}
