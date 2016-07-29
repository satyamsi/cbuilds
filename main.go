// +build linux,!darwin
package main 

import "C"
import "fmt"
import "github.com/satyamsi/cbuilds/a"

func main() {
	fmt.Println("Hello World")
        a.Somefunc()
}

