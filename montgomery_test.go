package snakes

import (
	"math/big"
	"testing"
)

func check(t *testing.T, a, b, c int) {
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	bigC := big.NewInt(int64(c))
	z := new(big.Int).Exp(bigA, bigB, bigC)
	got := MontgomeryLadderExp(bigA, bigB, bigC)
	if got.Cmp(z) != 0 {
		t.Errorf("pow(%d, %d, %d) Expected: %s, got %s", a, b, c, z.String(), got.String())
	}
}

func TestMontgomeryLadderModExp(t *testing.T) {
	check(t, 0, 0, 1)
	check(t, 0, 1, 1)
	check(t, 0, 2, 1)
	check(t, 1, 0, 1)
	check(t, 2, 0, 1)
	check(t, 2, 10, 1)
	check(t, 4, 13, 1)
	check(t, 13, 4, 1)
	check(t, 2, 10, 15)
	check(t, 3, 10, 15)
	check(t, 13, 4, 7)
	check(t, 19, 65, 3)
	check(t, 1432, 432, 123)
}
