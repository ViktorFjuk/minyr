package yr



import (

        "testing"
        //"bufio"
	//"encoding/csv"
	//"os"
	//"strings"
	//"io"
)


func  TestLesAntallLinjerFil(t *testing.T) {

        type test struct {

        input string

        want int
}



tests := []test{

{input: "/home/minyr/kjevik-temp-celsius-20220318-20230318.csv", want:  16756},

}


for _, tc := range tests {

        got := LesAntallLinjerFil(tc.input)

                if got != tc.want {
                        t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
}
}
}

func TestKonverteringAvLinjer(t *testing.T) {

	type test struct {

	input string

	want string
}

tests := []test{

{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;", want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Viktor Fjuk"},

}


for _, tc := range tests {

got := ProsesserLinjer(tc.input)

if !(tc.want == got) {

	t.Errorf("expected: %v, got: %v", tc.want, got)

}
}
}


//Siste manglende testen
//func test gjennomsnitt 
