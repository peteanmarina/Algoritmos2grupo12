package cola_prioridad_test

import (
	ColaPrioridad "tdas/cola_prioridad"
	"testing"

	"strings"

	"github.com/stretchr/testify/require"
)

const (
	numero1  = 1
	numero2  = 2
	numero3  = 3
	numero4  = 4
	numero5  = 5
	numero6  = 6
	numero7  = 7
	volumen  = 10000
	palabra1 = "Arbol"
	palabra2 = "Burro"
	palabra3 = "Casa"
	palabra4 = "Diario"
	palabra5 = "Sandia"
	palabra6 = "Walter"
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

type personas struct {
	edad   int
	nombre string
}

func compararPersonasEdad(p1, p2 personas) int {
	return compararInt(p1.edad, p2.edad)
}

func TestColaVacia(t *testing.T) {
	t.Log("Heap vacio se comporta como tal")
	heap := ColaPrioridad.CrearHeap(compararInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.Desencolar() })
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.VerMax() })
}

func TestComportamientoVaciarCola(t *testing.T) {
	t.Log("Cola desencolada hasta quedar vacia se comporta como cola vacia")
	heap := ColaPrioridad.CrearHeap(strings.Compare)
	heap.Encolar(palabra1)
	heap.Desencolar()
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.Desencolar() })
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.VerMax() })
}

func TestColaVaciaArreglo(t *testing.T) {
	t.Log("Crear heap desde arreglo vacio es un heap vacio")
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

func compararStrAlReves(s1, s2 string) int { return -strings.Compare(s1, s2) }

func TestHojasInpares(t *testing.T) {
	t.Log("Encolar y desencolar heap con hojas impares en ultimo nivel")
	heap := ColaPrioridad.CrearHeap[string](compararStrAlReves)
	heap.Encolar(palabra1)
	heap.Encolar(palabra3)
	heap.Encolar(palabra4)
	heap.Encolar(palabra2)
	require.EqualValues(t, palabra1, heap.VerMax())
	require.EqualValues(t, palabra1, heap.Desencolar())
	require.EqualValues(t, palabra2, heap.VerMax())
	require.EqualValues(t, palabra2, heap.Desencolar())
	require.EqualValues(t, palabra3, heap.VerMax())
	require.EqualValues(t, palabra3, heap.Desencolar())
	require.EqualValues(t, palabra4, heap.VerMax())
	require.EqualValues(t, palabra4, heap.Desencolar())
}

func TestVolumenEncolarYDesencolar(t *testing.T) {
	t.Log("Encolar y desencolar gran numero de elementos")
	heap := ColaPrioridad.CrearHeap[int](compararInt)
	for i := 0; i <= volumen; i++ {
		heap.Encolar(i)
	}
	for i := volumen; i >= 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
	}
}

func TestDuplicadosHeapArreglo(t *testing.T) {
	t.Log("Encolar elementos duplicados. Heap cumple propiedad de heap")
	heap := ColaPrioridad.CrearHeapArr([]int{numero1, numero2, numero1, numero3, numero2}, compararInt)
	require.EqualValues(t, numero3, heap.Desencolar())
	require.EqualValues(t, numero2, heap.Desencolar())
	require.EqualValues(t, numero2, heap.Desencolar())
	require.EqualValues(t, numero1, heap.Desencolar())
	require.EqualValues(t, numero1, heap.Desencolar())
	require.PanicsWithValue(t, ColaPrioridad.PANIC_VACIA, func() { heap.Desencolar() })
}

func TestVerMaximoHeapArreglo(t *testing.T) {
	t.Log("Siempre se muestra el maximo del arreglo")
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
	t.Log("Se actualiza la cantidad correctamente")
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
	t.Log("Heapify convierte en un heap a arreglo")
	arr := []int{4, 6, 5, 7, 1, 2, 3}
	heap := ColaPrioridad.CrearHeapArr(arr, compararInt)
	for i := numero7; i > 0; i-- {
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	t.Log("Heap Sort menor a mayor")
	arr := []int{numero7, numero4, numero1, numero3, numero5, numero2, numero6}
	arr_ordenado := []int{numero1, numero2, numero3, numero4, numero5, numero6, numero7}
	ColaPrioridad.HeapSort(arr, compararInt)
	require.Equal(t, arr_ordenado, arr)
}

func compararIntAlReves(s1, s2 int) int { return -compararInt(s1, s2) }

func TestHeapSortAlreves(t *testing.T) {
	t.Log("Heap Sort mayor a menor")
	arr := []int{numero7, numero4, numero1, numero3, numero5, numero2, numero6}
	arr_ordenado := []int{numero7, numero6, numero5, numero4, numero3, numero2, numero1}
	ColaPrioridad.HeapSort(arr, compararIntAlReves)
	require.Equal(t, arr_ordenado, arr)
}

func TestHeapSortVacio(t *testing.T) {
	t.Log("Heap Sort en heap vacio")
	arr := []int{}
	ColaPrioridad.HeapSort(arr, compararInt)
	require.Equal(t, []int{}, arr)
}

func TestCrearColaConArreglo(t *testing.T) {
	t.Log("Con struct de personas y mismos elementos Heap desde arreglo y heap vacio al que encolamos, desencolan lo mismo")
	marcos := personas{edad: numero2, nombre: "Marcos"}
	juan := personas{edad: numero7, nombre: "Juan"}
	pedro := personas{edad: numero6, nombre: "Pedro"}
	adrian := personas{edad: numero4, nombre: "Adrian"}
	marta := personas{edad: numero5, nombre: "Marta"}
	josefa := personas{edad: numero2, nombre: "Josefa"}
	merlina := personas{edad: numero1, nombre: "Merlina"}
	arr := []personas{marcos, juan, pedro, adrian, marta, josefa, merlina}

	arr_heap := ColaPrioridad.CrearHeap[personas](compararPersonasEdad)
	arr_heap.Encolar(marcos)
	arr_heap.Encolar(juan)
	arr_heap.Encolar(pedro)
	arr_heap.Encolar(adrian)
	arr_heap.Encolar(marta)
	arr_heap.Encolar(josefa)
	arr_heap.Encolar(merlina)
	heap := ColaPrioridad.CrearHeapArr(arr, compararPersonasEdad)

	for i := 0; i < len(arr); i++ {
		require.EqualValues(t, arr_heap.Desencolar(), heap.Desencolar())
	}

}
