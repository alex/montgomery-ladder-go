package snakes

import (
	"math/big"
	"testing"
)

func checkExpInt(t *testing.T, a, b, c int) {
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	bigC := big.NewInt(int64(c))
	checkExp(t, bigA, bigB, bigC)
}

func checkExp(t *testing.T, a, b, c *big.Int) {
	z := new(big.Int).Exp(a, b, c)
	got := MontgomeryLadderExp(a, b, c)
	if got.Cmp(z) != 0 {
		t.Errorf("pow(%s, %s, %s) Expected: %s, got %s", a, b, c, z, got)
	}
}

func TestMontgomeryLadderExp(t *testing.T) {
	checkExpInt(t, 0, 0, 1)
	checkExpInt(t, 0, 1, 1)
	checkExpInt(t, 0, 2, 1)
	checkExpInt(t, 1, 0, 1)
	checkExpInt(t, 2, 0, 1)
	checkExpInt(t, 2, 10, 1)
	checkExpInt(t, 4, 13, 1)
	checkExpInt(t, 13, 4, 1)
	checkExpInt(t, 2, 10, 15)
	checkExpInt(t, 3, 10, 15)
	checkExpInt(t, 13, 4, 7)
	checkExpInt(t, 19, 65, 3)
	checkExpInt(t, 1432, 432, 123)
	checkExpInt(t, 1, 65537, 2)
	a, _ := new(big.Int).SetString("2938462938472983472983659726349017249287491026512746239764525612965293865296239471239874193284792387498274256129746192347", 0)
	b, _ := new(big.Int).SetString("298472983472983471903246121093472394872319615612417471234712061", 0)
	c, _ := new(big.Int).SetString("29834729834729834729347290846729561262544958723956495615629569234729836259263598127342374289365912465901365498236492183464", 0)
	checkExp(t, a, b, c)
}

func BenchmarkBigIntExp(b *testing.B) {
	x := new(big.Int).Lsh(big.NewInt(1), 1374)
	y := new(big.Int).Lsh(big.NewInt(1), 4096)
	z := new(big.Int).Lsh(big.NewInt(1), 234)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		new(big.Int).Exp(x, y, z)
	}
}

func BenchmarkMontgomeryLadderExp(b *testing.B) {
	x := new(big.Int).Lsh(big.NewInt(1), 1374)
	y := new(big.Int).Lsh(big.NewInt(1), 4096)
	z := new(big.Int).Lsh(big.NewInt(1), 234)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MontgomeryLadderExp(x, y, z)
	}
}
