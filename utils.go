package utils

import "fmt"

type U struct {
	Index int8
}

func (u *U) Log(val string) {
	fmt.Println(u.Index, val)
	u.Index++
}
