package main


func FindLIS(input []int) []int{
	size:=len(input)
	if size ==0 {
		return input
	}
	maxLength := 1
	P := make([]int, size)
	M := make([]int, size + 1)

	for i := 1; i < size; i++ {
		lo := 1
		hi := maxLength
		for lo <= hi {
			mid := ceil(float64((lo+hi)/2))
			if input[M[mid]] < input[i] {
				lo = mid + 1
			} else {
				hi = mid-1
			}
		}

		P[i] = M[lo-1]
		M[lo] = i

		if lo> maxLength {
			maxLength = lo
		}

	}
	lis := make([]int, maxLength)
	indexOfElementOfLIS := M[maxLength]
	for i := maxLength-1; i >= 0; i--{
		lis[i] = input[indexOfElementOfLIS]
		indexOfElementOfLIS = P[indexOfElementOfLIS]
	}
	return lis
}

func ceil(num float64) int{
	iNum := int(num)
	if num == float64(iNum) {
		return iNum
	}
	return iNum + 1
}
