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
	pertenece, _, _ := abb.buscarNodo(abb.raiz, nil, clave, false)
	return pertenece != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _, _ := abb.buscarNodo(abb.raiz, nil, clave, false)
	if nodo == nil {
		panic(PANIC_NO_ENCONTRADO)
	}
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo, padre, _ := abb.buscarNodo(abb.raiz, nil, clave, false)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
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
	nuevo_nodo := crearNodoAbb[K, V](clave, valor)
	if abb.raiz == nil {
		abb.raiz = nuevo_nodo
		abb.cantidad++
		return
	}
	nodo, padre, izq := abb.buscarNodo(abb.raiz, nil, clave, false)
	if nodo == nil {
		abb.cantidad++
	}

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

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, f func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}

	mayor_nodo, _ := abb.raiz.buscarNodoMayor(nil)
	menor_nodo, _ := abb.raiz.buscarNodoMenor(nil)

	var desde_local, hasta_local K
	if desde == nil {
		desde_local = menor_nodo.clave
	} else {
		desde_local = *desde
	}
	if hasta == nil {
		hasta_local = mayor_nodo.clave
	} else {
		hasta_local = *hasta
	}

	if abb.cmp(desde_local, hasta_local) > 0 {
		return
	}

	abb.recorrerArbolAplicandoFuncion(&desde_local, &hasta_local, f, abb.raiz)
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}
	mas_chico, _ := abb.raiz.buscarNodoMenor(nil)
	mas_grande, _ := abb.raiz.buscarNodoMayor(nil)
	abb.recorrerArbolAplicandoFuncion(&mas_chico.clave, &mas_grande.clave, f, abb.raiz)
}

func (abb *abb[K, V]) recorrerArbolAplicandoFuncion(desde *K, hasta *K, f func(clave K, dato V) bool, nodo *nodoAbb[K, V]) bool {
	if nodo == nil {
		return true //no importa
	}

	clave_mayor_desde := abb.cmp(*desde, nodo.clave) <= 0
	clave_menor_hasta := abb.cmp(nodo.clave, *hasta) <= 0
	var sigo bool
	if clave_mayor_desde {
		sigo = abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_izq)
		if clave_menor_hasta && sigo {
			sigo = f(nodo.clave, nodo.dato)
			if sigo {
				abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_der)
			} else {
				return sigo // no importa
			}
		} else {
			return sigo // no importa
		}
		return sigo
	} else {
		abb.recorrerArbolAplicandoFuncion(desde, hasta, f, nodo.hijo_der)
		return true
	}
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

	menor_nodo, _ := abb.raiz.buscarNodoMenor(nil)
	mayor_nodo, _ := abb.raiz.buscarNodoMayor(nil)

	var desde_local, hasta_local K
	if desde == nil {
		desde_local = menor_nodo.clave
	} else {
		desde_local = *desde
	}
	if hasta == nil {
		hasta_local = mayor_nodo.clave
	} else {
		hasta_local = *hasta
	}
	no_en_rango := ((abb.cmp(desde_local, menor_nodo.clave) < 0) && (abb.cmp(menor_nodo.clave, hasta_local) > 0)) || ((abb.cmp(desde_local, mayor_nodo.clave) > 0) && (abb.cmp(mayor_nodo.clave, hasta_local) > 0))

	if abb.cmp(desde_local, hasta_local) > 0 || no_en_rango {
		return &iteradorDiccionarioOrdenado[K, V]{*abb, pila, &desde_local, &hasta_local}
	}

	abb.raiz.apilarHastaMenorDesde(pila, &desde_local, &hasta_local, abb.cmp)

	return &iteradorDiccionarioOrdenado[K, V]{*abb, pila, &desde_local, &hasta_local}
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
		if nodo_desapilado.hijo_der.hijo_izq.dentro_rango(i.desde, i.hasta, i.abb.cmp) {
			nodo_desapilado.hijo_der.hijo_izq.apilarHijosIzq(i.pila, i.desde, i.hasta, i.abb.cmp)
		}
	} else if nodo_desapilado.hijo_der != nil {
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

	comparacion_desde := cmp(*desde, nodo.clave) <= 0
	comparacion_hasta := cmp(*hasta, nodo.clave) >= 0

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

	if nodo.dentro_rango(desde, hasta, cmp) {
		pila.Apilar(nodo)
	}

	nodo.hijo_izq.apilarHijosIzq(pila, desde, hasta, cmp)
}

func (nodo *nodoAbb[K, V]) dentro_rango(desde *K, hasta *K, cmp func(K, K) int) bool {
	if nodo == nil {
		return false
	}
	return cmp(nodo.clave, *desde) >= 0 && cmp(nodo.clave, *hasta) <= 0
}