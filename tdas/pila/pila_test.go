package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Creo una pila de string y verifico si se comporta como una pila vacia")
	pila2 := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila2.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila2.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila2.VerTope() })
}

func TestUnicoElemento(t *testing.T) {
	t.Log("Creo una pila de bool y verifico que apile y desapile un elemento correctamente.")
	pila3 := TDAPila.CrearPilaDinamica[bool]()
	require.True(t, pila3.EstaVacia())
	pila3.Apilar(true)
	require.False(t, pila3.EstaVacia())
	require.Equal(t, true, pila3.VerTope())
	require.Equal(t, true, pila3.Desapilar())
	require.True(t, pila3.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila3.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila3.VerTope() })

}

func TestApilarAlgunosElementos(t *testing.T) {
	t.Log("Creo una pila de int. Apilo y desapilo algunos elementos de dos maneras.")
	t.Log("1) Primero apilo todos y despues desapilo todos")
	t.Log("2) Por cada elemento que apilo, desapilo inmediatamente despues")
	t.Log("Verifico al final de ambas que se comporte como pila vacia")

	pila1 := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 5; i++ {
		pila1.Apilar(i)
	}
	for i := 4; i > -1; i-- {
		require.Equal(t, i, pila1.VerTope())
		require.Equal(t, i, pila1.Desapilar())
	}
	require.True(t, pila1.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila1.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila1.VerTope() })

	for i := 0; i < 5; i++ { //ahora intercalando apilar y desapilar
		pila1.Apilar(i)
		require.Equal(t, i, pila1.VerTope())
		require.Equal(t, i, pila1.Desapilar())
	}
	require.True(t, pila1.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila1.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila1.VerTope() })
}

func TestVolumen(t *testing.T) {
	t.Log("Creo una pila de int. Apilo y desapilo una gran cantidad de elementos de dos maneras.")
	t.Log("1) Primero apilo todos y despues desapilo todos")
	t.Log("2) Por cada elemento que apilo, desapilo inmediatamente despues")
	t.Log("Verifico al final de ambas que se comporte como pila vacia")

	pila1 := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10000; i++ {
		pila1.Apilar(i)
		require.False(t, pila1.EstaVacia())
	}
	for i := 9999; i > -1; i-- {
		require.Equal(t, i, pila1.VerTope())
		require.Equal(t, i, pila1.Desapilar())
	}
	require.True(t, pila1.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila1.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila1.VerTope() })

	for i := 0; i < 5000; i++ { //ahora intercalando apilar y desapilar
		pila1.Apilar(i)
		require.False(t, pila1.EstaVacia())
		require.Equal(t, i, pila1.VerTope())
		require.Equal(t, i, pila1.Desapilar())
		require.True(t, pila1.EstaVacia())
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila1.Desapilar() })
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila1.VerTope() })
	}
}
