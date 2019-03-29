package armory

import (
	"bytes"
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

// ParseJSON 解析json
func (s *str) ParseJSON(v interface{}, pretty bool) (string, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	if pretty {
		encoder.SetIndent("", "    ")
	}
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	c := string(bytes.TrimSpace(buffer.Bytes()))
	return c, err
}

// PrettyJSON 美化json
func (s *str) PrettyJSON(bts []byte) (string, error) {
	v := interface{}(nil)
	json.Unmarshal(bts, &v)
	return s.ParseJSON(v, true)
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
