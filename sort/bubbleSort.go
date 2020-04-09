package sort

func Bsort(a []int)  {
	l := len(a)
	if l == 0 || l == 1 {
		return
	}
	flag := true
	for i:=0; i<l-1; i++ {
		flag = true
		for j:=0; j<l-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
