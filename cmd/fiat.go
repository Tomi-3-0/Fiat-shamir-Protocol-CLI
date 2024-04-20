package cmd

// x = r² mod n, n = p × q, v = s mod n, challenge c, random number r,
// 2 large primes p and q. a secret number `s` between 1 and n-1 i.e., `(1 ≤ s <n)`.
// She calculates `v = s mod n` s= private key, v= public key
import (
	"fmt"
	"math/big"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcd/btcec/v2"
)

// S256 returns a Curve which implements secp256k1
var secp256k1Order = btcec.S256().Params().N

// returns the order of the elliptic curve secp256k1
func Secp256k1CurveOrder() (order *big.Int){
	return secp256k1Order
}

type FSProof struct {
	x, r, c, p, q, order *big.Int
}

// creates a new FSProof struct with some default values initialized and sets p
func Initialise(p *big.Int) *FSProof {
	return &FSProof{
		x: big.NewInt(1),
		r: big.NewInt(1),
		c: big.NewInt(1),
		p: p,
		q: big.NewInt(1),
	}
}

// generates a random integer within the range of [0,q), where q represents the upper bound of the range
func RandomIntegerGenerator(n *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, n)
}

// calculates and returns a non-interactive proof of knowledge using the Fiat-Shamir heuristic
func Prover(x, p *big.Int) (*FSProof, error){
	res := Initialise(p)
	res.order= (&big.Int{}).Sub(p, big.NewInt(1))

	//q = p^x(mod order) and assigns the result to res.q
	res.q = (&big.Int{}).Exp(p, x, res.order)

	//generates a random point v within the group of order p
	v, err := RandomIntegerGenerator(p); if err != nil{
		return res, err;
	}

	// calculates x = p^v (mod order)
	res.x.Exp(p, v, res.order)

	// concatenates the byte representations of res.p, res.q, and res.x, and hashes them using SHA256
	challenge := sha256.Sum256(append(append(res.p.Bytes(), res.q.Bytes()...), res.x.Bytes()...))
	// assigns the result to res.c
	res.c.SetBytes(challenge[:])

	// calculates response r = c × x − v (mod order) and assigns the result to res.r.
	res.r.Mul(res.c, x)
	res.r.Sub(v, res.r)
	res.r.Mod(res.r, res.order)

	return res, nil
}

func (proof *FSProof) Verifier() bool {
	// left-hand side (LHS) of the Fiat-Shamir equation, calculated as p^r (mod n)
	left := (&big.Int{}).Exp(proof.p, proof.r, proof.order)

	// right-hand side (RHS) of the Fiat-Shamir equation, calculated as q^s (mod n)
	right := (&big.Int{}).Exp(proof.q, proof.c, proof.order)

	x := (&big.Int{}).Mul(left, right)

	// checks if x is equal to proof.x
	return x.Cmp(proof.x) == 0

}

// Prints the fiat shamir proof in a provided base.
func (p *FSProof) Print(base int) {
	fields := []struct {
		name string
		val  *big.Int
	}{
		{"witness", p.x}, 		// witness
		{"response", p.r}, 		// response
		{"challenge", p.c},		// challenge
		{"p", p.p},				
		{"q", p.q},				
	}

	for _, f := range fields {
		fmt.Printf("%s: %s\n", f.name, f.val.Text(base))
	}
}
