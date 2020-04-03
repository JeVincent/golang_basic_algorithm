package main

func qsort(a []int)  {
	quicksort(a, 0, len(a)-1)
}

func quicksort(a []int, left int, right int){
	if left >= right{
		return
	}
	p := partition(a, left, right)
	quicksort(a, left, p-1)
	quicksort(a, p+1, right)
}

func partition(a [] int, left int, right int) int{
	l := a[left]
	i, j := -1, 0
	for {
		if (j > right){
			break
		}
		if (a[j] <= l){
			i++
			a[i], a[j] = a[j], a[i]
		}
		j++
	}
	a[left], a[i] = a[i], a[left]
	return i
}


