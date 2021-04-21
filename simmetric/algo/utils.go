package algo

import (
	"errors"
	"flag"
	"io/ioutil"
	"strings"
)

type Visitor struct {
	Name  string
	Bilet int
}


type Parameters struct {
	FileName  string
	NumBilets int
	Parameter int
}

// ParseFlags parses passed args in program
func ParseFlags() (*Parameters, error) {
	parsedFlags := Parameters{}
	flag.StringVar(&parsedFlags.FileName, "file", "", "[--file]")
	flag.IntVar(&parsedFlags.NumBilets, "numbilets", -1, "[--numbilets]")
	flag.IntVar(&parsedFlags.Parameter, "parameter", -1, "[--parameter]")

	flag.Parse()

	if parsedFlags.FileName == "" || parsedFlags.Parameter == -1 || parsedFlags.NumBilets == -1 {
		return nil, errors.New("Unexpected usage")
	}

	return &parsedFlags, nil
}

func ReadFile(name string) ([]string, error) {
	stringData, err := ioutil.ReadFile(name)
	return strings.Split(string(stringData), "\n"), err
}

func Swap (x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}