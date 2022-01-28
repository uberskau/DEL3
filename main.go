package main

import (
	"bufio"
	"fmt"
	"github.com/uberskau/DEL3/event"
	"github.com/uberskau/DEL3/state"
	"os"
)

func main() {

	fmt.Println("########################################################")
	fmt.Println(state.ViewState())
	fmt.Println("Hvem er det som må krysse først: kylling, rev eller korn?")
	fmt.Println("########################################################")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "kylling" {
			fmt.Println("########################################################")
			state.PutInBoat()
			fmt.Println(state.ViewState())

			state.CrossRiver()
			fmt.Println(state.ViewState())

			state.PutOutBoat()
			fmt.Println(state.ViewState())

			fmt.Println("GRATULERER, DET VAR RIKTIG!")
			fmt.Println("########################################################")
		} else if text == "rev" {
			fmt.Println("--------------------------------------------------------")
			fmt.Println(event.GetIn(text))
			fmt.Println(event.GetIn(text))
			fmt.Println(event.CrossRiver(text))
			fmt.Println(event.PutOutBoat(text))
			fmt.Println("")
			fmt.Println("Å NEI, DET VAR VISST IKKE RIKTIG DET!")
			fmt.Println("--------------------------------------------------------")
		} else if text == "korn" {
			fmt.Println("--------------------------------------------------------")
			fmt.Println(event.GetIn(text))
			fmt.Println(event.GetIn(text))
			fmt.Println(event.CrossRiver(text))
			fmt.Println(event.PutOutBoat(text))
			fmt.Println("")
			fmt.Println("Å NEI, DET VAR VISST IKKE RIKTIG DET!")
			fmt.Println("--------------------------------------------------------")
		} else {
			fmt.Println("--------------------------------------------------------")
			fmt.Println("Du må skrive inn en av objektene")
			fmt.Println("--------------------------------------------------------")
		}
	}

}
