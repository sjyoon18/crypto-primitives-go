# ECB Repeated-Block Leakage

## Overview

Electronic Codebook (ECB) mode encrypts identical plaintext blocks into identical ciphertext blocks.

This leaks structural information about the plaintext and makes repeated patterns visible in the ciphertext.

## Demonstration

This module:

- encrypts repeated plaintext blocks using AES-128 ECB,

- detects repeated ciphertext blocks,

- demonstrates structural leakage.

## Security Issue

ECB does not provide semantic security because plaintext patterns remain observable.

## Files

- `leakage.go`

- `demo.go`

- `leakage_test.go`