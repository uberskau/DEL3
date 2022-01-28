package state

var state = "[kylling rev korn hs --V--I(__)__________________I--Ø--]"

func ViewState() string {
	return state
}

func PutInBoat() {
	state = "[rev korn --V--I(_hs_kylling_)__________________I--Ø--]"
}

func CrossRiver() {
	state = "[rev korn --V--I__________________(_hs_kylling_)I--Ø--]"
}

func PutOutBoat() {
	state = "[rev korn --V--I__________________(_hs_)I--Ø-- kylling]"
}
