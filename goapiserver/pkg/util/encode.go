// Go Api server
// @jeffotoni
// 2019-01-04

package util

import (
	"encoding/base64"
	"fmt"
	"log"
)

func EncodeBase64(str string) string {

	data := []byte(str)
	str64 := base64.StdEncoding.EncodeToString(data)
	return str64
}

// decode base64
func DecodeBase64(str64 string) string {

	data, err := base64.StdEncoding.DecodeString(str64)
	if err != nil {
		log.Println("error base64:", err, " Dado: ", str64)
		return str64
	}

	// use bufio write
	return fmt.Sprintf("%q", data)
}
