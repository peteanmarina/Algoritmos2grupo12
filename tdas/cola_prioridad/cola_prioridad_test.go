package cola_prioridad_test

import (
	ColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	numero1 = 1
	numero2 = 2
	numero3 = 3
	numero4 = 4
	numero5 = 5
	numero6 = 6
	numero7 = 7
	volumen = 1000
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

func TestColaVacia(t *testing.T) {
	heap := ColaPrioridad.CrearHeap(compararInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.Desencolar() })
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.VerMax() })
}

func TestColaVaciaArreglo(t *testing.T) {
	heap := ColaPrioridad.CrearHeapArr([]int{}, compararInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.Desencolar() })
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.VerMax() })
}

func TestEncolarYDesencolar(t *testing.T) {
	heap := ColaPrioridad.CrearHeap[int](compararInt)
	heap.Encolar(numero1)
	heap.Encolar(numero2)
	heap.Encolar(numero3)
	require.EqualValues(t, numero3, heap.Desencolar())
	require.EqualValues(t, numero2, heap.Desencolar())
	require.EqualValues(t, numero1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestVolumenEncolarYDesencolar(t *testing.T) {
	heap := ColaPrioridad.CrearHeap[int](compararInt)
	for i := 0; i <= volumen; i++ {
		heap.Encolar(i)
	}
	for i := volumen; i >= 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
	}
}

func TestDuplicadosHeapArreglo(t *testing.T) {
	heap := ColaPrioridad.CrearHeapArr([]int{numero1, numero2, numero1, numero3, numero2}, compararInt)
	require.EqualValues(t, numero3, heap.Desencolar())
	require.EqualValues(t, numero2, heap.Desencolar())
	require.EqualValues(t, numero2, heap.Desencolar())
	require.EqualValues(t, numero1, heap.Desencolar())
	require.EqualValues(t, numero1, heap.Desencolar())
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.Desencolar() })
}

func TestVerMaximoHeapArreglo(t *testing.T) {
	heap := ColaPrioridad.CrearHeapArr([]int{numero1, numero2, numero3, numero4}, compararInt)
	require.EqualValues(t, numero4, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, numero3, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, numero2, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, numero1, heap.VerMax())
	heap.Desencolar()
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.VerMax() })
}

func TestCantidadHeapArreglo(t *testing.T) {
	heap := ColaPrioridad.CrearHeapArr([]int{numero4, numero5, numero6, numero7}, compararInt)
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
	for i := numero7; i > 0; i-- {
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	arr := []int{numero7, numero4, numero1, numero3, numero5, numero2, numero6}
	arr_ordenado := []int{numero1, numero2, numero3, numero4, numero5, numero6, numero7}
	ColaPrioridad.HeapSort(arr, compararInt)
	require.Equal(t, arr_ordenado, arr)
}
