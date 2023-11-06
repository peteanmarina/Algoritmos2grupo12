package cola_prioridad

const (
	FACTOR_AUMENTO      = 2
	FACTOR_REDUCCION    = 2
	CONDICION_REDUCCION = 4
	CAPACIDAD_INICIAL   = 13
	PANIC_VACIA         = "La cola esta vacia"
)

type fcmpHeap[T any] func(T, T) int

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

func CrearHeap[T any](cmp fcmpHeap[T]) ColaPrioridad[T] {
	return &heap[T]{make([]T, CAPACIDAD_INICIAL), 0, cmp}
}

func CrearHeapArr[T any](arr []T, cmp fcmpHeap[T]) ColaPrioridad[T] {
	if len(arr) > 0 {
		aux := make([]T, len(arr))
		copy(aux, arr)
		heap := &heap[T]{aux, len(aux), cmp}
		heapify(aux, cmp)
		copy(heap.datos, aux)
		return heap
	}
	return CrearHeap(cmp)
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(dato T) {
	heap.datos[heap.cantidad] = dato
	heap.upHeap(heap.cantidad)
	heap.cantidad++
	heap.redimensionar()
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(PANIC_VACIA)
	}
	max := heap.datos[0]
	heap.datos[0] = heap.datos[heap.cantidad-1]
	heap.cantidad--
	downHeap(heap.datos, 0, heap.cantidad, heap.cmp)
	heap.redimensionar()
	return max
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic(PANIC_VACIA)
	}
	return heap.datos[0]
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func HeapSort[T any](elementos []T, cmp fcmpHeap[T]) {
	heapify(elementos, cmp)
	n := 0
	for i := len(elementos) - 1; i > 0; i-- {
		swap(elementos, 0, i)
		downHeap(elementos, 0, i, cmp)
		n++
	}
}

func heapify[T any](arreglo []T, cmp fcmpHeap[T]) {
	for i := len(arreglo) - 1; i >= 0; i-- {
		downHeap(arreglo, i, len(arreglo), cmp)
	}
}

func (heap *heap[T]) redimensionar() {
	capacidadNueva := cap(heap.datos)
	if cap(heap.datos) == heap.cantidad {
		capacidadNueva *= FACTOR_AUMENTO
	} else if (heap.cantidad * CONDICION_REDUCCION) <= cap(heap.datos) {
		capacidadNueva /= FACTOR_REDUCCION
	} else {
		return
	}
	if capacidadNueva < CAPACIDAD_INICIAL {
		capacidadNueva = CAPACIDAD_INICIAL
	}
	slice := make([]T, capacidadNueva)
	copy(slice, heap.datos)
	heap.datos = slice
}

func (heap *heap[T]) upHeap(indice_elemento int) {
	if indice_elemento == 0 {
		return
	}
	indice_padre := (indice_elemento - 1) / 2
	if heap.cmp(heap.datos[indice_padre], heap.datos[indice_elemento]) < 0 {
		swap(heap.datos, indice_padre, indice_elemento)
		heap.upHeap(indice_padre)
	}
}

func downHeap[T any](arreglo []T, indice_elemento int, cantidad int, cmp fcmpHeap[T]) {
	izquierdo := 2*indice_elemento + 1
	derecho := 2*indice_elemento + 2
	max := indice_elemento

	if izquierdo < cantidad && cmp(arreglo[izquierdo], arreglo[max]) > 0 {
		max = izquierdo
	}
	if derecho < cantidad && cmp(arreglo[derecho], arreglo[max]) > 0 {
		max = derecho
	}

	if max != indice_elemento {
		arreglo[indice_elemento], arreglo[max] = arreglo[max], arreglo[indice_elemento]
		downHeap(arreglo, max, cantidad, cmp)
	}
}

func swap[T any](arr []T, n1 int, n2 int) {
	arr[n1], arr[n2] = arr[n2], arr[n1]
}
