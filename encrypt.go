package armory

import (
	"bytes"
	"crypto"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

type encrypt struct{}

// Encrypt encrypt
var Encrypt *encrypt

// GenerateRsaKey GenerateRsaKey
func (s *encrypt) GenerateRsaKey(bit int, privateKeyPath, publicKeyPath string) error {
	private, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		return err
	}
	//x509私钥序列化
	privateStream := x509.MarshalPKCS1PrivateKey(private)
	//将私钥设置到pem结构中
	block := pem.Block{
		Type:  "Rsa Private Key",
		Bytes: privateStream,
	}
	//保存磁盘
	file, err := os.Create(privateKeyPath)
	if err != nil {
		return err
	}
	//pem编码
	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}
	//=========public=========
	public := private.PublicKey
	//509序列化
	publicStream, err := x509.MarshalPKIXPublicKey(&public)
	if err != nil {
		return err
	}
	//公钥赋值pem结构体
	pubblock := pem.Block{Type: "Rsa Public Key", Bytes: publicStream}
	//保存磁盘
	pubfile, err := os.Create(publicKeyPath)
	if err != nil {
		return err
	}
	//pem编码
	err = pem.Encode(pubfile, &pubblock)
	if err != nil {
		return err
	}
	return nil
}

// SignatureRSA SignatureRSA
func (s *encrypt) RsaSignature(sourceData, privateKey []byte) ([]byte, error) {
	msg := []byte("")
	//解析
	block, _ := pem.Decode(privateKey)
	pKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	//哈希加密
	myHash := sha256.New()
	myHash.Write(sourceData)
	hashRes := myHash.Sum(nil)
	//对哈希结果进行签名
	res, err := rsa.SignPKCS1v15(rand.Reader, pKey, crypto.SHA256, hashRes)
	if err != nil {
		return msg, err
	}
	return res, nil
}

// VerifyRSA VerifyRSA
func (s *encrypt) VerifyRsaSignature(sourceData, signedData, publicKey []byte) error {
	//pem解密
	block, _ := pem.Decode(publicKey)
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pKey := publicInterface.(*rsa.PublicKey)
	//元数据哈希加密
	mySha := sha256.New()
	mySha.Write(sourceData)
	res := mySha.Sum(nil)

	//校验签名
	err = rsa.VerifyPKCS1v15(pKey, crypto.SHA256, res, signedData)
	if err != nil {
		return err
	}
	return nil
}

// RsaEncrypt RsaEncrypt
func (s *encrypt) RsaEncryptWithPublicKey(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt RsaDecrypt
func (s *encrypt) RsaDecryptWithPrivateKey(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// DesEncrypt DesEncrypt
func (s *encrypt) DesEncrypt(origData, key, offset []byte, mode, padding string) []byte {
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(key)
	//对明文先进行补码操作
	origData = s.pkcs5Padding(origData, block.BlockSize())
	//设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, offset)
	//创建明文长度的字节数组
	crypted := make([]byte, len(origData))
	//加密明文,加密后的数据放到数组中
	blockMode.CryptBlocks(crypted, origData)
	//将字节数组转换成字符串
	return crypted

}

// 实现明文的补码
func (s *encrypt) pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	//计算出需要补多少位
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
	// 需要补padding位的padding值
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//把补充的内容拼接到明文后面
	return append(ciphertext, padtext...)
}

// DesDecrypt DesDecrypt
func (s *encrypt) DesDecrypt(crypted, key, offset []byte, mode, padding string) []byte {
	if len(crypted) == 0 {
		return nil
	}
	//倒叙执行一遍加密方法
	//将字符串转换成字节数组
	// crypted, _ := base64.StdEncoding.DecodeString(data)
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(key)
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block, offset)
	//创建密文大小的数组变量
	origData := make([]byte, len(crypted))
	//解密密文到数组origData中
	blockMode.CryptBlocks(origData, crypted)
	//去补码
	origData = s.pkcs5UnPadding(origData)
	//打印明文
	return origData
	// fmt.Println(string(origData))
}

// 去除补码
func (s *encrypt) pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	//解密去补码时需取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文
	return origData[:(length - unpadding)]
}
