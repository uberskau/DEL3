package event

import "testing"

func TestGetIn(t *testing.T) {
	// Hva forventer jeg?
	wanted := "[--V--I(_hs_" + item + ")__________________I--Ø--]"
	got := Put("korn") // Hva fikk jeg?
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	// Hva forventer jeg?
	wanted := "[--V--I__________________(_hs_" + item + "_)I--Ø--]"
	got := Put("korn") // Hva fikk jeg?
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}

func TestPutOutBoat(t *testing.T) {
	// Hva forventer jeg?
	wanted := "[--V--I__________________(_hs_)I--Ø--" + item + "]"
	got := Put("korn") // Hva fikk jeg?
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
	}
}
