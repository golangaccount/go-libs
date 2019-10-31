package blockcipher

import (
	"io/ioutil"
	"testing"
)

func init() {
	defaultNonce = make([]byte, 12)
	defaultNonce[0] = 12

	//DefaultIV([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
}

func TestAESGCMEncryption(t *testing.T) {
	password := "123"
	text := []byte("1234")
	ciphertext, nonce, err := AESGCMEncryption(password, text, nil)
	if err != nil {
		t.Log(err.Error())
	}
	text, err = AESGCMDecryption(password, ciphertext, nil)
	t.Log(nonce)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(text))
}

func TestAESCBCEncryption(t *testing.T) {
	password := "123"
	text := []byte("1234")
	ciphertext, err := AESCBCEncryption(password, text)
	if err != nil {
		t.Log(err.Error())
	}
	text, err = AESCBCDecryption(password, ciphertext)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(text))
}

func TestAESStreamEncryption(t *testing.T) {
	testAESStreamEncryption(t, CFB)
	testAESStreamEncryption(t, CTR)
	testAESStreamEncryption(t, OFB)
}

func testAESStreamEncryption(t *testing.T, mode MODE) {
	password := "123"
	text := []byte("1234werrtty中国")
	ciphertext, err := AESStreamEncryption(password, mode, text)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(ciphertext)
	text, err = AESStreamDecryption(password, mode, ciphertext)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(text))
}

func TestAESStreamEncryptionFile(t *testing.T) {
	testAESStreamEncryptionFile(t, CFB)
	testAESStreamEncryptionFile(t, CTR)
	testAESStreamEncryptionFile(t, OFB)
}

func testAESStreamEncryptionFile(t *testing.T, mode MODE) {
	t.Log(AESStreamEncryptionFile("123", mode, "./in", "./out"))
	t.Log(AESStreamDecryptionFile("123", mode, "./out", "./in"))
	bts, err := ioutil.ReadFile("./in")
	t.Log(string(bts), err)
}

func TestAESStream(t *testing.T) {
	testAESStream(t, CFB)
	testAESStream(t, CTR)
	testAESStream(t, OFB)
}

func testAESStream(t *testing.T, mode MODE) {
	password := "123"
	text := []byte("1234werrtty中国")
	ciphertext, err := AESStreamEncryption(password, mode, text)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(ioutil.WriteFile("./out", ciphertext, 0666))

	t.Log(AESStreamDecryptionFile("123", mode, "./out", "./in"))
	bts, err := ioutil.ReadFile("./in")
	t.Log(string(bts), err)

}

func TestAESStreamR(t *testing.T) {
	testAESStreamR(t, CFB)
	testAESStreamR(t, CTR)
	testAESStreamR(t, OFB)
}

func testAESStreamR(t *testing.T, mode MODE) {
	t.Log(AESStreamEncryptionFile("123", mode, "./in", "./out"))
	bts, err := ioutil.ReadFile("./out")
	if err != nil {
		t.Log(err.Error())
	}
	bts, err = AESStreamDecryption("123", mode, bts)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(bts))
}
