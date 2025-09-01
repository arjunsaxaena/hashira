package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
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

func mustAtoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("invalid base %q: %v", s, err)
	}
	return v
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go <json_file>")
	}
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open %s: %v", filename, err)
	}
	defer f.Close()

	var raw map[string]json.RawMessage
	if err := json.NewDecoder(f).Decode(&raw); err != nil {
		log.Fatalf("failed to parse JSON: %v", err)
	}

	var keys Keys
	if err := json.Unmarshal(raw["keys"], &keys); err != nil {
		log.Fatalf("failed to parse keys: %v", err)
	}
	if keys.K < 2 {
		log.Fatalf("k must be >= 2")
	}

	m := keys.K - 1
	roots := []*big.Int{}

	idxs := make([]int, 0, keys.N)
	for i := 1; i <= keys.N; i++ {
		if _, ok := raw[strconv.Itoa(i)]; ok {
			idxs = append(idxs, i)
		}
	}
	sort.Ints(idxs)

	for _, i := range idxs {
		if len(roots) >= m {
			break
		}
		var r Root
		if err := json.Unmarshal(raw[strconv.Itoa(i)], &r); err != nil {
			log.Fatalf("failed to parse root %d: %v", i, err)
		}
		base := mustAtoi(r.Base)
		if base < 2 || base > 62 {
			log.Fatalf("base %d out of range [2,62]", base)
		}
		val, ok := new(big.Int).SetString(r.Value, base)
		if !ok {
			log.Fatalf("invalid value %q for base %d at index %d", r.Value, base, i)
		}
		roots = append(roots, val)
	}

	if len(roots) < m {
		log.Fatalf("not enough roots: need %d, got %d", m, len(roots))
	}

	prod := big.NewInt(1)
	for _, r := range roots {
		prod.Mul(prod, r)
	}
	if m%2 == 1 {
		prod.Neg(prod)
	}

	fmt.Println("c =", prod.String())
}
