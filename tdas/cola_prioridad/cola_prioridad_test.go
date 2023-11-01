package cola_prioridad_test

import (
	ColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const valor1 = 1
const valor2 = 2
const valor3 = 3
const valor4 = 4
const valor5 = 5
const valor6 = 6
const valor7 = 7

func compararInt(s1, s2 int) int {
	if s1 < s2 {
		return -1
	}
	if s1 > s2 {
		return 1
	}
	return 0
}

func TestColaVacia(t *testing.T) {
	heap := ColaPrioridad.CrearHeap(compararInt)
	require.True(t, heap.EstaVacia())
}

func TestColaVaciaArreglo(t *testing.T) {
	heap := ColaPrioridad.CrearHeapArr([]int{}, compararInt)
	require.True(t, heap.EstaVacia())
}

func TestEncolarYDesencolar(t *testing.T) {
	heap := ColaPrioridad.CrearHeap[int](compararInt)
	heap.Encolar(valor1)
	heap.Encolar(valor2)
	heap.Encolar(valor3)
	require.EqualValues(t, valor3, heap.Desencolar())
	require.EqualValues(t, valor2, heap.Desencolar())
	require.EqualValues(t, valor1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestDuplicadosHeapArreglo(t *testing.T) {
	heap := ColaPrioridad.CrearHeapArr([]int{valor1, valor2, valor1, valor3, valor2}, compararInt)
	require.EqualValues(t, valor3, heap.Desencolar())
	require.EqualValues(t, valor2, heap.Desencolar())
	require.EqualValues(t, valor2, heap.Desencolar())
	require.EqualValues(t, valor1, heap.Desencolar())
	require.EqualValues(t, valor1, heap.Desencolar())
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.Desencolar() })
}

func TestVerMaximoHeapArreglo(t *testing.T) {
	heap := ColaPrioridad.CrearHeapArr([]int{valor1, valor2, valor3, valor4}, compararInt)
	require.EqualValues(t, valor4, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, valor3, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, valor2, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, valor1, heap.VerMax())
	heap.Desencolar()
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.VerMax() })
}

func TestCantidadHeapArreglo(t *testing.T) {
	heap := ColaPrioridad.CrearHeapArr([]int{valor4, valor5, valor6, valor7}, compararInt)
	require.Equal(t, 4, heap.Cantidad())
	heap.Desencolar()
	require.Equal(t, 3, heap.Cantidad())
	heap.Desencolar()
	require.Equal(t, 2, heap.Cantidad())
	heap.Desencolar()
	require.Equal(t, 1, heap.Cantidad())
	heap.Desencolar()
	require.Equal(t, 0, heap.Cantidad())
}

func TestHeapify(t *testing.T) {
	arr := []int{4, 6, 5, 7, 1, 2, 3}
	heap := ColaPrioridad.CrearHeapArr(arr, compararInt)
	for i := valor7; i > 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	arr := []int{valor7, valor4, valor1, valor3, valor5, valor2, valor6}
	arr_ordenado := []int{valor1, valor2, valor3, valor4, valor5, valor6, valor7}
	ColaPrioridad.HeapSort(arr, compararInt)
	require.Equal(t, arr_ordenado, arr)
}
