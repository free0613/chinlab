package main

import (
	"fmt"
)

type shouldRecover struct{}
type emptyStruct struct{}

func f6() (err error) {
	defer func() {
		switch p := recover(); p {
		case nil: //donoting
		case shouldRecover{}:
			err = fmt.Errorf("occur panic but had recovered")
		default:
			panic(p)
		}
	}()

	func() {
		func() int {
			panic(shouldRecover{})
			//panic(emptyStruct{})
			x := 0
			y := 5 / x
			return y
		}()
	}()

	return
}

func main() {
	err := f6()
	if err != nil {
		fmt.Printf("fail %v\n", err)
	} else {
		fmt.Printf("success\n")
	}
}
