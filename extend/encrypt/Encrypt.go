package encrypt

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

/**
notes: 加密工具
*/

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func SHA1(data interface{}) string {

	byt := []byte(fmt.Sprintf("%s", data))

	_sha := sha1.New()
	_sha.Write(byt)
	return hex.EncodeToString(_sha.Sum(nil))
}
