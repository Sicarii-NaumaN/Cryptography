package main

import (
	"flag"
	"fmt"
	"log"
	"simmetric/algo"
)

func main() {
	args := &algo.Parameters{}
	var err error
	if args, err = algo.ParseFlags(); err != nil {
		flag.Usage()
	}

	data := make([]algo.Visitor, args.NumBilets)

	names, err := algo.ReadFile(args.FileName)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < args.NumBilets; i++ { // Initialization of structure
		data[i].Name = names[i]
		data[i].Bilet = i + 1
	}

	// XOR-randomizer
	for j := 0; j < 1001; j++ { // Swap the tickets
		for i := 0; i < args.NumBilets-1; i++ {
			algo.Swap(&data[(args.Parameter^i+12&j)%args.NumBilets].Bilet,
				&data[(args.Parameter^j+12&i)%args.NumBilets].Bilet)

		}
	}

	for _, visitor := range data {
		fmt.Printf("%s: %d\n", visitor.Name, visitor.Bilet)
	}
}
