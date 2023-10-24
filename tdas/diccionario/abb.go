package diccionario

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
	//comprobar bien los tipo de dato
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
	pertenece, _, _ := abb.reutilizable(abb.raiz, nil, clave, false)
	return pertenece != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	nodo, _, _ := abb.reutilizable(abb.raiz, nil, clave, false)
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	//obtenermos el nodo que deseamos borrar
	nodo, padre, _ := abb.reutilizable(abb.raiz, nil, clave, false)
	valor := nodo.dato
	abb.cantidad--

	if nodo.hijo_izq == nil && nodo.hijo_der == nil {
		// nodo sin hijos
		if padre == nil {
			abb.raiz = nil
		} else {
			padre.reemplazarHijo(nodo, nil)
		}
	} else if nodo.hijo_izq != nil && nodo.hijo_der != nil {
		// nodo con dos hijos
		reemplazo, padreReemplazo := nodo.hijo_izq.mas_der(nodo)
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
		// nodo con un solo hijo
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

func (nodo *nodoAbb[K, V]) mas_der(padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo.hijo_der == nil {
		return nodo, padre
	}
	return nodo.hijo_der.mas_der(nodo)
}

func (abb *abb[K, V]) Guardar(clave K, valor V) {
	//si sobreescribimos no habia que sumar en cantidad
	if !abb.Pertenece(clave) {
		abb.cantidad++
	}
	//si el diccionario esta vacia habia que crear el primer nodo
	if abb.raiz == nil {
		abb.raiz = crearNodoAbb[K, V](clave, valor)
		return
	}
	//obtenemos el nodo donde deberia de estar
	nodo, padre, izq := abb.reutilizable(abb.raiz, nil, clave, false)

	//por si se sobrescribe la raiz
	if padre == nil {
		nodo.dato = valor
		return
	}

	//caso generico
	if izq {
		padre.hijo_izq = crearNodoAbb[K, V](clave, valor)
	} else {
		padre.hijo_der = crearNodoAbb[K, V](clave, valor)
	}

}

func (abb *abb[K, V]) reutilizable(nodo *nodoAbb[K, V], padre *nodoAbb[K, V], clave K, izq bool) (*nodoAbb[K, V], *nodoAbb[K, V], bool) {
	if nodo == nil || abb.cmp(nodo.clave, clave) == 0 {
		return nodo, padre, izq
	}

	if abb.cmp(nodo.clave, clave) > 0 {
		return abb.reutilizable(nodo.hijo_izq, nodo, clave, true)
	}
	return abb.reutilizable(nodo.hijo_der, nodo, clave, false)
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (nodo *nodoAbb[K, V]) iterar(f func(clave K, dato V) bool) {
	//propuesta: modificar esta funcion para que sea un sub-caso de IterarRango, es decir, rearmarla
	//para que Iterar() solamente sea abb.raiz.iterar(f,inicio,fin) (es pseudocodigo)
	if nodo == nil {
		return
	}
	if f(nodo.clave, nodo.dato) {
		nodo.hijo_izq.iterar(f)
		nodo.hijo_der.iterar(f)
	}
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	abb.raiz.iterar(f)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	panic("a")
}

// Propuesta: es la misma que para Iterar.
func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorDiccionarioOrdenado[K, V]{abb.raiz, nil, nil}
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

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	panic("a")
}
