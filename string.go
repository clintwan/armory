package armory

import (
	"bytes"
	"crypto/md5"
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

// String string
var String *str

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

// MD5Encode MD5 Encode
func (s *str) MD5Encode(source []byte) string {
	r := fmt.Sprintf("%x", md5.Sum(source))
	return r
}

// RandomString RandomString
func (s *str) RandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := []byte(strconv.Itoa(rand.Int()))
	return s.MD5Encode(b)
}

// GetIntArray getIntArray
func (s *str) GetIntArray(str string) []int {
	arr := []int{}
	if len(str) > 0 {
		for _, s := range strings.Split(str, ",") {
			i, _ := strconv.Atoi(s)
			if i > 0 {
				arr = append(arr, i)
			}
		}
	}
	return arr
}
