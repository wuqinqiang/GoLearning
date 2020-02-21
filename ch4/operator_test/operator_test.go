package operator_test

import (
	"fmt"
	"testing"
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	fmt.Println(a == b)

}
