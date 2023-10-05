package pila

const cantidadAumento = 2
const cantidadReduccion = 2
const condicionReduccion = 4

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{make([]T, 1), 0}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(dato T) {
	pila.datos[pila.cantidad] = dato
	pila.cantidad++
	pila.modificarCapacidad()
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	pila.modificarCapacidad()
	pila.cantidad--
	return pila.datos[pila.cantidad]
}

func (pila *pilaDinamica[T]) modificarCapacidad() {
	capacidadNueva := cap(pila.datos)
	if cap(pila.datos) == pila.cantidad {
		capacidadNueva *= cantidadAumento
	} else if (pila.cantidad * condicionReduccion) <= cap(pila.datos) {
		capacidadNueva /= cantidadReduccion
	}
	slice := make([]T, capacidadNueva)
	copy(slice, pila.datos)
	pila.datos = slice
}

func (p *pilaDinamica[T]) Invertir() {
	for i := 0; i < p.cantidad/2; i++ {
		p.datos[i], p.datos[p.cantidad-i-1] = p.datos[p.cantidad-i-1], p.datos[i]
	}
}
