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
		reemplazo, _ := nodo.hijo_izq.buscarNodoMayor(nodo)
		clave_nueva := reemplazo.clave
		dato_nuevo := abb.Borrar(reemplazo.clave)
		abb.cantidad++ //porque se va a "borrar 2 veces" cuando en realidad es solo 1
		nodo.clave = clave_nueva
		nodo.dato = dato_nuevo

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

func (nodo *nodoAbb[K, V]) buscarNodoMenor(padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo.hijo_izq == nil {
		return nodo, padre
	}
	return nodo.hijo_izq.buscarNodoMenor(nodo)
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
	clave_mayor_desde := abb.cmp(*desde, nodo.clave) <= 0
	clave_menor_hasta := abb.cmp(nodo.clave, *hasta) <= 0
	if clave_mayor_desde {
		abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_izq)
		if clave_menor_hasta {
			condicion := f(nodo.clave, nodo.dato)
			if !condicion {
				return
			}
			abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_der)
		}
	}
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, f func(clave K, dato V) bool) {
	abb.recorrerArbolAplicandoFuncion(desde, hasta, f, abb.raiz)
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}
	mas_chico, _ := abb.raiz.buscarNodoMenor(nil)
	mas_grande, _ := abb.raiz.buscarNodoMayor(nil)
	abb.recorrerArbolAplicandoFuncion(&mas_chico.clave, &mas_grande.clave, f, abb.raiz)
}

//////////////////////////// ITERADOR EXTERNO ////////////////////////////

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if abb.raiz != nil {
		pila.Apilar(abb.raiz)
	}
	return &iteradorDiccionarioOrdenado[K, V]{abb.raiz, *abb, pila, nil, nil}
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if abb.cmp(*desde, *hasta) > 0 {
		hasta, desde = desde, hasta
	}
	primero := abb.buscarPrimeroEnRango(abb.raiz, *desde, *hasta)
	if primero != nil {
		pila.Apilar(primero)
	}

	return &iteradorDiccionarioOrdenado[K, V]{abb.raiz, *abb, pila, desde, hasta}
}

func (abb *abb[K, V]) buscarPrimeroEnRango(nodo *nodoAbb[K, V], desde K, hasta K) *nodoAbb[K, V] {
	if abb.cmp(nodo.clave, hasta) <= 0 && abb.cmp(desde, nodo.clave) <= 0 {
		return abb.raiz
	}
	if nodo.hijo_izq == nil && nodo.hijo_der == nil {
		return nil
	}
	if abb.cmp(abb.raiz.clave, hasta) > 0 {
		return abb.buscarPrimeroEnRango(abb.raiz.hijo_izq, desde, hasta)
	}
	if abb.cmp(abb.raiz.clave, desde) < 0 {
		return abb.buscarPrimeroEnRango(abb.raiz.hijo_der, desde, hasta)
	}
	return nil
}

func (i *iteradorDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	return !i.pila.EstaVacia()
}

func (i *iteradorDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	if i.pila.EstaVacia() {
		panic(PANIC_TERMINO_ITERAR)
	}
	return i.pila.VerTope().clave, i.nodo_actual.dato
}

func (i *iteradorDiccionarioOrdenado[K, V]) Siguiente() {
	i.lanzarPanicTerminoIterar()
	var izq *nodoAbb[K, V]
	var der *nodoAbb[K, V]
	i.nodo_actual = i.pila.Desapilar()
	if i.nodo_actual == nil {
		return
	}
	if i.nodo_actual.hijo_der != nil && (i.hasta == nil || i.abb.cmp(i.nodo_actual.hijo_der.clave, *i.hasta) <= 0) {
		der = i.nodo_actual.hijo_der
		i.pila.Apilar(der)
	}
	if i.nodo_actual.hijo_izq != nil && (i.desde == nil || i.abb.cmp(*i.desde, i.nodo_actual.hijo_izq.clave) <= 0) {
		izq = i.nodo_actual.hijo_izq
		i.pila.Apilar(izq)
	}
}

func (i *iteradorDiccionarioOrdenado[K, V]) lanzarPanicTerminoIterar() {
	if i.HaySiguiente() {
		return
	}
	panic(PANIC_TERMINO_ITERAR)
}
