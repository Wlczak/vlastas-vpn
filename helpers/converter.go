package helpers

import (
	"encoding/base64"
	"encoding/hex"
)

func Base64tohex(base64String string) string {
	base64Bytes := make([]byte, base64.StdEncoding.DecodedLen(len(base64String)))
	n, err := base64.StdEncoding.Decode(base64Bytes, []byte(base64String))
	if err != nil {
		panic(err)
	}
	base64Bytes = base64Bytes[:n]

	hex := hex.EncodeToString(base64Bytes)

	return hex
}
