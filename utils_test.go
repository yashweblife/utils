package utils

import (
	"testing"
)

func TestLog(t *testing.T) {
	u := U{}
	u.Log("hello")
	u.Log("world")

	if u.index != 2 {
		t.Errorf("expected 2, got %d", u.index)
	}
}
