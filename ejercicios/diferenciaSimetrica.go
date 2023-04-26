package ejercicios

import (
	"guia4/set"
)

func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {
	
	set3 := set.NewSet[T]()

	for _, v := range s1.Values() {

		if !s2.Contains(v){

			set3.Add(v)
		}
	}

	for _, v := range s2.Values() {
		
		if !s1.Contains(v){

			set3.Add(v)
		}
	}

	return set3
}
