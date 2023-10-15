package utilidades

const NO_ENCONTRADO = -1

func IngresoBinario(elem int, array []int) []int {
	indice := buscar_ultimo_dni(elem, array)
	if indice == -1 {
		return append(array, elem)
	}
	retornado := make([]int, 0)
	retornado = append(retornado, array[:indice]...)
	retornado = append(retornado, elem)
	retornado = append(retornado, array[indice:]...)
	return retornado
}

func QuitaBinaria(elem int, array []int) []int {
	indice := BusquedaBinaria(array, 0, len(array), elem)

	if indice == NO_ENCONTRADO {
		return array
	}

	retornado := make([]int, 0)

	if indice == 0 {
		der := array[indice+1:]
		retornado = append(retornado, der...)
	} else if indice == len(array)-1 {
		izq := array[:indice]
		retornado = append(retornado, izq...)
	} else {
		izq := array[:indice]
		der := array[indice+1:]
		retornado = append(retornado, izq...)
		retornado = append(retornado, der...)
	}

	return retornado
}

func buscar_ultimo_dni(elem int, array []int) int {
	izq, der := 0, len(array)-1

	for izq <= der {
		medio := izq + (der-izq)/2

		if array[medio] == elem {
			return medio
		} else if array[medio] < elem {
			izq = medio + 1
		} else {
			der = medio - 1
		}
	}

	return izq
}

func BusquedaBinaria(arreglo []int, inicio, fin, elemento int) int {

	medio := (inicio + fin) / 2

	if inicio > fin || medio >= len(arreglo) || medio < 0 {
		return NO_ENCONTRADO
	}

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
