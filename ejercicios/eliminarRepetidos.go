package ejercicios

import "guia4/set"

func EliminarRepetidos[T comparable](arreglo []T) []T {

	conjunto := set.NewSet(arreglo...)

	return conjunto.Values()
}
