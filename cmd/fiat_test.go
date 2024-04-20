package cmd

import (
	"math/big"
	"testing"
)

func TestFiat(t *testing.T) {
    // test parameters
    message := []byte("I know the password")
    x := big.NewInt(0).SetBytes(message)
    Order := Secp256k1CurveOrder()

    t.Run("secpk1", func(t *testing.T) {
        // Call Prover function
        result, err := Prover(x, Order)
        
        // Check for errors
        if err != nil {
            t.Errorf("Failed to Prove: %v", err)
        }

        // Verify proof
        if !result.Verifier() {
            t.Error("Failed to verify proof")
        }

    })
}
