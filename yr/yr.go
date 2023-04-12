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

)


func LesAntallLinjerFil(filename string) int{

src, err := os.Open("/home/ViktorFjuk/minyr/yr/kjevik-temp-celsius-20220318-20230318.csv")

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

