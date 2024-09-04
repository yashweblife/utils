package utils

import "fmt"

type U struct {
	index int8
}

func (u *U) Log(val string) {
	fmt.Printf("%d: %s\n", u.index, val)
	u.index++
}
