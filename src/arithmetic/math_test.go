package arithmetic

import "testing"

func TestAdd(t *testing.T) {

	got := Add(2, 3)
	if got != 5 {
		t.Errorf(" 2 + 3 did not equal 5! Received %d", got)
	}

}

func TestSub(t *testing.T) {

	if Sub(5, 3) != 2 {
		t.Error(" 5 - 3 did not equal 2!")
	}

}
