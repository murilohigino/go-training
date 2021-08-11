package main

import "fmt"

func main(){

	var name string
	var idade int

	fmt.Printf("\nDigite seu nome:\n")
	fmt.Scanf("%s", &name)

	fmt.Printf("Digite sua idade:\n")
	fmt.Scanf("%d", &idade)

	fmt.Printf("\nNome: %s\nidade: %d\n", name ,idade)

}