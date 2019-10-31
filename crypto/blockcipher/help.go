package blockcipher

import "crypto/aes"

//AESPassword aes加密密钥处理，使用0进行填充
func AESPassword(password string) []byte {
	bt := []byte(password)
	length := len(bt)
	var result []byte
	if length == 16 || length == 24 || length == 32 {
		return bt
	}

	if length < 16 {
		result = make([]byte, 16)
	} else if length < 24 {
		result = make([]byte, 24)
	} else if length < 32 {
		result = make([]byte, 32)
	} else {
		panic(ErrAESPasswordTolang)
	}
	copy(result, bt)
	return result
}

//CBCPlaintext cbc模式下使用0进行填充
func CBCPlaintext(plaintext []byte) []byte {
	length := len(plaintext)
	size := length % aes.BlockSize
	result := make([]byte, aes.BlockSize*(length/aes.BlockSize+1))
	result[0] = byte(size)
	copy(result[1:], plaintext)
	return result
}

//CBCCiphertext cbc模式下对0进行填充的数据进行原始数据提取
func CBCCiphertext(ciphertext []byte) []byte {
	size := int(ciphertext[0])
	length := len(ciphertext)
	result := make([]byte, length-16+size)
	copy(result, ciphertext[1:])
	return result
}
