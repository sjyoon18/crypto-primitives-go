# CTR Nonce Reuse Attack

## Overview

CTR mode transforms a block cipher into a stream cipher.

If the same nonce and key are reused, the same keystream is reused.

This causes:

```text

C1 XOR C2 = P1 XOR P2

```

which leaks information about both plaintexts.

## Demonstration

This module:

- encrypts two plaintexts with the same nonce,

- XORs the ciphertexts,

- recovers the XOR of the original plaintexts.

## Security Issue

Nonce reuse in CTR mode completely breaks confidentiality.

## Files

- `nonce_reuse.go`

- `demo.go`

- `nonce_reuse_test.go`