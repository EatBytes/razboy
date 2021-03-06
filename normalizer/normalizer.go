package normalizer

import (
	"encoding/base64"

	"github.com/eatbytes/razboy/ferror"
)

func Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Decode(str string) (string, error) {
	var (
		sDec []byte
		err  error
	)

	sDec, err = base64.StdEncoding.DecodeString(str)

	if err != nil {
		return str, ferror.NormalizeErr(err)
	}

	return string(sDec), nil
}

func PHPEncode(str string) string {
	return "base64_encode(" + str + ")"
}

func PHPDecode(str string) string {
	return "base64_decode(" + str + ")"
}
