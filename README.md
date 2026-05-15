# crypto-primitives-go

A cryptography-focused Go project implementing symmetric cryptographic primitives from scratch, with an emphasis on transparency, diffusion analysis, visualization, and educational tooling.

This repository currently focuses on AES-128 internals and includes:
- full AES encryption/decryption,
- round-by-round trace generation,
- avalanche-effect experiments,
- diffusion metrics,
- visualization utilities,
- and FIPS-197 verification tests.

---

# Features

## AES-128 Implementation
Implemented fully from scratch in Go:
- AES-128 encryption
- AES-128 decryption
- AES key expansion
- SubBytes / InvSubBytes
- ShiftRows / InvShiftRows
- MixColumns / InvMixColumns
- GF(2⁸) arithmetic

## AES Trace System
Round-by-round tracing of AES internals:
- plaintext state
- round keys
- intermediate round states
- ciphertext state

Useful for:
- learning AES internals,
- debugging,
- diffusion analysis,
- cryptographic visualization.

## Differential / Avalanche Analysis
Implements tooling for analyzing diffusion behavior under single-bit perturbations:
- single-bit flip experiments,
- state difference tracking,
- per-round diffusion growth,
- ciphertext bit difference statistics.

## Visualization Utilities
Utilities for visualizing AES behavior:
- AES state rendering,
- diffusion heatmaps,
- avalanche progression,
- S-box byte substitution,
- MixColumns transformations.

## Verification Tests
Includes correctness tests using official FIPS-197 vectors:
- AES encryption/decryption
- SubBytes
- ShiftRows
- MixColumns
- state conversion
- trace consistency

---

# Project Structure

```text
crypto-primitives-go/
│
├── symmetric/
│   ├── aes-128/
│   │   ├── encrypt.go
│   │   ├── decrypt.go
│   │   ├── subbytes.go
│   │   ├── shiftrows.go
│   │   ├── mixcolumns.go
│   │   ├── roundkeys.go
│   │   ├── trace.go
│   │   ├── experiments.go
│   │   ├── metrics.go
│   │   └── *_test.go
│   │
│   └── des/
│
├── visualization/
│   ├── aesviz/
│   └── desviz/
│
└── cmd/
    └── test/
```

---

# Example Usage

## AES Encryption

```go
plaintext := [16]byte{
    0x00, 0x11, 0x22, 0x33,
    0x44, 0x55, 0x66, 0x77,
    0x88, 0x99, 0xaa, 0xbb,
    0xcc, 0xdd, 0xee, 0xff,
}

key := [16]byte{
    0x00, 0x01, 0x02, 0x03,
    0x04, 0x05, 0x06, 0x07,
    0x08, 0x09, 0x0a, 0x0b,
    0x0c, 0x0d, 0x0e, 0x0f,
}

ciphertext := aes.EncryptBlock(plaintext, key)
```

---

## AES Trace Visualization

```go
trace := aes.EncryptWithTrace(plaintext, key)

aesviz.PrintTrace(trace)
```

---

## Avalanche Experiment

```go
experiment := aes.RunSingleBitFlipExperiment(
    plaintext,
    key,
    0,
)

aesviz.PrintSingleBitFlipAvalanche(experiment)
```

---

# Running Tests

Run all tests:

```bash
go test ./...
```

Run AES tests only:

```bash
go test -v -count=1 ./symmetric/aes-128
```

All AES correctness tests are verified against official FIPS-197 examples.

---

# Educational Focus

This repository is intentionally designed to expose AES internals clearly rather than prioritizing optimization or production deployment.

The goal is to build:
- cryptographic intuition,
- systems-level understanding,
- and practical experience implementing primitives from first principles.

---

# Future Work

Planned additions include:
- AES block cipher modes (CBC / CTR / GCM)
- SHA-256 implementation
- HMAC implementation
- DES analysis tooling
- differential cryptanalysis experiments
- serialization / JSON export
- benchmark tooling
- Rust/C implementations for comparison

---

# References

- FIPS-197: Advanced Encryption Standard (AES)
- NIST Cryptographic Standards
- The Design of Rijndael
- Understanding Cryptography (Paar & Pelzl)

---

# Disclaimer

This repository is intended for educational and research purposes.

Do not use these implementations directly in production systems.
Use well-audited cryptographic libraries for real-world applications.