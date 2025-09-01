package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

type Keys struct {
	N int `json:"n"`
	K int `json:"k"`
}

type Root struct {
	Base  string `json:"base"`
	Value string `json:"value"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <json_file>")
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	var data map[string]json.RawMessage
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	var keys Keys
	if err := json.Unmarshal(data["keys"], &keys); err != nil {
		log.Fatalf("Failed to parse keys: %v", err)
	}

	roots := []*big.Int{}

	// Collect first k-1 roots (for quadratic, k=3 → 2 roots)
	for i := 1; i <= keys.N && len(roots) < keys.K-1; i++ {
		key := strconv.Itoa(i)
		if raw, ok := data[key]; ok {
			var r Root
			if err := json.Unmarshal(raw, &r); err != nil {
				log.Fatalf("Failed to parse root %d: %v", i, err)
			}
			base, _ := strconv.Atoi(r.Base)

			// Convert string in given base to big.Int
			val, success := new(big.Int).SetString(r.Value, base)
			if !success {
				log.Fatalf("Invalid value %s for base %d", r.Value, base)
			}
			roots = append(roots, val)
		}
	}

	if len(roots) < 2 {
		log.Fatalf("Not enough roots to compute quadratic")
	}

	// c = r1 * r2
	c := new(big.Int).Mul(roots[0], roots[1])
	fmt.Println("c =", c)
}
