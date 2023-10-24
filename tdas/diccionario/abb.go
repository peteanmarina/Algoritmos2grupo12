package diccionario

func (abb *abb[K, V]) Diccionario() {}

type abb[K comparable, V any] struct {
	raiz     *arbol[K, V]
	cmp      func(K, K) int
	cantidad int
}

type arbol[K comparable, T any] struct {
	clave    K
	dato     T
	hijo_izq *arbol[K, T]
	hijo_der *arbol[K, T]
}

type iteradorDiccionarioOrdenado[K comparable, V any] struct {
	abb *arbol[K, V]
	//...
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, funcion_cmp, 0}
}

func crearArbol[K comparable, V any](clave K, dato V) *arbol[K, V] {
	return &arbol[K, V]{clave, dato, nil, nil}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	panic("a")
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	panic("a")
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	panic("a")
}

func (abb *abb[K, V]) Obtener(clave K) V {
	panic("a")
}

func (abb *abb[K, V]) Borrar(clave K) V {
	panic("a")
}

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	panic("a")
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	panic("a")
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	panic("a")
}

func (i *iteradorDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	panic("a")
}

func (i *iteradorDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	panic("a")
}

func (i *iteradorDiccionarioOrdenado[K, V]) Siguiente() {
	panic("a")
}
