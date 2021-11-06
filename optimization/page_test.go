package optimization

import (
	"log"
	"testing"
)

func Test__page_One(t *testing.T) {
	nums := []int{}

	start, end := Page.One(2, 40, len(nums))

	log.Println("result = ", start, end, "nums = ", nums[start:end])
}
