package cola_prioridad

type fcmpHeap[T comparable] func(T, T) int

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

func CrearHeap[T comparable](cmp fcmpHeap[T]) ColaPrioridad[T] {
	return &heap[T]{make([]T, 1), 0, cmp}
}

func CrearHeapArr[T comparable](arr []T, cmp fcmpHeap[T]) ColaPrioridad[T] {
	heap := &heap[T]{make([]T, 10), 0, cmp}
	heap.datos = heapify[T](arr, cmp)
	heap.cantidad = len(arr)
	return heap
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(elem T) {
	heap.cantidad++
}

func (heap *heap[T]) Desencolar() T {
	heap.cantidad--
	panic("")
}

func (heap *heap[T]) VerMax() T {
	return heap.datos[0]
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func HeapSort[T comparable](elementos []T, cmp fcmpHeap[T]) {

	arr_heap := heapify[T](elementos, cmp)

	for i := len(arr_heap) - 1; i > 0; i-- {
		Swap[T](arr_heap, 0, i)
		elementos[i] = arr_heap[i]

		arr_heap = arr_heap[:i]

		heapDown[T](arr_heap, 0, cmp)
	}
	elementos[0] = arr_heap[0]
}

func heapify[T comparable](arr []T, cmp fcmpHeap[T]) []T {
	n := len(arr)
	heap := make([]T, n)
	copy(heap, arr)

	for i := (n / 2) - 1; i >= 0; i-- {
		heapDown[T](heap, i, cmp)
	}

	return heap
}

func heapUp[T comparable](arr []T, hijo int, cmp fcmpHeap[T]) {
}

func heapDown[T comparable](arr []T, padre int, cmp fcmpHeap[T]) {
	if padre < -1 || padre > len(arr)/2 {
		return
	}

	if padre == -1 {
		padre = 0
	}

	var hijo_izq int
	if (2*padre + 1) >= len(arr) {
		hijo_izq = padre
	} else {
		hijo_izq = 2*padre + 1
	}

	var hijo_der int
	if (2*padre + 2) >= len(arr) {
		hijo_der = padre
	} else {
		hijo_der = 2*padre + 2
	}

	var max int = padre
	if cmp(arr[hijo_izq], arr[max]) > 0 {
		max = hijo_izq
	}
	if cmp(arr[hijo_der], arr[max]) > 0 {
		max = hijo_der
	}

	if max != padre {
		Swap[T](arr, max, padre)
		heapDown[T](arr, max, cmp)
	}
}

func Swap[T comparable](arr []T, n1 int, n2 int) {
	arr[n1], arr[n2] = arr[n2], arr[n1]
}
