package main

import (
	"fmt"
	"math/big"
)

func main() {
	roots := []struct {
		value string
		base  int
	}{
		{"13444211440455345511", 6},
		{"aed7015a346d635", 15},
		{"6aeeb69631c227c", 15},
		{"e1b5e05623d881f", 16},
		{"316034514573652620673", 8},
		{"2122212201122002221120200210011020220200", 3},
	}

	product := big.NewInt(1)

	for _, r := range roots {
		num := new(big.Int)
		_, ok := num.SetString(r.value, r.base)
		if !ok {
			fmt.Printf("Error converting %s from base %d\n", r.value, r.base)
			return
		}
		product.Mul(product, num)
	}

	fmt.Println("Product of roots =")
	fmt.Println(product.String())
}
