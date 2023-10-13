package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	entero1  = 10
	entero2  = 17
	entero3  = 22
	entero4  = 26
	entero5  = 7
	entero6  = 10
	cadena1  = "hola!"
	cadena2  = "como"
	cadena3  = "estas"
	cadena4  = "??"
	float1   = 0.1
	float2   = 0.2
	volumen1 = 5000
	volumen2 = 10000
)

func TestRecienCreada(t *testing.T) {
	t.Log("Verificamos que lista recien creada se comporte como vacia")
	lista := TDALista.CrearListaEnlazada[string]()
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, TDALista.Mensaje_lista_vacia, func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
}

func TestParElementos(t *testing.T) {
	t.Log("Insertamos dos elementos al principio y los borramos dejando lista vacia")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero(cadena1)
	lista.InsertarPrimero(cadena2)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, cadena2, lista.BorrarPrimero())
	require.Equal(t, cadena1, lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
}

func TestInsertarPrimeroYUltimo(t *testing.T) {
	t.Log("Verificamos que insertar pocos elementos al principio y final, y borrarlos se realiza correctamente")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(entero1)
	lista.InsertarUltimo(entero2)
	lista.InsertarPrimero(entero3)
	lista.InsertarUltimo(entero4)
	require.Equal(t, 4, lista.Largo())
	enteros := []int{entero3, entero1, entero2, entero4}
	for i := 0; i < len(enteros); i++ {
		require.Equal(t, enteros[i], lista.BorrarPrimero())
	}
}

func TestVolumen(t *testing.T) {
	t.Log("Insertamos gran cantidad de elementos al principio y final")
	lista := TDALista.CrearListaEnlazada[int64]()
	var i int64
	for i = volumen1; i >= 0; i-- {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
	}
	for i = volumen1 + 1; i < volumen2; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
	}

	t.Log("Borramos el primer elemento hasta dejar lista vacia")
	for i = 0; i < volumen2; i++ {
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.Equal(t, 0, lista.Largo())
}

/////////////////////////////ITERADOR INTERNO////////////////////////////////

func TestIteradorInternoSumarTodosLosElementos(t *testing.T) {
	t.Log("Test con funcion que suma todos los elementos de la lista")
	var (
		suma    int16  = 0
		puntero *int16 = &suma
		lista          = TDALista.CrearListaEnlazada[int16]()
	)

	elementos := []int{entero1, entero2, entero3, entero4, entero5, entero6}
	for _, elemento := range elementos {
		lista.InsertarPrimero(int16(elemento))
	}
	lista.Iterar(func(v int16) bool {
		*puntero += v
		return true
	})

	require.Equal(t, int16(entero1+entero2+entero3+entero4+entero5+entero6), *puntero)

	*puntero = 0
}

func TestIteradorInternoSumarParesFinSiEsSiete(t *testing.T) {
	t.Log("Test con funcion que itera la lista, sumando todos los numeros pares. Finaliza si el numero es 7")
	var (
		suma    int16  = 0
		puntero *int16 = &suma
		lista          = TDALista.CrearListaEnlazada[int16]()
	)
	elementos := []int{entero1, entero2, entero3, entero4, entero5, entero6}
	for _, elemento := range elementos {
		lista.InsertarPrimero(int16(elemento))
	}

	lista.Iterar(func(v int16) bool {
		if v == 7 {
			return false
		}
		if v%2 == 0 {
			*puntero += v
		}
		return true
	})
	require.Equal(t, int16(entero6), *puntero)
}

/////////////////////////////ITERADOR EXTERNO////////////////////////////////

func TestInsertarUnElementoIterador(t *testing.T) {

	t.Log("Verificamos el comportamiento general al insertar un solo elemento con el iterador")
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())
	iter.Insertar(float1)
	require.True(t, iter.HaySiguiente())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, float1, lista.VerPrimero())
	require.Equal(t, float1, lista.VerUltimo())
	require.Equal(t, float1, iter.VerActual())
	iter.Siguiente()
	require.PanicsWithValue(t, TDALista.Mensaje_iterador, func() { iter.Siguiente() })
	require.PanicsWithValue(t, TDALista.Mensaje_iterador, func() { iter.VerActual() })
}

func TestInsertarElementosEnMedioIterador(t *testing.T) {
	t.Log("En lista vacia insertamos elementos en el medio. Iteramos con otro iterador")
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(entero1)
	iterador.Insertar(entero2)
	iterador.Siguiente()
	elementos := []int{entero3, entero4, entero5}
	for _, elemento := range elementos {
		iterador.Insertar(elemento)
		require.Equal(t, elemento, iterador.VerActual())
	}
	elementos = []int{entero4, entero3, entero1}
	for i := 0; i < len(elementos); i++ {
		iterador.Siguiente()
	}
	iterador.Siguiente()
	iterador.Insertar(entero6)
	iterador.Siguiente()

	iterador2 := lista.Iterador()
	elementos = []int{entero2, entero5, entero4, entero3, entero1, entero6}
	for _, elemento := range elementos {
		require.Equal(t, elemento, iterador2.VerActual())
		iterador2.Siguiente()
	}

}

func TestInsertarEnListaVaciaIterador(t *testing.T) {
	t.Log("Verificamos que podamos insertar elementos cuando iteramos una lista vacia")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(entero1)
	require.Equal(t, entero1, iter.VerActual())
	iter.Insertar(entero2)
	require.Equal(t, entero2, iter.VerActual())
	iter.Insertar(entero3)
	require.Equal(t, entero3, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, entero2, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, entero1, iter.VerActual())
}

func TestBorrarPrimeroYUltimo(t *testing.T) {
	t.Log("Borramos estando el primer y ultimo elemento, verificando que se actualicen")
	lista := TDALista.CrearListaEnlazada[int]()

	iter := lista.Iterador()
	require.PanicsWithValue(t, TDALista.Mensaje_iterador, func() { iter.Borrar() })
	iter.Insertar(entero1)
	require.Equal(t, 1, lista.Largo())
	iter.Borrar()
	iter.Insertar(entero2)
	iter.Siguiente()
	iter.Insertar(entero3)
	require.Equal(t, entero3, lista.VerUltimo())
	require.Equal(t, entero2, lista.VerPrimero())

}
