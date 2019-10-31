package blockcipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"

	gos "github.com/golangaccount/go-libs/os"
)

var defaultNonce []byte

//DefaultNonce 设置GCM模式下nonce的默认值
func DefaultNonce(bt []byte) {
	if defaultNonce != nil {
		return
	}
	if bt == nil || len(bt) == 0 {
		return
	}
	defaultNonce := make([]byte, len(bt))
	copy(defaultNonce, bt)
}

var defaultIV []byte

//DefaultIV 设置默认的iv
func DefaultIV(bt []byte) error {
	if defaultIV != nil {
		return nil
	}
	if len(bt) != aes.BlockSize {
		return ErrAESIVLength
	}
	defaultIV = make([]byte, aes.BlockSize)
	copy(defaultIV, bt)
	return nil
}

//AESGCMEncryption AES加密的GCM模式
func AESGCMEncryption(password string, plaintext []byte, nonce []byte) ([]byte, []byte, error) {
	key := AESPassword(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	if nonce == nil {
		if defaultNonce != nil {
			nonce = make([]byte, len(defaultNonce))
			copy(nonce, defaultNonce)
		} else {
			nonce = make([]byte, 12)
			if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
				return nil, nil, err
			}
		}
	}

	gcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return nil, nil, err
	}
	return gcm.Seal(nil, nonce, plaintext, nil), nonce, nil
}

//AESGCMDecryption AES解密的GCM模式
func AESGCMDecryption(password string, ciphertext []byte, nonce []byte) ([]byte, error) {
	key := AESPassword(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if nonce == nil && defaultNonce != nil {
		nonce = make([]byte, len(defaultNonce))
		copy(nonce, defaultNonce)
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, nonce, ciphertext, nil)
}

//AESCBCEncryption AES CBC加密
func AESCBCEncryption(password string, plaintext []byte) ([]byte, error) {
	key := AESPassword(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = CBCPlaintext(plaintext)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

//AESCBCDecryption AES CBC解密
func AESCBCDecryption(password string, ciphertext []byte) ([]byte, error) {
	key := AESPassword(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return CBCCiphertext(ciphertext), nil
}

//AESStreamEncryption AES CFB/CTR/OFB
func AESStreamEncryption(password string, mode MODE, plaintext []byte) ([]byte, error) {
	key := AESPassword(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if defaultIV != nil {
		copy(iv, defaultIV)
	} else {
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return nil, err
		}
	}
	var stream cipher.Stream
	switch mode {
	case CFB:
		stream = cipher.NewCFBEncrypter(block, iv)
	case CTR:
		stream = cipher.NewCTR(block, iv)
	case OFB:
		stream = cipher.NewOFB(block, iv)
	default:
		return nil, ErrMODENotExists
	}
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

//AESStreamDecryption AES CFB/CTR/OFB
func AESStreamDecryption(password string, mode MODE, ciphertext []byte) ([]byte, error) {
	key := AESPassword(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	var stream cipher.Stream
	switch mode {
	case CFB:
		stream = cipher.NewCFBDecrypter(block, iv)
	case CTR:
		stream = cipher.NewCTR(block, iv)
	case OFB:
		stream = cipher.NewOFB(block, iv)
	default:
		return nil, ErrMODENotExists
	}
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

//AESStreamEncryptionFile AES CFB/CTR/OFB 文件
func AESStreamEncryptionFile(password string, mode MODE, inpath, outpath string) error {
	in, err := os.Open(inpath)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := gos.Create(outpath)
	if err != nil {
		return err
	}
	defer out.Close()
	return AESStreamEncryptionStream(password, mode, in, out)
}

//AESStreamDecryptionFile AES CFB/CTR/OFB 文件
func AESStreamDecryptionFile(password string, mode MODE, inpath, outpath string) error {
	in, err := os.Open(inpath)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := gos.Create(outpath)
	if err != nil {
		return err
	}
	defer out.Close()
	return AESStreamDecryptionStream(password, mode, in, out)
}

//AESStreamEncryptionStream AES CFB/CTR/OFB 文件
func AESStreamEncryptionStream(password string, mode MODE, in io.Reader, out io.Writer) error {
	key := AESPassword(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	iv := make([]byte, aes.BlockSize)
	if defaultIV != nil {
		copy(iv, defaultIV)
	} else {
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return err
		}
	}
	var stream cipher.Stream
	switch mode {
	case CFB:
		stream = cipher.NewCFBEncrypter(block, iv)
	case CTR:
		stream = cipher.NewCTR(block, iv)
	case OFB:
		stream = cipher.NewOFB(block, iv)
	default:
		return ErrMODENotExists
	}
	if defaultIV == nil {
		out.Write(iv)
	}
	writer := &cipher.StreamWriter{
		S: stream,
		W: out,
	}
	_, err = io.Copy(writer, in)
	return err
}

//AESStreamDecryptionStream AES CFB/CTR/OFB 文件
func AESStreamDecryptionStream(password string, mode MODE, in io.Reader, out io.Writer) error {
	key := AESPassword(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	iv := make([]byte, aes.BlockSize)
	if defaultIV != nil {
		copy(iv, defaultIV)
	} else {
		length, err := in.Read(iv)
		if err != nil || length != aes.BlockSize {
			return ErrAESStreamIv
		}
	}
	var stream cipher.Stream
	switch mode {
	case CFB:
		stream = cipher.NewCFBDecrypter(block, iv)
	case CTR:
		stream = cipher.NewCTR(block, iv)
	case OFB:
		stream = cipher.NewOFB(block, iv)
	default:
		return ErrMODENotExists
	}
	reader := &cipher.StreamReader{
		S: stream,
		R: in,
	}
	_, err = io.Copy(out, reader)
	return err
}
