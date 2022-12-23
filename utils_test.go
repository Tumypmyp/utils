package utils_test

import (
	"fmt"

	"github.com/tumypmyp/utils"
)

func ExampleHead() {
	lines, err := utils.Head(`testdata/somefile`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)

	// Output:
	// [one two]
}
