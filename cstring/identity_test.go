package cstring

import (
	"fmt"
	"github.com/icarus-go/utils/cstring/constant"
	"github.com/icarus-go/utils/cstring/identity"
	"regexp"
	"testing"
)

func Test_Identity_Substring(t *testing.T) {

	valid := identity.IsValid("D5296952")

	fmt.Printf("%v", valid)
	fmt.Printf("%v", len("D5296952"))
}

func Test_Identity_birthday(t *testing.T) {
	r := regexp.MustCompile(constant.Birthday.Value())
	s := "19821026"

	strings := r.FindAllStringSubmatch(s, -1)

	fmt.Printf("%v", strings[0])

	fmt.Printf("%d", len(strings[0]))
}
