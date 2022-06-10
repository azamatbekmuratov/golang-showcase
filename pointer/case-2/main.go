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

	var a = []int{1, 2, 3, 4, 5}
	b := a[2:]
	b[0] = 0
	

}
