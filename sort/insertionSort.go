package sort

func Isort(a []int)  {
	l := len(a)
	if l==0 || l==1 {
		return
	}
	for i:=1; i<l; i++ {
		tmp := a[i]
		for j:=i-1; j>=0 && tmp <= a[j]; j-- {
			a[j+1], a[j] = a[j], a[j+1]
		}
	}
}
