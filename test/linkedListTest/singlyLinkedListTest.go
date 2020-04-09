package main

import (
	"fmt"
	"golang_basic_algorithm/linkedList"
)

//func main(){
//	singlyLinkedListTest()
//}
func singlyLinkedListTest()  {
	list := linkedList.SinglyList{}
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
	_, err := list.Insert(3, 1)
	if err == nil {
		list.Iterates()
	}
	// 0, 3, 1 ,4, 2
	_, err = list.Insert(4, -2)
	if err == nil {
		list.Iterates()
	}
	// 0, 3, 1, 4, 2
	// error
	_, err = list.Insert(5, 6)
	if err != nil {
		list.Iterates()
		fmt.Println(err)
	}
	// 0, 1, 4, 2
	list.Remove(3)
	list.Iterates()
	// 0, 1, 4
	list.RemoveAtIndex(-1)
	list.Iterates()
	//[]
	list.RemoveAtIndex(2)
	list.RemoveAtIndex(1)
	list.RemoveAtIndex(0)
	list.Iterates()
	// true
	fmt.Println(list.IsEmpty())
}
