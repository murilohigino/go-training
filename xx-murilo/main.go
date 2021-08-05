package main

import "fmt"

func main(){

	x:=1
	y:=2
	fmt.Printf("\nValor de X: %v\nValor de y: %v", x, y)
	//OpEquals
}

func OpEquals(x int, y int) bool {
	  fmt.Println("x == y:", x == y)
	  return x == y
}