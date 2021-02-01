package cli

import (
	"fmt"
	"os"
)

// Exit finishs requests
func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	fmt.Println(err)
	os.Exit(code)
}
