package yr



import (

        "testing"
        "bufio"
	"encoding/csv"
	"os"
	"strings"
	"io"
)


func  TestLesAntallLinjerFil(t *testing.T) {

        type test struct {

        input string

        want int
}



tests := []test{

{input: "kjevik-temp-celsius-20220318-20230318.csv", want:  1656},

}


for _, tc := range tests {

        got := LesAntallLinjerFil(tc.input)

                if got != tc.want {
                        t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
}
}
}

