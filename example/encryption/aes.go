package main

import (
	"fmt"

	"github.com/ea67/gocore/encryption/aes"
)

func main() {

	str, _ := aes.AesEncrypt("sunmi", "sunmiWorkOnesunmiWorkOne")
	fmt.Println(str)
	str2, _ := aes.AesDecrypt(str, "sunmiWorkOnesunmiWorkOne")
	fmt.Println(str2)
}
