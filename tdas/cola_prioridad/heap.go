package cola_prioridad

const (
	FACTOR_AUMENTO      = 2
	FACTOR_REDUCCION    = 2
	CONDICION_REDUCCION = 4
	CAPACIDAD_INICIAL   = 1
	PANIC_VACIA         = "La cola esta vacia"
)

type fcmpHeap[T comparable] func(T, T) int

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

func CrearHeap[T comparable](cmp fcmpHeap[T]) ColaPrioridad[T] {
	return &heap[T]{make([]T, CAPACIDAD_INICIAL), 0, cmp}
}

func CrearHeapArr[T comparable](arr []T, cmp fcmpHeap[T]) ColaPrioridad[T] {
	if len(arr) > 0 {
		aux := make([]T, len(arr))
		copy(aux, arr)
		heap := &heap[T]{aux, len(aux), cmp}
		heap.heapify()
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
	heap.redimencionar()
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic(PANIC_VACIA)
	}
	max := heap.datos[0]
	heap.datos[0] = heap.datos[heap.cantidad-1]
	heap.cantidad--
	heap.downHeap(0)
	heap.redimencionar()
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

func HeapSort[T comparable](elementos []T, cmp fcmpHeap[T]) {
	heapify := &heap[T]{elementos, len(elementos), cmp}
	heapify.heapify()
	for i := len(heapify.datos) - 1; i > 0; i-- {
		heapify.datos[0], heapify.datos[i] = heapify.datos[i], heapify.datos[0]
		heapify.cantidad--
		heapify.downHeap(0)
	}
}

func (heap *heap[T]) heapify() {
	for i := heap.cantidad; i >= 0; i-- {
		heap.downHeap(i)
	}
}

func (heap *heap[T]) redimencionar() {
	capacidadNueva := cap(heap.datos)
	if cap(heap.datos) == heap.cantidad {
		capacidadNueva *= FACTOR_AUMENTO
	} else if (heap.cantidad * CONDICION_REDUCCION) <= cap(heap.datos) {
		capacidadNueva /= FACTOR_REDUCCION
	} else {
		return
	}
	slice := make([]T, capacidadNueva)
	copy(slice, heap.datos)
	heap.datos = slice
}

func (heap *heap[T]) upHeap(padre int) {
	if padre == 0 {
		return
	}
	indice_padre := (padre - 1) / 2
	if heap.cmp(heap.datos[indice_padre], heap.datos[padre]) < 0 {
		heap.datos[indice_padre], heap.datos[padre] = heap.datos[padre], heap.datos[indice_padre]
		heap.upHeap(indice_padre)
	}
}

func (heap *heap[T]) downHeap(padre int) {
	izquierdo := 2*padre + 1
	derecho := 2*padre + 2
	max := padre

	if izquierdo < heap.cantidad && heap.cmp(heap.datos[izquierdo], heap.datos[max]) > 0 {
		max = izquierdo
	}
	if derecho < heap.cantidad && heap.cmp(heap.datos[derecho], heap.datos[max]) > 0 {
		max = derecho
	}

	if max != padre {
		heap.datos[padre], heap.datos[max] = heap.datos[max], heap.datos[padre]
		heap.downHeap(max)
	}
}

func swap[T comparable](arr []T, n1 int, n2 int) {
	arr[n1], arr[n2] = arr[n2], arr[n1]
}
