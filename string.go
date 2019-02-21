package armory

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type str struct{}

var String *str

// ParseJSON 解析json
func (s *str) ParseJSON(v interface{}, pretty bool) (*string, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if pretty {
		encoder.SetIndent("", "    ")
	}
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	c := string(buffer.Bytes())
	return &c, err
}

// PrettyJSON 美化json
func (s *str) PrettyJSON(bts []byte) (*string, error) {
	v := interface{}(nil)
	json.Unmarshal(bts, &v)
	return s.ParseJSON(v, true)
}

// MD5Encode MD5 Encode
func (s *str) MD5Encode(source *[]byte) *string {
	r := fmt.Sprintf("%x", md5.Sum(*source))
	return &r
}

// RandomString RandomString
func (s *str) RandomString() *string {
	rand.Seed(time.Now().UnixNano())
	b := []byte(strconv.Itoa(rand.Int()))
	return s.MD5Encode(&b)
}

// DesEncrypt DesEncrypt
func (s *str) DesEncrypt(origData, key, offset []byte, mode, padding string) []byte {
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
func (s *str) pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	//计算出需要补多少位
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
	// 需要补padding位的padding值
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	//把补充的内容拼接到明文后面
	return append(ciphertext, padtext...)
}

// DesDecrypt DesDecrypt
func (s *str) DesDecrypt(crypted, key, offset []byte, mode, padding string) []byte {
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
func (s *str) pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	//解密去补码时需取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文
	return origData[:(length - unpadding)]
}

// FindStringInSlice FindStringInSlice
func (s *str) FindStringInSlice(sa *[]string, str string) int {
	idx := -1
	for i, s := range *sa {
		if s == str {
			idx = i
		}
	}
	return idx
}

// GetIntArray getIntArray
func (s *str) GetIntArray(str *string) *[]int {
	arr := []int{}
	if len(*str) > 0 {
		for _, s := range strings.Split(*str, ",") {
			i, _ := strconv.Atoi(s)
			if i > 0 {
				arr = append(arr, i)
			}
		}
	}
	return &arr
}

// Utf8ToGbk Utf8ToGbk
func (s *str) Utf8ToGbk(bts []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(bts), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// GbkToUtf8 GbkToUtf8
func (s *str) GbkToUtf8(bts []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(bts), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
