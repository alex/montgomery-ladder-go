// DO NOT USE THIS CODE. SERIOUSLY.
package snakes

import (
	"math/big"
)

// Compute `a ** b % c`, just like `big.Int.Exp`, but do it in a way that is
// more constant time. This is useful for crypto stuff.
func MontgomeryLadderExp(a, b, c *big.Int) *big.Int {
	if a.Sign() == -1 || b.Sign() == -1 || c.Sign() == -1 {
		panic("MontgomeryLadderExp called with negative parameter.")
	}
	if b.Sign() == 0 {
		return big.NewInt(1)
	}

	a1 := new(big.Int).Set(a)
	a2 := new(big.Int).Mul(a, a)
	for pos := b.BitLen() - 2; pos >= 0; pos-- {
		if b.Bit(pos) == 0 {
			a2.Mul(a2, a1)
			a1.Mul(a1, a1)
		} else {
			a1.Mul(a1, a2)
			a2.Mul(a2, a2)
		}

		a1.Mod(a1, c)
		a2.Mod(a2, c)
	}
	return a1
}
