package main

import (
	"encoding/base64"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"log"
	mrand "math/rand"
	"strings"

	shamir "shamir/shamir"
)

func main() {
	keyTest := "Qlco35ne5pbmhu5eBbFOp2yCLMdNFD/IJ9sq+ZcdoRY="
	fmt.Println("keyTest: ", keyTest)
	keyShares := 5
	keyThreshold := 5
	useKeys := 5

	keyBytes, err := base64.StdEncoding.DecodeString(keyTest)

	if err != nil {
		log.Fatalf("Unexpected error occurred while decoding key; %v", err)
	}
	var key [32]byte
	copy(key[:], keyBytes)

	stringKeys := make([]string, 0)
	keys, err := shamir.GenKeyShares(key, keyThreshold, keyShares)
	for i := range keys {
		stringKeys = append(stringKeys, string(keys[i]))
	}
	fmt.Println("KEYS: ", stringKeys)
	pointsMap := make(map[int]bool, useKeys)
	points := make([]shamir.Point, 0, useKeys)
	fmt.Println("KEYBYTES: ", keyBytes)

	// Computed just for readable test output
	var encodedPoints strings.Builder
	encodedPoints.WriteString("\n")

	for len(points) < useKeys {
		x := mrand.Intn(keyShares) + 1
		if ok := pointsMap[x]; ok {
			continue
		}
		pointsMap[x] = true
		points = append(points, shamir.Point{X: x, Fx: keys[x-1]})
		encodedPoints.WriteString(fmt.Sprintf("\t'%d-%s'\n", x, base64.StdEncoding.EncodeToString(keys[x-1][:])))
	}

	derivedKey, err := shamir.GetKeyFromKeyShares(points)
	fmt.Println("DERIVEDKEY" , derivedKey)
	decodedKey := base64.StdEncoding.EncodeToString(derivedKey[:])
	fmt.Println("DECODEDKEY " , decodedKey)
	diff := cmp.Diff(keyTest, decodedKey)
	fmt.Println(diff)
}
