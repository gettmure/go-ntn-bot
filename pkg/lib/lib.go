package lib

import "encoding/base64"

func StringToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
