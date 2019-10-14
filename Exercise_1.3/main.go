// Exercise 1.3
package main

import (
	"strings"
)

func main() {
	echo1([]string{"o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o"})

	echo2([]string{"o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o"})

	echo3([]string{"o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o"})
}

func echo1(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[0:] {
		s += sep + arg
		sep = " "
	}
}

func echo3(args []string) {
	strings.Join(args[0:], " ")
}
