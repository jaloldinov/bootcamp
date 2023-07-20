package comparetriplets

func CompareTriplets(a []int32, b []int32) []int32 {
	var aScore int32
	var bScore int32
	var arr []int32

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			aScore++
		} else if a[i] < b[i] {
			bScore++
		}
	}
	arr = append(arr, aScore, bScore)
	return arr
}
