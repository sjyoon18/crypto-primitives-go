package modes

import "errors"

func PadPKCS7(data []byte, blockSize int) []byte {
	if blockSize <= 0 || blockSize > 255 {
		panic("invalid block size")
	}

	paddingSize := blockSize - (len(data) % blockSize)

	padded := make([]byte, len(data)+paddingSize)
	copy(padded, data)

	for i := len(data); i < len(padded); i++ {
		padded[i] = byte(paddingSize)
	}

	return padded
}

func UnpadPKCS7(padded []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 || blockSize > 255 {
		return nil, errors.New("invalid block size")
	}

	if len(padded) == 0 || len(padded)%blockSize != 0 {
		return nil, errors.New("invalid padded data length")
	}

	paddingSize := int(padded[len(padded)-1])

	if paddingSize == 0 || paddingSize > blockSize {
		return nil, errors.New("invalid padding")
	}

	if paddingSize > len(padded) {
		return nil, errors.New("invalid padding length")
	}

	for i := len(padded) - paddingSize; i < len(padded); i++ {
		if int(padded[i]) != paddingSize {
			return nil, errors.New("invalid padding bytes")
		}
	}

	data := make([]byte, len(padded)-paddingSize)
	copy(data, padded[:len(padded)-paddingSize])

	return data, nil
}
