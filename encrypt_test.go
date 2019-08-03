package armory

import (
	"encoding/base64"
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
	Encrypt.RsaGenerateKey(2048, privateKeyPath, publicKeyPath)
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

	bts, err := Encrypt.RsaEncrypt(sourceData, publicKey)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(bts))

	r, err := Encrypt.RsaDecrypt(bts, privateKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r) == string(sourceData))
}

func TestRsaSignature(t *testing.T) {
	sourceData := []byte("asdf123456789")
	fmt.Println(string(sourceData))

	signDataMD5, errMD5 := Encrypt.RsaSignatureWithMD5(sourceData, privateKey)
	if errMD5 != nil {
		fmt.Println(errMD5)
	}
	errMD5 = Encrypt.RsaSignatureVerifyWithMD5(sourceData, signDataMD5, publicKey)
	signDataMD5String := base64.StdEncoding.EncodeToString(signDataMD5)
	fmt.Println(signDataMD5String)
	if errMD5 != nil {
		fmt.Println("Verify error:", errMD5)
	} else {
		fmt.Println("Verify success:")
	}

	signDataSha1, errSha1 := Encrypt.RsaSignatureWithSha1(sourceData, privateKey)
	if errSha1 != nil {
		fmt.Println(errSha1)
	}
	errSha1 = Encrypt.RsaSignatureVerifyWithSha1(sourceData, signDataSha1, publicKey)
	signDataSha1String := base64.StdEncoding.EncodeToString(signDataSha1)
	fmt.Println(signDataSha1String)
	if errSha1 != nil {
		fmt.Println("Verify error:", errSha1)
	} else {
		fmt.Println("Verify success:")
	}

	signDataSha256, errSha256 := Encrypt.RsaSignatureWithSha256(sourceData, privateKey)
	if errSha256 != nil {
		fmt.Println(errSha256)
	}
	errSha256 = Encrypt.RsaSignatureVerifyWithSha256(sourceData, signDataSha256, publicKey)
	signDataSha256String := base64.StdEncoding.EncodeToString(signDataSha256)
	fmt.Println(signDataSha256String)
	if errSha256 != nil {
		fmt.Println("Verify error:", errSha256)
	} else {
		fmt.Println("Verify success:")
	}
}
