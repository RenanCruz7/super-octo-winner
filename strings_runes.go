package main

import (
	"fmt"
	"unicode/utf8"
)

/*
String is a sequence of bytes
A string can contain any bytes, including bytes that are not valid UTF-8. The range of valid Unicode code points is U+0000 to U+10FFFF.
The UTF-8 encoding of a Unicode code point is a sequence of one to four bytes.
*/
func main() {
	const s = "Hello,世界"

	fmt.Println("Len:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// count real characters
	fmt.Println(utf8.RuneCountInString(s))
}
