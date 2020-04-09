package sort

func Shsort(a []int) {
	l := len(a)
	if l==0 || l == 1 {
		return
	}
	for step:=l/2; step>0; step/=2 {
		for i:=step; i<l; i++ {
			tmp := a[i]
			for j:=i-step; j>=0 && tmp<a[j]; j-=step {
				a[j], a[j+step] = a[j+step], a[j]
			}
		}
	}
}
