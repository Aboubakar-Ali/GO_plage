package main

import "fmt"

func main() {
	var n, max, pos int
	max = 0
	for i := 1; i < 11; i++ {
		fmt.Println("entrez un nombre")
		fmt.Scan(&n)
		if i == 1 || n > max {
			max = n
			pos = i
		}
	}
	fmt.Println("le nombre le plus grand est", max, ", sa position de saisi est", pos)
}
