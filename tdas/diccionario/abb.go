package diccionario

import (
	TDAPila "tdas/pila"
)

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
	abb   abb[K, V] //PARA TENER CMP EN INTERAR X RANGOS
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, funcion_cmp, 0}
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave, dato, nil, nil}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	pertenece := abb.buscarNodo(&abb.raiz, clave)
	return *pertenece != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb.buscarNodo(&abb.raiz, clave)
	if *nodo == nil {
		panic(PANIC_NO_ENCONTRADO)
	}
	return (*nodo).dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	vinculo := abb.buscarNodo(&abb.raiz, clave)
	if *vinculo == nil {
		panic(PANIC_NO_ENCONTRADO)
	}
	valor := (*vinculo).dato

	if (*vinculo).hijo_izq != nil && (*vinculo).hijo_der != nil {
		reemplazo, _ := (*vinculo).hijo_izq.buscarNodoMayor(*vinculo)
		clave_nueva := reemplazo.clave
		dato_nuevo := abb.Borrar(reemplazo.clave)
		(*vinculo).clave = clave_nueva
		(*vinculo).dato = dato_nuevo
	} else {
		if (*vinculo).hijo_der == nil {
			*vinculo = (*vinculo).hijo_izq
		} else {
			*vinculo = (*vinculo).hijo_der
		}
		abb.cantidad--
	}

	return valor
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
	vinculo := abb.buscarNodo(&abb.raiz, clave)
	if *vinculo == nil {
		abb.cantidad++
	} else {
		(*vinculo).dato = valor
		return
	}

	nuevo_nodo := crearNodoAbb[K, V](clave, valor)
	*vinculo = nuevo_nodo
}

func (abb *abb[K, V]) buscarNodo(vinculo **nodoAbb[K, V], clave K) **nodoAbb[K, V] {
	if *vinculo == nil || (*vinculo).clave == clave {
		return vinculo
	}
	if abb.cmp((*vinculo).clave, clave) > 0 {
		return abb.buscarNodo(&(*vinculo).hijo_izq, clave)
	}
	return abb.buscarNodo(&(*vinculo).hijo_der, clave)
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

//////////////////////////// ITERADOR INTERNO ////////////////////////////

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, f func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}
	abb.recorrerArbolAplicandoFuncion(desde, hasta, f, abb.raiz)
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}
	abb.recorrerArbolAplicandoFuncion(nil, nil, f, abb.raiz)
}

func (abb *abb[K, V]) recorrerArbolAplicandoFuncion(desde *K, hasta *K, f func(clave K, dato V) bool, nodo *nodoAbb[K, V]) bool {
	if nodo == nil {
		return true // no importa
	}

	clave_mayor_desde := desde == nil || abb.cmp(*desde, nodo.clave) <= 0
	clave_menor_hasta := hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0

	if !clave_mayor_desde {
		abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_der)
		return true
	}

	sigo := abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_izq)

	if !clave_menor_hasta || !sigo {
		return sigo // no importa
	}

	sigo = f(nodo.clave, nodo.dato)

	if sigo {
		abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_der)
	}

	return sigo
}

//////////////////////////// ITERADOR EXTERNO ////////////////////////////

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if abb.raiz == nil {
		return &iteradorDiccionarioOrdenado[K, V]{pila: pila}
	}

	abb.raiz.apilarHastaMenorDesde(pila, desde, hasta, abb.cmp)

	return &iteradorDiccionarioOrdenado[K, V]{*abb, pila, desde, hasta}
}

func (i *iteradorDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	return !i.pila.EstaVacia()
}

func (i *iteradorDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	i.lanzarPanicTerminoIterar()
	return i.pila.VerTope().clave, i.pila.VerTope().dato
}

func (i *iteradorDiccionarioOrdenado[K, V]) Siguiente() {
	i.lanzarPanicTerminoIterar()
	nodo_desapilado := i.pila.Desapilar()

	if nodo_desapilado.hijo_der.dentro_rango(i.desde, i.hasta, i.abb.cmp) {
		i.pila.Apilar(nodo_desapilado.hijo_der)

	}
	if nodo_desapilado.hijo_der != nil {
		nodo_desapilado.hijo_der.hijo_izq.apilarHijosIzq(i.pila, i.desde, i.hasta, i.abb.cmp)
	}
}

func (i *iteradorDiccionarioOrdenado[K, V]) lanzarPanicTerminoIterar() {
	if i.HaySiguiente() {
		return
	}
	panic(PANIC_TERMINO_ITERAR)
}

func (nodo *nodoAbb[K, V]) apilarHastaMenorDesde(pila TDAPila.Pila[*nodoAbb[K, V]], desde *K, hasta *K, cmp func(K, K) int) {
	if nodo == nil {
		return
	}

	comparacion_desde := desde == nil || cmp(*desde, nodo.clave) <= 0
	comparacion_hasta := hasta == nil || cmp(*hasta, nodo.clave) >= 0

	if comparacion_desde {
		if comparacion_hasta {
			pila.Apilar(nodo)
		}
		nodo.hijo_izq.apilarHastaMenorDesde(pila, desde, hasta, cmp)
	} else {
		nodo.hijo_der.apilarHastaMenorDesde(pila, desde, hasta, cmp)
	}
}

func (nodo *nodoAbb[K, V]) apilarHijosIzq(pila TDAPila.Pila[*nodoAbb[K, V]], desde *K, hasta *K, cmp func(K, K) int) {
	if nodo == nil {
		return
	}
	pila.Apilar(nodo)
	nodo.hijo_izq.apilarHijosIzq(pila, desde, hasta, cmp)
}

func (nodo *nodoAbb[K, V]) dentro_rango(desde *K, hasta *K, cmp func(K, K) int) bool {
	if nodo == nil {
		return false
	}
	return cmp(nodo.clave, *desde) >= 0 && cmp(nodo.clave, *hasta) <= 0

}
