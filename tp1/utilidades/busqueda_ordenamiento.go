package utilidades

import "rerepolez/votos"

const NO_ENCONTRADO = -1

func BusquedaBinaria(arreglo []votos.Votante, inicio, fin, elemento int) int {

	medio := (inicio + fin) / 2

	if inicio > fin || medio >= len(arreglo) || medio < 0 {
		return NO_ENCONTRADO
	}

	if arreglo[medio].LeerDNI() == elemento {
		return medio
	} else if arreglo[medio].LeerDNI() < elemento {
		return BusquedaBinaria(arreglo, medio+1, fin, elemento)
	} else {
		return BusquedaBinaria(arreglo, inicio, medio-1, elemento)
	}
}

func CountingSort(arreglo []votos.Votante, cifra int) []votos.Votante {
	n := len(arreglo)
	nuevo_arreglo := make([]votos.Votante, n)
	contadores := make([]int, 10)
	for i := 0; i < n; i++ {
		indice := (arreglo[i].LeerDNI() / cifra) % 10
		contadores[indice]++
	}
	acumuladores := make([]int, 10)
	for i := 1; i < 10; i++ {
		acumuladores[i] = acumuladores[i-1] + contadores[i-1]
	}
	for i := 0; i < n; i++ {
		indice := (arreglo[i].LeerDNI() / cifra) % 10
		nuevo_arreglo[acumuladores[indice]] = arreglo[i]
		acumuladores[indice]++
	}
	for i := 0; i < n; i++ {
		arreglo[i] = nuevo_arreglo[i]
	}
	return arreglo
}

func RadixSort(arr []votos.Votante, max int) []votos.Votante {
	for cifra := 1; max/cifra > 0; cifra *= 10 {
		arr = CountingSort(arr, cifra)
	}
	return arr
}
