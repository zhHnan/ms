package encrypts

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Md5(pwd string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, pwd)
	return hex.EncodeToString(hash.Sum(nil))
}

func Encrypt(plainText string, keyText string) (cipherStr string, err error) {
	// 转换成字节数据, 方便加密
	plainByte := []byte(plainText)
	keyByte := []byte(keyText)
	// 创建加密算法aes
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	//加密字符串
	cfb := cipher.NewCTR(c, commonIV)
	cipherByte := make([]byte, len(plainByte))
	cfb.XORKeyStream(cipherByte, plainByte)
	cipherStr = hex.EncodeToString(cipherByte)
	return
}
func Decrypt(cipherStr string, keyText string) (plainText string, err error) {
	// 转换成字节数据, 方便加密
	keyByte := []byte(keyText)
	// 创建加密算法aes
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	// 解密字符串
	cfbdec := cipher.NewCTR(c, commonIV)
	cipherByte, _ := hex.DecodeString(cipherStr)
	plainByte := make([]byte, len(cipherByte))
	cfbdec.XORKeyStream(plainByte, cipherByte)
	plainText = string(plainByte)
	return
}
func EncryptInt64(value int64, aesKey string) (cipherStr string, err error) {
	return Encrypt(strconv.FormatInt(value, 10), aesKey)
}
