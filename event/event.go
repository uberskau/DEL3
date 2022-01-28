package event

var event string

func Tilstander(tilstanden string) string {
	event = tilstanden
	return tilstanden
}

func GetIn(item string) string {
	event = "[--V--I(_hs_" + item + "_)__________________I--Ø--]"
	return Tilstander(event)
}

func CrossRiver(item string) string {
	event = "[--V--I__________________(_hs_" + item + "_)I--Ø--]"
	return Tilstander(event)
}

func PutOutBoat(item string) string {
	event = "[--V--I__________________(_hs_)I--Ø--" + item + "]"
	return Tilstander(event)
}
