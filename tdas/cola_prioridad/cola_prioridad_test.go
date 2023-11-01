package cola_prioridad_test

import (
	ColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararInt(s1, s2 int) int {
	if s1 < s2 {
		return -1
	}
	if s1 > s2 {
		return 1
	}
	return 0
}

func TestPruebaHeapify(t *testing.T) {
	arr := []int{4, 6, 5, 7, 8, 1, 9, 2, 3, 10}
	heap := ColaPrioridad.CrearHeapArr(arr, compararInt)

	require.EqualValues(t, 10, heap.Desencolar())
	require.EqualValues(t, 9, heap.Desencolar())
	require.EqualValues(t, 8, heap.Desencolar())
	require.EqualValues(t, 7, heap.Desencolar())
	require.EqualValues(t, 6, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 4, heap.Desencolar())
	require.EqualValues(t, 3, heap.Desencolar())
	require.EqualValues(t, 2, heap.Desencolar())
	require.EqualValues(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestPruebaHeapSort(t *testing.T) {
	arr := []int{4, 6, 5, 7, 8, 1, 9, 2, 3, 10}
	arr_ordenado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ColaPrioridad.HeapSort[int](arr, compararInt)

	require.Equal(t, arr_ordenado, arr)
}
