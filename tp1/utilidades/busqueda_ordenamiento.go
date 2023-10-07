package utilidades

const NO_ENCONTRADO = -1

func BusquedaBinaria(arreglo []int, inicio, fin, elemento int) int {
	if inicio > fin {
		return NO_ENCONTRADO
	}
	medio := (inicio + fin) / 2
	if arreglo[medio] == elemento {
		return medio
	} else if arreglo[medio] < elemento {
		return BusquedaBinaria(arreglo, medio+1, fin, elemento)
	} else {
		return BusquedaBinaria(arreglo, inicio, medio-1, elemento)
	}
}

func CountingSort(arreglo []int, cifra int) []int {
	n := len(arreglo)
	nuevo_arreglo := make([]int, n)
	contadores := make([]int, 10)
	for i := 0; i < n; i++ {
		indice := (arreglo[i] / cifra) % 10
		contadores[indice]++
	}
	acumuladores := make([]int, 10)
	for i := 1; i < 10; i++ {
		acumuladores[i] = acumuladores[i-1] + contadores[i-1]
	}
	for i := 0; i < n; i++ {
		indice := (arreglo[i] / cifra) % 10
		nuevo_arreglo[acumuladores[indice]] = arreglo[i]
		acumuladores[indice]++
	}
	for i := 0; i < n; i++ {
		arreglo[i] = nuevo_arreglo[i]
	}
	return arreglo
}

func RadixSort(arr []int, max int) []int {
	for cifra := 1; max/cifra > 0; cifra *= 10 {
		arr = CountingSort(arr, cifra)
	}
	return arr
}
