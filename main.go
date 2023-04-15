package main


import (

	"fmt"
	"github.com/ViktorFjuk/minyr/yr"
	"os"
	"bufio"
)


func main() {

var input string

scanner := bufio.NewScanner(os.Stdin)



for {

	fmt.Println("Velg average, convert eller exit")

	if !scanner.Scan() {

		break

	}

	input = scanner.Text()

	switch input {

	case "convert":

	yr.KonverteringAvLinjer()

	fmt.Println("Konvertering ferdig. Sjekk den nye filen som ble opprettet")

	case "average":

	fmt.Println("Velg enhet for gjennomsnittstemperatur (c/f):")

	if !scanner.Scan() {

	break

	}

	enhet := scanner.Text()

	switch enhet {

	case "c":

	yr.GjennomsnittsBeregningCelsius()

	case "f":

	yr.GjennomsnittsBeregningFahr()

	default:

	fmt.Println("Ugyldig enhet valgt. Velg c eller f.")

}

 	case "exit":

        fmt.Println("Avslutter program")
        return
}
}
}
