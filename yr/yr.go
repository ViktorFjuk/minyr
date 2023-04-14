package yr


import (

        "fmt"
        "strconv"
        "strings"
        //"errors"
        "log"
        "os"
        //"io"
        "bufio"
        "github.com/ViktorFjuk/funtemps/conv"
	//"encoding/csv"
)


func LesAntallLinjerFil(filename string) int{

src, err := os.Open("/home/minyr/kjevik-temp-celsius-20220318-20230318.csv")

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

//Apner input filen


inputFil, err := os.Open("/home/minyr/kjevik-temp-celsius-20220318-20230318.csv")

if err != nil {

	log.Fatal(err)

}
defer inputFil.Close()


//Lager output filen

outputFil, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")

if err != nil {
	log.Fatal(err)

}

defer outputFil.Close()


outputWriter := bufio.NewWriter(outputFil)


inputScanner := bufio.NewScanner(inputFil)



for inputScanner.Scan() {

linje := inputScanner.Text()


konvertertLinje := ProsesserLinjer(linje)


_, err := outputWriter.WriteString(konvertertLinje + "\n")

if err != nil {

	log.Fatal(err)

}

}

err = outputWriter.Flush()

if err != nil {

	log.Fatal(err)
}

}



func ProsesserLinjer(linje string) string {


if strings.Contains(linje, "Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur") {

return linje

}



if strings.Contains(linje, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;") {

linje = strings.Replace(linje, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;", "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Viktor Fjuk", 1)

return linje

}

elementer := strings.Split(linje, ";")

celsiusStr := elementer[len(elementer)-1]

celsius, err := strconv.ParseFloat(celsiusStr, 64)

if err != nil {

log.Fatal(err)

}


fahrenheit := conv.CelsiusToFahrenheit(celsius)

elementer[len(elementer)-1] = fmt.Sprintf("%.2f", fahrenheit)


konvertertLinje := strings.Join(elementer, ";")

return konvertertLinje

}





