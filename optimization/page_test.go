package optimization

import (
	"log"
	"testing"
)

func Test__page_One(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16}

	start, end := Page.One(2, 40, len(nums))

	log.Println("result = ", start, end, "nums = ", nums[start:end])
}
