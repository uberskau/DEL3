package state

import "testing"

func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn hs --V--I(__)__________________I--Ø--]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestPutInBoat(t *testing.T) {
	wanted := "[rev korn --V--I(_hs_kylling_)__________________I--Ø--]"
	state := PutInBoat()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	wanted := "[rev korn --V--I__________________(_hs_kylling_)I--Ø--]"
	state := CrossRiver()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestPutOutBoat(t *testing.T) {
	wanted := "[rev korn --V--I__________________(_hs_)I--Ø-- kylling]"
	state := PutOutBoat()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
