package arr

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	left := 0
	right := 0
	total := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			right = i
			if right-left-2 >= 0 {
				if flowerbed[left] == 0 {
					total += (right - left - 1) / 2
				} else {
					total += (right - left - 2) / 2
				}
			}
			left = i
		} else {
			if i == len(flowerbed)-1 {
				total += (i - left - 1) / 2
			}
		}
	}
	return total >= n
}
