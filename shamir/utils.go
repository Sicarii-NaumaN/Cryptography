package crypto

import (
	"errors"
	"flag"
	"io/ioutil"
	"strings"
)

type Parameters struct {
	Split  	 bool
	Recover  bool
	Test     bool
}

// ParseFlags parses passed args in program
func ParseFlags() (*Parameters, error) {
	parsedFlags := Parameters{}
	flag.BoolVar(&parsedFlags.Split, "split", false, "[-split]")
	flag.BoolVar(&parsedFlags.Recover, "recover", false, "[-recover]")
	flag.BoolVar(&parsedFlags.Test, "test", false, "[-test]")

	flag.Parse()

	if (parsedFlags.Split && parsedFlags.Recover) ||
		(!parsedFlags.Split && !parsedFlags.Recover && !parsedFlags.Test) {
		return nil, errors.New("Unexpected usage")
	}

	return &parsedFlags, nil
}

func ReadFile(name string) ([]string, error) {
	stringData, err := ioutil.ReadFile(name)
	return strings.Split(string(stringData), "\n"), err
}

func Swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
