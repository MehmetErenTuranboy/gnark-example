package main

import (
	"fmt"
	"log"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// MyCircuit is a struct representing the zk-SNARK circuit.
type MyCircuit struct {
	X frontend.Variable `gnark:",secret"` // X is a secret variable
	Y frontend.Variable `gnark:",public"`  // Y is a public variable
}

// Define defines the logic of the circuit. It declares the set of constraints that
// the witness must satisfy to create a valid zk-SNARK.
func (circuit *MyCircuit) Define(api frontend.API) error {
	cubedX := api.Mul(circuit.X, circuit.X, circuit.X) // Compute X**3
	api.AssertIsEqual(circuit.Y, api.Add(cubedX, circuit.X, 7)) // Assert that X**3 + X + 7 == Y
	return nil
}

func main() {
	var circuit MyCircuit

	// Compile the circuit into a set of constraints
	ccs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuit)
	if err != nil {
		log.Fatalf("Failed to compile the circuit: %v", err)
	}

	// Setup the Proving and Verifying keys
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		log.Fatalf("Failed to setup the proving and verifying keys: %v", err)
	}

	// Define the witness
	assignment := MyCircuit{X: 4, Y: 75}

	// Create a witness from the assignment
	witness, err := frontend.NewWitness(&assignment, ecc.BN254)
	if err != nil {
		log.Fatalf("Failed to create a witness: %v", err)
	}

	// Extract the public part of the witness
	publicWitness, err := witness.Public()
	if err != nil {
		log.Fatalf("Failed to extract the public witness: %v", err)
	}

	// Prove the witness
	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		log.Fatalf("Failed to prove the witness: %v", err)
	}

	fmt.Println("Generated Proof:", proof)

	// Verify the proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		fmt.Println("Verification Result: Failed")
		log.Fatalf("Failed to verify the proof: %v", err)
	} else {
		fmt.Println("Verification Result: Success")
	}
}
