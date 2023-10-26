package diccionario

import (
	TDAPila "tdas/pila"
)

func (abb *abb[K, V]) Diccionario() {}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cmp      func(K, K) int
	cantidad int
}

type nodoAbb[K comparable, T any] struct {
	clave    K
	dato     T
	hijo_izq *nodoAbb[K, T]
	hijo_der *nodoAbb[K, T]
}

type iteradorDiccionarioOrdenado[K comparable, V any] struct {
	nodo_actual *nodoAbb[K, V]
	abb         abb[K, V] //PARA TENER CMP EN INTERAR X RANGOS
	pila        TDAPila.Pila[*nodoAbb[K, V]]
	desde       *K
	hasta       *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, funcion_cmp, 0}
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave, dato, nil, nil}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	pertenece, _, _ := abb.buscarNodo(abb.raiz, nil, clave, false)
	return pertenece != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	nodo, _, _ := abb.buscarNodo(abb.raiz, nil, clave, false)
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	nodo, padre, _ := abb.buscarNodo(abb.raiz, nil, clave, false)
	valor := nodo.dato
	abb.cantidad--

	if nodo.hijo_izq == nil && nodo.hijo_der == nil {
		if padre == nil {
			abb.raiz = nil
		} else {
			padre.reemplazarHijo(nodo, nil)
		}
	} else if nodo.hijo_izq != nil && nodo.hijo_der != nil {
		reemplazo, padreReemplazo := nodo.hijo_izq.buscarNodoMayor(nodo)
		if nodo.hijo_izq.hijo_izq != nil {
			reemplazo.hijo_izq = nodo.hijo_izq.hijo_izq
		}
		reemplazo.hijo_der = nodo.hijo_der
		padreReemplazo.hijo_der = nil

		if padre == nil {
			abb.raiz = reemplazo
		} else {
			padre.reemplazarHijo(nodo, reemplazo)
		}

	} else {
		var hijo *nodoAbb[K, V]
		if nodo.hijo_izq != nil {
			hijo = nodo.hijo_izq
		} else {
			hijo = nodo.hijo_der
		}

		if padre == nil {
			abb.raiz = hijo
		} else {
			padre.reemplazarHijo(nodo, hijo)
		}
	}

	return valor
}

func (nodo *nodoAbb[K, V]) reemplazarHijo(viejo, nuevo *nodoAbb[K, V]) {
	if nodo.hijo_izq == viejo {
		nodo.hijo_izq = nuevo
	} else if nodo.hijo_der == viejo {
		nodo.hijo_der = nuevo
	}
}

func (nodo *nodoAbb[K, V]) buscarNodoMayor(padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo.hijo_der == nil {
		return nodo, padre
	}
	return nodo.hijo_der.buscarNodoMayor(nodo)
}

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	if !abb.Pertenece(clave) {
		abb.cantidad++
	}
	nuevo_nodo := crearNodoAbb[K, V](clave, valor)
	if abb.raiz == nil {
		abb.raiz = nuevo_nodo
		return
	}
	nodo, padre, izq := abb.buscarNodo(abb.raiz, nil, clave, false)

	if padre == nil {
		nodo.dato = valor
		return
	}

	if izq {
		padre.hijo_izq = nuevo_nodo
	} else {
		padre.hijo_der = nuevo_nodo
	}

}

func (abb *abb[K, V]) buscarNodo(nodo *nodoAbb[K, V], padre *nodoAbb[K, V], clave K, izq bool) (*nodoAbb[K, V], *nodoAbb[K, V], bool) {
	if nodo == nil || abb.cmp(nodo.clave, clave) == 0 {
		return nodo, padre, izq
	}

	if abb.cmp(nodo.clave, clave) > 0 {
		return abb.buscarNodo(nodo.hijo_izq, nodo, clave, true)
	}
	return abb.buscarNodo(nodo.hijo_der, nodo, clave, false)
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

//////////////////////////// ITERADOR INTERNO ////////////////////////////

func (abb *abb[K, V]) recorrerArbolAplicandoFuncion(desde *K, hasta *K, f func(clave K, dato V) bool, nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	if (desde == nil || abb.cmp(*desde, nodo.clave) <= 0) && (hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0) {
		if !f(nodo.clave, nodo.dato) {
			return
		}
	}
	if desde == nil || abb.cmp(*desde, nodo.clave) <= 0 {
		abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_izq)
	}
	if desde == nil || abb.cmp(nodo.clave, *hasta) <= 0 {
		abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_der)
	}
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, f func(clave K, dato V) bool) {
	abb.recorrerArbolAplicandoFuncion(desde, hasta, f, abb.raiz)
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	abb.recorrerArbolAplicandoFuncion(nil, nil, f, abb.raiz)
}

//////////////////////////// ITERADOR EXTERNO ////////////////////////////

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	abb.apilarNodosEnRango(&pila, abb.raiz, nil, nil)
	return &iteradorDiccionarioOrdenado[K, V]{abb.raiz, *abb, pila, nil, nil}
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if abb.cmp(*desde, *hasta) > 0 {
		hasta, desde = desde, hasta
	}
	abb.apilarNodosEnRango(&pila, abb.raiz, desde, hasta)

	return &iteradorDiccionarioOrdenado[K, V]{abb.raiz, *abb, pila, desde, hasta}
}

func (abb *abb[K, V]) apilarNodosEnRango(pila *TDAPila.Pila[*nodoAbb[K, V]], nodo *nodoAbb[K, V], desde *K, hasta *K) {
	if nodo == nil {
		return
	}
	if desde == nil || abb.cmp(*desde, nodo.clave) <= 0 {
		abb.apilarNodosEnRango(pila, nodo.hijo_izq, desde, hasta)
	}
	if (desde == nil || abb.cmp(*desde, nodo.clave) <= 0) && (hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0) {
		(*pila).Apilar(nodo)
	}
	if hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0 {
		abb.apilarNodosEnRango(pila, nodo.hijo_der, desde, hasta)
	}
}

func (i *iteradorDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	return !i.pila.EstaVacia()
}

func (i *iteradorDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	i.lanzarPanicTerminoIterar()
	return i.nodo_actual.clave, i.nodo_actual.dato
}

func (i *iteradorDiccionarioOrdenado[K, V]) Siguiente() {
	i.lanzarPanicTerminoIterar()
	if i.pila.EstaVacia() {
		i.nodo_actual = nil
		return
	}
	i.nodo_actual = i.pila.Desapilar()
}

func (i *iteradorDiccionarioOrdenado[K, V]) lanzarPanicTerminoIterar() {
	if !i.HaySiguiente() {
		panic(PANIC_TERMINO_ITERAR)
	}
}
