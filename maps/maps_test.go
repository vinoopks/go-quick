package maps

import "testing"

func TestMaps(t *testing.T) {
	v := PrintMaps()
	PrintMaps()
	if v != "vin" {
		t.Errorf("issue")
	}
}
