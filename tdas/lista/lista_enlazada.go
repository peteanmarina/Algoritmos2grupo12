package lista

const (
	mensaje_lista_vacia  = "La lista esta vacia"
	mensaje_no_elementos = "No hay mas elementos en la lista"
	mensaje_iterador     = "El iterador termino de iterar"
)

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
	principio *nodo[T]
	fin       *nodo[T]
	largo     int
}

type iterador[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *listaEnlazada[T]
}

func crearNodo[T any](elemento T, sig *nodo[T]) *nodo[T] {
	return &nodo[T]{dato: elemento, siguiente: sig}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) InsertarPrimero(elememt T) {
	nuevo_nodo := crearNodo[T](elememt, l.principio)
	if l.EstaVacia() {
		l.fin = nuevo_nodo
	}
	l.principio = nuevo_nodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elememt T) {
	nuevo_nodo := crearNodo[T](elememt, nil)
	if l.EstaVacia() {
		l.principio = nuevo_nodo
	} else {
		l.fin.siguiente = nuevo_nodo
	}
	l.fin = nuevo_nodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic(mensaje_lista_vacia)
	}
	dato_a_retornar := l.principio.dato
	l.principio = l.principio.siguiente
	if l.principio == nil {
		l.fin = l.principio
	}
	l.largo--
	return dato_a_retornar
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic(mensaje_lista_vacia)
	}
	return l.principio.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic(mensaje_lista_vacia)
	}
	return l.fin.dato
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterador[T]{l.principio, nil, l}
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.principio
	for i := 0; i < l.largo; i++ {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

func (i *iterador[T]) VerActual() T {
	if !i.HaySiguiente() {
		panic(mensaje_iterador)
	}
	return i.actual.dato
}

func (i *iterador[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iterador[T]) Siguiente() {
	if !i.HaySiguiente() {
		panic(mensaje_iterador)
	}
	i.anterior = i.actual
	i.actual = i.actual.siguiente
}

func (i *iterador[T]) Insertar(element T) {
	nuevo_nodo := crearNodo[T](element, i.actual)
	if i.actual == nil {
		i.lista.fin = nuevo_nodo
	}
	if i.anterior == nil {
		i.actual = nuevo_nodo
		i.lista.principio = i.actual
		if i.lista.fin == nil {
			i.lista.fin = i.actual
		}
	} else {
		i.anterior.siguiente = nuevo_nodo
		i.actual = i.anterior.siguiente
	}
	i.lista.largo++
}

func (i *iterador[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic(mensaje_iterador)
	}
	if i.anterior != nil {
		i.anterior.siguiente = i.actual.siguiente
	}
	if i.actual == i.lista.fin {
		i.lista.fin = i.anterior
	}
	if i.actual == i.lista.principio {
		i.lista.principio = i.actual.siguiente
	}
	dato_a_retornar := i.actual.dato
	i.actual = i.actual.siguiente
	i.lista.largo--
	return dato_a_retornar
}
