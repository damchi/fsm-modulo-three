#  FSM Modulo

##  Overview

This project implements a **Finite State Machine (FSM)** generator in Go.  
It demonstrates how to define, initialize, and execute state machines using **generic programming**.

As an example, the repository includes the implementation of a **Modulo Three** FSM — a simple automaton that tracks the remainder when dividing a binary number by a chosen number.

---

##  Finite Automaton Definition

A **Finite Automaton (FA)** is formally defined as a 5-tuple:

> **FA = (Q, Σ, q₀, F, δ)**

Where:

- **Q** → a finite set of states
- **Σ** → a finite input alphabet
- **q₀ ∈ Q** → the initial state
- **F ⊆ Q** → the set of accepting (final) states
- **δ : Q × Σ → Q** → the transition function

For any element **q ∈ Q** and symbol **σ ∈ Σ**,  
`δ(q, σ)` represents the next state when the automaton is in **q** and receives **σ**.

---


## Project Structure 
```
fsm-module-three/
├── cmd/
│   └── app/
│       └── main.go               # Starts the FSM example
├── internal/
│   ├── fsm/
│   │   ├── fsm.go                # FSM structure and logic
│   │   ├── modulo.go             # Modulo service using FSM
│   ├── adapters/
│   │   ├── routes/
│   │   │   └── routes.go         # Gin routes setup
│   │   └── handler/
│   │       ├── modulo_handler.go # Gin handler for /check endpoint
│   │       └── modulo_handler_integration_test.go # Integration tests
├── doc/                          # Generated API documentation
├── go.mod
└── README.md
```

### Make File commands

```makefile
make run #start the project
make test #run the tests
make doc  #generate the doc 
make clean  #clean the generated doc 
```