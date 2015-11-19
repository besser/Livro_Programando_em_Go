package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	expr := regexp.MustCompile("\\b\\w")

	transformadora := func(s string) string {
		return strings.ToUpper(s)
	}

	texto := "roberto mota besser"

	fmt.Println(transformadora(texto))
	fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
}
