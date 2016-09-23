package main

import (
	"api-foundations/bootstrap"
	"fmt"
)

func main() {
	fmt.Printf("Time: %.4f\n", bootstrap.Now())
	fmt.Println("Hello world!")
	fmt.Printf("Time: %.4f\n", bootstrap.Now())

	bootstrap.Reset()
	fmt.Printf("Time after reset: %.4f\n", bootstrap.Now())
}
