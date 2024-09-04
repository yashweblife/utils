package utils

import "fmt"

type U struct {
	index int8
}

func (u *U) Log(val string) {
	fmt.Println(u.index, ": ", val)
	u.index++
}
