package event

import "testing"

func TestGetIn(t *testing.T) {
	wanted := "[--V--I(_hs_" + "" + "_)__________________I--Ø--]"
	got := GetIn("")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	wanted := "[--V--I__________________(_hs_" + "" + "_)I--Ø--]"
	got := CrossRiver("")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestPutOutBoat(t *testing.T) {
	wanted := "[--V--I__________________(_hs_)I--Ø--" + "" + "]"
	got := PutOutBoat("")
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
