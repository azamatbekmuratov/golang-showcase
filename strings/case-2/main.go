//concatenete strings
package main

import (
	"bytes"
)

func main() {
	var b bytes.Buffer

	b.WriteString("test")
	b.WriteString("test continue")

	fmt.Println("String: ", b.String())

	//or use fmt.Sprintf()
	str1 := "test"
	str2 := "test2"
	result := fmt.Sprintf("%s,%s", str1, str2)
	fmt.Println(result)
}
