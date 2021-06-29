package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	oo, _ := ioutil.ReadAll(os.Stdin)
	x0 := string(oo)
	re := regexp.MustCompile(`(\x60.*?\x60|\x22[^\n"]*\x22|/\*.*?\*/|\bpackage|import|func|interface|if|else|for|range\b)`)
	fmt.Println(re.ReplaceAllStringFunc(x0, func(x string) string {
		switch x[0] {
		case 0x22:
			return "\x1b[36m" + x + "\x1b[0m"
		case 0x60:
			return "\x1b[36m" + x + "\x1b[0m"
		case 0x2f:
			return "\x1b[90m" + x + "\x1b[0m"
		default:
			return "\x1b[0;1m" + x + "\x1b[0m"
		}
	}))
}
