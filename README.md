# crypto-primitives-go

An educational Go implementation of classical cryptographic primitives with
a strong emphasis on **bit-level transparency**, **diffusion/confusion analysis**,
and **visualization-driven understanding**.

This project is designed for learning, inspection, and experimentation —
not for production use.

---

## Motivation

Modern cryptographic libraries prioritize performance and security guarantees,
but often obscure the internal transformations that give block ciphers their
strength.

This project exists to make those transformations **visible**.

By explicitly implementing primitives such as DES and AES-128 from first
principles and pairing them with visualization tools, the repository allows
developers and students to *see* how properties like diffusion, avalanche,
and non-linearity emerge across rounds.

---

## Design Philosophy

- **Correctness over performance**
- **Explicit bit-level operations**
- **No reliance on standard crypto libraries for core algorithms**
- **Visualization as a first-class feature**
- **Clear separation between cryptographic logic and analysis tools**

---

## Implemented Primitives

### Symmetric Ciphers

#### DES
- Full encryption and decryption
- Explicit Feistel structure
- Round-by-round tracing
- Avalanche effect measurement
- Bit-level permutation inspection

#### AES-128
- Full AES-128 block cipher
- SubBytes, ShiftRows, MixColumns, AddRoundKey
- Galois Field arithmetic implemented explicitly
- Round-by-round state inspection

---

### AES Visualizations

#### MixColumns Column Transformation
```go
aesviz.VisualizeMixColumn([4]byte{0xdb, 0x13, 0x53, 0x45})
```
Displays how a single AES column transforms under MixColumns.

#### MixColumns Influence Analysis
```go
aesviz.VisualizeMixColumnInfluence(0, 0x01)
```
Demonstrates how changing one byte affects the entire column,
illustrating linear diffusion in GF(2⁸).

#### S-box Byte Substitution
```go
aesviz.VisualizeSBoxByte(0x53, sbox)
```
Visualizes byte substitution and bit changes induced by the AES S-box.

---

### Example output
```code
Input column:
  [0] 01
  [1] 00
  [2] 00
  [3] 00

Output column:
  [0] 02
  [1] 01
  [2] 01
  [3] 03
```
This demonstrates how MixColumns spreads a single-byte difference across multiple output bytes.

---

### Project Structure
```code
symmetric/
├── des/
└── aes-128/

visualization/
├── desviz/
└── aesviz/
```
- **symmetric/ contains core cipher implementations**
- **visualization/ contains analysis and educational tooling**
-  **The visualization code lives separately from core implementations and never modifies core cipher logic**

---

### Intended Use

This repository is intended for:
- **Cryptography/Security engineers seeking deeper intuition**
- **Researchers analyzing diffusion and avalanche behavior**
- **Developers learning block cipher internals**

It is not intended for:
- **Production cryptography**
- **Performance benchmarking**
- **Security-critical applications**

---

### References
	
- **FIPS PUB 46-3 — Data Encryption Standard (DES)**
- **FIPS PUB 197 — Advanced Encryption Standard (AES)**

---

### Future work

- **Round-level diffusion visualization for AES**
- **Comparative avalanche analysis (DES vs AES)**
- **NIST test vector verification**
- **Additional primitives (e.g. SHA-family internals)**
