
# Simle Gnark Implementation
## Installation
Fallow https://go.dev/doc/install for installation golang.

## Installing dependencies
```bash
go mod tidy  
```

## Run Code
```bash
go run main.go 
```

```mermaid
sequenceDiagram
    participant M as main()
    participant C as Compile()
    participant S as Setup()
    participant NW as NewWitness()
    participant P as Prove()
    participant V as Verify()

    M->>C: Compile circuit into constraints
    C-->>M: Return compiled circuit (ccs)
    M->>S: Setup Proving and Verifying keys
    S-->>M: Return Proving Key (pk) and Verifying Key (vk)
    M->>NW: Create a witness from assignment
    NW-->>M: Return witness
    M->>P: Prove the witness
    P-->>M: Return proof
    M->>V: Verify the proof
    V-->>M: Verification result

```

### Circuit
 - A representation of a computation in zk-SNARKs, involving inputs, operations, and outputs.

 ### Compile
 - Transforming the circuit into a set of mathematical constraints for the zero-knowledge proof.

 ### Setup
 - Generating proving and verification keys based on the compiled circuit.

 ### Witness 
 - A solution to the circuit's constraints, consisting of both public and secret inputs.

 ### Prove
- Creating a proof that a valid witness is known, without revealing any information about the witness.

### Verify
- Checking if a proof is valid without learning anything about the witness.