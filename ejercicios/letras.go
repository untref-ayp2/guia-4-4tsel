package ejercicios

import (
	"github.com/agrison/go-commons-lang/stringUtils"
	"guia4/set"
)

func Letras(s string) *set.Set[string] {

	result := set.NewSet[string]()

	for _, v := range s {

		letra := string(v)

		if !stringUtils.IsBlank(letra) {

			result.Add(letra)
		}

	}

	return result
}
