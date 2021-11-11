func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(arr []int)[]int{
	if len(arr)>1{
		mid := len(arr)/2
		l := mergeSort(arr[:mid])
		r := mergeSort(arr[mid:])
		return merge(l,r)
	}
	return arr
}
func merge(a,b []int)[]int{
	tmp := []int{}
	for len(a)>0 || len(b)>0 {
		if len(a)==0{
			tmp=append(tmp,b...)
			return tmp
		}
		if len(b)==0{
			tmp=append(tmp,a...)
			return tmp
		}
		if a[0]<=b[0]{
			tmp=append(tmp,a[0])
			a=a[1:]
		}else{
			tmp=append(tmp,b[0])
			b=b[1:]
		}
	}
	return tmp
}
