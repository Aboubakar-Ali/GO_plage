package main

import "fmt"

func main() {

	var n, max, pst, i int
	i = 0
	n = 1
	max = 0
	for n != 0 {
		fmt.Println("saisir un nombre")
		fmt.Scanln(&n)

		i = i + 1

		if i == 1 || n > max {
			max = n
			pst = i

		}
	}

	fmt.Println("le nombre le plus grand est", max, ",et sa position est", pst)

}
