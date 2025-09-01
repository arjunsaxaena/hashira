package main

import (
	"fmt"
	"math/big"
)

func main() {
	roots := []string{
		"196563650089608567",
		"1016509518118225951",
		"3711974121218449851",
		"10788619898233492461",
		"995085094601491",
		"320923294898495900",
	}

	product := big.NewInt(1)

	for _, r := range roots {
		num := new(big.Int)
		num.SetString(r, 10)
		product.Mul(product, num)
	}

	fmt.Println("Product of roots =")
	fmt.Println(product.String())
}
