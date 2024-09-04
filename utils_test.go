package utils

import "testing"

func TestLog(t *testing.T) {
	u := U{Index: 0}
	u.Log("Hello World")

	if u.Index != 1 {
		t.Fail()
	}
}
