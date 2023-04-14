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


	fmt.Println("Velg average, convert eller exit for å avslutte programmet")

	if !scanner.Scan() {

		break


	}

	input = scanner.Text()


	switch input {

	case "convert":

	yr.KonverteringAvLinjer()

	fmt.Println("Konvertering ferdig. Sjekk den nye filen som ble opprettet")


	case "exit":

	fmt.Println("Avslutter program")
	return

}
}
}
