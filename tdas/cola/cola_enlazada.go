package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato    T
	proximo *nodoCola[T]
}

func crearNodoCola[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato, nil}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{nil, nil}
}

func (cola *colaEnlazada[T]) Invertir_cola() {
	actual := cola.primero
	var anterior *nodoCola[T] = nil
	siguiente := cola.primero.proximo
	cola.ultimo = cola.primero
	for actual != nil {
		actual.proximo = anterior
		anterior = actual
		actual = siguiente
		if siguiente == actual {
			actual.proximo = nil
		}
	}
	cola.primero = anterior
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := crearNodoCola[T](dato)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.proximo = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	desencolado := cola.primero.dato
	cola.primero = cola.primero.proximo
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return desencolado
}
