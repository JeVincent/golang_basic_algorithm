package main

func ssort(a []int)  {
	l := len(a)
	if l == 0 || l == 1 {
		return
	}

	for i:=0; i<l-1; i++ {
		k := i
		for j:=i+1; j<l; j++ {
			if a[j] < a[k] {
				k = j
			}
		}
		a[i], a[k] = a[k], a[i]
	}
}
