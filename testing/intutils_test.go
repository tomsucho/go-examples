package intutils

import "testing"

func TestIntMin(t *testing.T) {
	a, b, wants := 2, -2, -2
	res := IntMin(a, b)
	if res != wants {
		t.Errorf("IntMin(%d,%d) wants %d", a, b, wants)
	}
}
