package lista

type Lista[T any] interface {
	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool
	// InsertarPrimero agrega un nuevo elemento a la primer posicion de la lista.
	InsertarPrimero(T)
	// InsertarUltimo agrega un nuevo elemento a la ultima posicion de la lista.
	InsertarUltimo(T)
	// BorrarPrimero elimina el primer elemento de la lista y lo duelve. Si la lista esta vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T
	// VerPrimero devuelve el primer elemento de la lista. Si la lista esta vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T
	// VerUltimo devuelve el ultimo elemento de la lista. Si la lista esta vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T
	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int
	// Iterar itera internamente en la lista aplicandole a cada elemento la funcion visitar recibida por parametro.
	Iterar(visitar func(T) bool)
	// Iterador devuelve un iterador externo para recorrer la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// VerActual devuelve el elemento que esta siendo apuntado por el iterador. Si ya no hay elementos entra en pánico con el mensaje "El iterador termino de iterar"
	VerActual() T
	// HaySiguiente devuelve true si el iterador apunta a un elemento de la lista, false en caso contrario.
	HaySiguiente() bool
	// Siguiente actualiza el iterador para que apunte al siguiente elemento de la lista. Si ya no hay elementos entra en pánico con el mensaje "El iterador termino de iterar"
	Siguiente()
	// Insertar agrega un nuevo elemento a la lista entre el apuntando por el iterador y el anterior. Luego, el elemento apuntado sera el agregado.
	Insertar(T)
	// Borrar elimina el elemento apuntado por el iterador y luego apunta hacia el proximo. Si la lista esta vacía, entra en pánico con un mensaje "La lista esta vacia".
	Borrar() T
}
