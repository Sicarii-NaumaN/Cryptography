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
		pkBytes, err := hex.DecodeString("a11b0a4e1a132305652ee7a8eb7848f6ad5ea381e3ce20a2c086a2e388230811")
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		// The number of people who need to give out the sharing key
		keyShares := 13
		// The number of people needed to get the initial private key
		keyThreshold := 5

		fmt.Printf("Passed key:\n%x\n N = %d, T = %d\n", pkBytes, keyShares, keyThreshold)

		fmt.Println("-------------------------------------")

		// Spliting input key to keyShares sharing keys
		keys := Split(pkBytes, keyThreshold, keyShares)

		// Generating graph points of a function in a coordinate system
		// These points uses to restore polynomial
		points := PointsGeneration(len(keys), keys)

		stringKeys := make([]string, 0)
		fmt.Println("Points:")
		for i := range points {
			stringKeys = append(stringKeys, fmt.Sprintf("%d %x", points[i].X, points[i].Fx))
		}
		for i := range stringKeys {
			fmt.Println(stringKeys[i])
		}

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
		fmt.Println("Enter N and T...")
		fmt.Scan(&keyShares, &keyThreshold)
		hexSecret, _ := hex.DecodeString(secret)
		keys := Split(hexSecret, keyThreshold, keyShares)

		stringKeys := make([]string, 0)

		points := PointsGeneration(len(keys), keys)
		fmt.Println("keys:")
		for i := range points {
			stringKeys = append(stringKeys, fmt.Sprintf("%d %x", points[i].X, points[i].Fx))
		}
		for i := range stringKeys {
			fmt.Println(stringKeys[i])
		}
		return
	}

	// If flag -recover is passed
	if args.Recover {
		var X int
		var sharingKey string
		points := make([]shamir.Point, 0)

		fmt.Println("Enter sharing points...")
		for {
			_, err = fmt.Scan(&X, &sharingKey)
			if err != nil {
				break
			}
			pkBytes, _ := hex.DecodeString(sharingKey)
			point := shamir.Point{X: X, Fx: pkBytes}
			points = append(points, point)
		}

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
