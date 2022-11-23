package main

import (
	"fmt"
	"math/big"
)

func main() {
	//console.Menu()

	var res []*big.Int

	res = append(res, big.NewInt(10))
	res = append(res, big.NewInt(10))
	res = append(res, big.NewInt(10))
	res = append(res, big.NewInt(10))

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

}
