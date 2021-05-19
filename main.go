package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	mrand "math/rand"
	shamir "shamir/shamir"
)

func main() {
	args := &shamir.Parameters{}
	var err error

	if args, err = shamir.ParseFlags(); err != nil {
		flag.Usage()
		return
	}

	// If flag -test is passed
	if args.Test {
		// Any private key
		pkBytes, err := hex.DecodeString("a11b0a4e1a132305652ee7a8eb7848f6ad" +
			"5ea381e3ce20a2c086a2e388230811")
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		fmt.Printf("Passed key:\n%x\n", pkBytes)

		fmt.Println("-------------------------------------")

		// The number of people who need to give out the sharing key
		keyShares := 3

		// The number of people needed to get the initial private key
		keyThreshold := 2

		// Spliting input key to keyShares sharing keys
		keys := Split(pkBytes, keyThreshold, keyShares)

		stringKeys := make([]string, 0)
		for i := range keys {
			stringKeys = append(stringKeys, fmt.Sprintf("%x", keys[i]))
		}
		fmt.Println("keys:")
		for i := range stringKeys {
			fmt.Println(stringKeys[i])
		}

		// Generating graph points of a function in a coordinate system
		// These points uses to restore polynomial
		points := PointsGeneration(len(keys), keys)

		fmt.Println("-------------------------------------")
		// Getting the result
		fmt.Printf("%x <---- This is your key", Recover(points))
		return
	}

	// If flag -split is passed
	if args.Split {
		var secret string
		var keyShares, keyThreshold int
		fmt.Println("Enter secret...")
		fmt.Scan(&secret)
		fmt.Println("Enter N and T...", )
		fmt.Scan(&keyShares, &keyThreshold)

		keys := Split([]byte(secret), keyThreshold, keyShares)

		stringKeys := make([]string, 0)
		for i := range keys {
			stringKeys = append(stringKeys, fmt.Sprintf("%x", keys[i]))
		}
		fmt.Println("keys:")
		for i := range stringKeys {
			fmt.Println(stringKeys[i])
		}
		return
	}

	// If flag -recover is passed
	if args.Recover {
		keyParts := make([][]byte, 0)
		var sharingKey string


		fmt.Println("Enter sharing keys...")
		for {
			_, err = fmt.Scan(&sharingKey)
			if err != nil {
				break
			}
			pkBytes, _ := hex.DecodeString(sharingKey)
			keyParts = append(keyParts, pkBytes)
		}

		points := PointsGeneration(len(keyParts), keyParts)
		fmt.Printf("%x <---- This is your key", Recover(points))
		return
	}
}

func Split(secret []byte, keyThreshold, keyShares int) [][]byte {
	keys, err := shamir.GenKeyShares(secret, keyThreshold, keyShares)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return keys
}

func Recover(points []shamir.Point) string {
	derivedKey, err := shamir.GetKeyFromKeyShares(points)
	if err != nil {
		log.Fatalf(err.Error())
	}
	decodedKey := string(derivedKey[:])

	return decodedKey
}

func PointsGeneration(numOfKeys int, keys [][]byte) []shamir.Point {
	pointsMap := make(map[int]bool, numOfKeys)
	points := make([]shamir.Point, 0, numOfKeys)

	for len(points) < numOfKeys {
		x := mrand.Intn(numOfKeys) + 1
		if ok := pointsMap[x]; ok {
			continue
		}
		pointsMap[x] = true
		points = append(points, shamir.Point{X: x, Fx: keys[x-1]})
	}

	return points
}
