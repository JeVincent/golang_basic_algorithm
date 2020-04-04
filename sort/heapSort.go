package main

func hsort(a []int) {
	l := len(a)
	if l==0 || l==1 {
		return
	}

	for k:=l/2; k>=1; k-- {
		sink(a, k, l)
	}
	for i:=l; i>1; i-- {
		a[0], a[i-1] = a[i-1], a[0]
		sink(a, 1, i-1)
	}
}

func sink(a []int, k, l int)  {
	// k, left and right is the index of the heap
	// tmp is the index of slice
	for {
		left := 2 * k
		right := left + 1
		if left > l {
			break
		}
		tmp := k - 1
		if a[left-1] > a[tmp] {
			tmp = left - 1
		}
		if right <= l && a[right-1] > a[tmp] {
			tmp = right - 1
		}
		if tmp == k - 1 {
			break
		}
		a[k-1], a[tmp] = a[tmp], a[k-1]
		k = tmp + 1
	}
}
