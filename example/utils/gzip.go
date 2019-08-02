package main

import (
	"fmt"
	"github.com/ea67/gocore/utils"
)

func main() {
	fmt.Println(utils.GzipEncode("dsxdjdhskfjkdsfhsdjlaal"))
	var m = utils.GzipEncode("dsxdjdhskfjkdsfhsdjlaal")
	fmt.Println(utils.GzipDecode(m))
}
