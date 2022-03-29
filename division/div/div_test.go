package div

import "testing"

func TestDivisionOK(t *testing.T) {

	_, error := Division(8, 4)

	if error != nil {
		t.Error()
		return
	}
}

func TestDivisionERROR(t *testing.T) {

	_, error := Division(8, 0)

	if error == nil {
		t.Error()
		return
	}
}
