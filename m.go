package snakes

import (
    "math/big"
)


func MontgomeryLadderExp(a, b, c *big.Int) *big.Int {
    zero := big.NewInt(0)
    if a.Cmp(zero) == 0 && b.Cmp(zero) == 0 {
        return big.NewInt(1)
    }
    a1 := a
    a2 := new(big.Int).Mul(a, a)
    for pos := b.BitLen() - 1; pos >= 0; pos -- {
        if b.Bit(pos) == 0 {
            a2 = new(big.Int).Mul(a2, a1)
            a1 = new(big.Int).Mul(a1, a1)
        } else {
            a1 = new(big.Int).Mul(a1, a2)
            a2 = new(big.Int).Mul(a2, a2)
        }

        a1 = new(big.Int).Mod(a1, c)
        a2 = new(big.Int).Mod(a2, c)
    }
    return a1
}
