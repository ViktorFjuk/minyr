package yr


import (

        "fmt"
        "strconv"
        "strings"
        "errors"
        "log"
        "os"
        "io"
        "bufio"
        "github.com/ViktorFjuk/funtemps/conv"
	"encoding/csv"
)


func LesAntallLinjerFil(filename string) int{

src, err := os.Open("/home/minyr/yr/kjevik-temp-celsius-20220318-20230318.csv")

if err != nil {

        log.Fatal(err)

}

defer src.Close()

scanner := bufio.NewScanner(src)

antall := 0

for scanner.Scan() {

antall++

}


return antall

}

func KonverteringAvLinjer() {

inputFil, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")

if err != nil {

	fmt.Println("error", err)

}
defer inputFil.Close()

inputScanner := bufio.NewScanner(inputFil)

outputFile, err := os.Create("kjevik-temp-fahr-20220318-20230319.csv")

if err !=nil {
	fmt.Println("error", err)

