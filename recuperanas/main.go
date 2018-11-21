package main

import "fmt"

func main() {
	listalistanas := recuperaNAS()

	var i int
	for _, listanas := range listalistanas {
		for n, nas := range listanas {
			i++
			fmt.Println(n, nas.Name)
		}
	}
	fmt.Printf("I Nas trovati sono %d\n", i)
}
