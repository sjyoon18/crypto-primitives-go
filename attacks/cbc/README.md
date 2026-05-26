# CBC Bit-Flipping Attack

## Overview

CBC encryption provides confidentiality but does not provide integrity.

An attacker can modify ciphertext block `C(i-1)` to cause controlled changes in decrypted plaintext block `P(i)`.

## Demonstration

This module:

- modifies ciphertext bytes,

- flips selected plaintext values after decryption,

- demonstrates controlled plaintext manipulation without knowing the AES key.

## Security Issue

CBC mode alone is malleable and vulnerable to active ciphertext manipulation.

Authenticated encryption modes such as GCM prevent this issue.

## Files

- `bitflip.go`

- `demo.go`

- `bitflip_test.go`