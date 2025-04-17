package encrypts

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5(pwd string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, pwd)
	return hex.EncodeToString(hash.Sum(nil))
}
