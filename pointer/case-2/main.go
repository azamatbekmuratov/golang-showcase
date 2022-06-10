// Playing with pointers in slice
package main

import "fmt"

func slc(input *[]*int) {
	for i, v := range *input {
		fmt.Println(i, v)
	}
}

func main() {
	frst := 5
	sec := 4
	prms := []*int{&frst, &sec}
	slc(&prms)

}
