package state

var state = "[kylling rev korn hs --V--I(__)__________________I--Ø--]"

func ViewState() string {
	return state
}

func PutInBoat() string {
	state = "[rev korn --V--I(_hs_kylling_)__________________I--Ø--]"
	return state
}

func CrossRiver() string {
	state = "[rev korn --V--I__________________(_hs_kylling_)I--Ø--]"
	return state
}

func PutOutBoat() string {
	state = "[rev korn --V--I__________________(_hs_)I--Ø-- kylling]"
	return state
}
