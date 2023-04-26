package ejercicios

import (
	"guia4/set"
	"fmt"
)

func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {

	interseccion := conjuntos[0]
	fmt.Println(interseccion)

	for i := 1; i < len(conjuntos); i++ {
		
		fmt.Println(conjuntos[i])
		interseccion = set.Intersection(interseccion, conjuntos[i])
		fmt.Println(interseccion)
	}

	return interseccion
}
