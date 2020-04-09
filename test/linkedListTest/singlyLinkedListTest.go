package main

import "golang_basic_algorithm/linkedList"

func main(){
	singlyLinkedListTest()
}


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
}
