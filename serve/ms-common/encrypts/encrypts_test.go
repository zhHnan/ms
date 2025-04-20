package encrypts

import (
	"fmt"
	"hnz.com/ms_serve/ms-project/pkg/model"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plain := "100123213123"
	// AES 规定有3种长度的key: 16, 24, 32分别对应AES-128, AES-192, or AES-256
	key := "abcdefgehjhijkmlkjjwwoew"
	// 加密
	cipherByte, err := Encrypt(plain, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %s\n", plain, cipherByte)
	// 解密
	plainText, err := Decrypt(cipherByte, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %s\n", cipherByte, plainText)
}
func TestDecrypt(t *testing.T) {
	var code = "3ba2043c16"
	text, err := Decrypt(code, model.AESKey)
	fmt.Println(text, err)
}
