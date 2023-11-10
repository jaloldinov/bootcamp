package candles

import (
	"fmt"
	"sort"
)

func BirthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	sort.Slice(candles, func(i, j int) bool {
		return candles[i] > candles[j]
	})

	fmt.Println(candles)

	var count int32 = 1

	for i := 1; i < len(candles); i++ {
		if candles[i] == candles[0] {
			count++
		} else {
			break
		}
	}
	return count
}
