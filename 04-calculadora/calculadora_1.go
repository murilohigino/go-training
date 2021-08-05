package main

import (
	"fmt"
	"math"
)

func main() {

		fmt.Println("Soma: ", soma(4 , 5))
		fmt.Println("subtração: ", subtracao(3, 5))
		fmt.Println("Multiplicação: ", mult(4, 6))
		fmt.Println("Divisão: ", div(6, 2))
		fmt.Println("porcentagem: ",addpercentual(50 , 30) )
		fmt.Println("Quanto porcento um representa do outro : ", achepercentual(30, 50))
		fmt.Println("Potencia: ", potencia(10 , 2))

}

func soma(x float64, y float64) float64{
	return x + y
}

func subtracao(x float64, y float64)  float64{
	return x - y
}

func mult(x float64, y float64)  float64{
	return x * y
}

func div(x float64, y float64)  float64{
	return x / y
}

func addpercentual(x float64, y float64)  float64{
	return x + (x  *  y / 100)
}

func achepercentual(x float64, y float64)  float64{
	return (x / y * 100)
}

func potencia(x float64, y float64)  float64{
	return math.Pow( x , y)
}