package practice

import "math/big"

func uint128FromBytes(bytes []byte) *big.Int {
	res := big.Int{}
	return res.SetBytes(bytes)
}
