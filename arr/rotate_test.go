package arr

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a)
}
