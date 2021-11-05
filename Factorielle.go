package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("donnez une valeur à x")
	fmt.Scanln(&x)
	y = 1

	for i := 1; i <= x; i++ {
		y = y * i
	}
	fmt.Println("le factoriel de", x, "est", y)

}
