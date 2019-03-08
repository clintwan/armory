package armory

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var publicKey, privateKey []byte

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	// teardown()
	os.Exit(retCode)
}

func setup() {
	privateKeyPath := "./private.pem"
	publicKeyPath := "./public.pem"
	Encrypt.GenerateRsaKey(1024, privateKeyPath, publicKeyPath)
	var err error
	publicKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println(err)
	}
	privateKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println(err)
	}
}
func TestRsaEncrypt(t *testing.T) {
	sourceData := []byte("asdf123456789")

	bts, err := Encrypt.RsaEncryptWithPublicKey(sourceData, publicKey)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(bts))

	r, err := Encrypt.RsaDecryptWithPrivateKey(bts, privateKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r))
}

func TestRsaSignature(t *testing.T) {
	sourceData := []byte("asdf123456789")
	signData, err := Encrypt.RsaSignature(sourceData, privateKey)
	if err != nil {
		fmt.Println(err)
	}
	err = Encrypt.VerifyRsaSignature(sourceData, signData, publicKey)
	if err != nil {
		fmt.Println("校验出错:", err)
	} else {
		fmt.Println("校验正确:")
	}
}
