package main

import (
	"fmt"
	"guia4/ejercicios"
	"guia4/set"
)

func main() {
	// Creo un conjunto de letras
	// letras := ejercicios.Letras("Hola Mundo")
	// fmt.Println(letras)
	// arreglo := []int{1, 4, 5, 2, 1, 4, 5, 6}
	// fmt.Println(ejercicios.EliminarRepetidos(arreglo))
	// Creo dos conjuntos de números
	n1 := set.NewSet(1, 10, 5)
	n2 := set.NewSet(5, 15, 1)
	n3 := set.NewSet(10, 0, 2)
	fmt.Println(n1)
	fmt.Println("Diferencia Simetrica: ", ejercicios.DiferenciaSimetrica(n1, n2))
	fmt.Println("Diferencia: ", set.Difference(n1, n2))

	fmt.Println("Intersección: ", ejercicios.Interseccion(n1, n2, n3))
}
