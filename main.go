package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 != 0:
			fmt.Println("fizz")
		case i%5 == 0 && i%3 != 0:
			fmt.Println("buzz")
		case i%5 == 0 && i%3 == 0:
			fmt.Println("fizzbuzz")
		default:
			fmt.Println(i)
		}
	}

}
