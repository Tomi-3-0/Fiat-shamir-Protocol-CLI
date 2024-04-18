package cmd

// x = r² mod n, n = p × q, v = s mod n, challenge c, random number r,
// 2 large primes p and q. a secret number `s` between 1 and n-1 i.e., `(1 ≤ s <n)`.
// She calculates `v = s mod n` s= private key, v= public key
import (
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

// S256 returns a Curve which implements secp256k1
var secp256k1Order = btcec.S256().Params().N

func Main() {
	fmt.Println("Order of secp256k1 curve:", secp256k1Order)
}

type FSProof struct {
	x, r, s, p, q, n *big.Int
}

// creates a new FSProof struct with some default values initialized and sets p
func initialise(p *big.Int) *FSProof {
	return &FSProof{
		x: big.NewInt(1),
		r: big.NewInt(1),
		s: big.NewInt(1),
		p: p,
		q: big.NewInt(1),
	}
}

func (proof *FSProof) Verifier() bool {
	// left-hand side (LHS) of the Fiat-Shamir equation, calculated as p^r (mod n)
	left := (&big.Int{}).Exp(proof.p, proof.r, proof.n)

	// right-hand side (RHS) of the Fiat-Shamir equation, calculated as q^s (mod n)
	right := (&big.Int{}).Exp(proof.q, proof.s, proof.n)

	x := (&big.Int{}).Mul(left, right)

	// checks if x is equal to proof.x
	return x.Cmp(proof.x) == 0

}
