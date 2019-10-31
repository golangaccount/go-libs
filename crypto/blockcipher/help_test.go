package blockcipher

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"io"
	"testing"
)

func TestFunc(t *testing.T) {
	cip, err := aes.NewCipher([]byte("123344123456123412345678"))
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(cip.BlockSize())

}

func TestAESPassword(t *testing.T) {
	t.Log(AESPassword("123"))
	t.Log(AESPassword("12345678901234567890"))
	t.Log(AESPassword("123456789012345678901234567890"))
	t.Log(AESPassword("1234567890123456789012345678908888888888888888"))
}

func TestCBCPlaintext(t *testing.T) {
	var bts []byte
	for i := 1; i < 20; i++ {
		bts = make([]byte, i)
		io.ReadFull(rand.Reader, bts)
		t.Log(CBCPlaintext(bts))
		t.Log(bytes.Equal(bts, CBCCiphertext(CBCPlaintext(bts))))
	}
}
