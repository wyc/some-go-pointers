package main

import (
	"fmt"
	"time"
)

type A struct {
	data [65536]byte
}

func valFunc(a A) int {
	return len(a.data)
}

func ptrFunc(a *A) int {
	return len(a.data)
}

type B struct {
	data [1]byte
}

func valFuncSmol(b B) int {
	return len(b.data)
}

func ptrFuncSmol(b *B) int {
	return len(b.data)
}

func main() {
	a := A{}
	t := time.Now()
	for i := 0; i < 1000000; i++ {
		valFunc(a)
	}
	fmt.Println("A: passing the val:", time.Now().Sub(t))

	t = time.Now()
	for i := 0; i < 1000000; i++ {
		ptrFunc(&a)
	}
	fmt.Println("A: passing the ptr:", time.Now().Sub(t))

	b := B{}
	t = time.Now()
	for i := 0; i < 1000000; i++ {
		valFuncSmol(b)
	}
	fmt.Println("B: passing the val:", time.Now().Sub(t))

	t = time.Now()
	for i := 0; i < 1000000; i++ {
		ptrFuncSmol(&b)
	}
	fmt.Println("B: passing the ptr:", time.Now().Sub(t))
}
