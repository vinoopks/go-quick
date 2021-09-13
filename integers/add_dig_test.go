package integers

import (
	"testing"
)

func TestAddDig(t *testing.T) {
	sum := Add(3, 4)
	expected := 7

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}
