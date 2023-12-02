package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var VALORES_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

type CompararSTRING func(int, int) int

func compararInt(s1, s2 int) int {
	if s1 < s2 {
		return -1
	}
	if s1 > s2 {
		return 1
	}
	return 0
}

func compararString(s1, s2 string) int {
	return strings.Compare(s1, s2)
}

func TestABBIteradorExternoRecorreInOrder(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	dic.Guardar(14, "A")
	dic.Guardar(4, "B")
	dic.Guardar(22, "C")
	dic.Guardar(3, "D")
	dic.Guardar(9, "E")
	dic.Guardar(16, "F")
	dic.Guardar(24, "G")
	iter := dic.Iterador()

	clave, valor := iter.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, "D", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 4, clave)
	require.EqualValues(t, "B", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 9, clave)
	require.EqualValues(t, "E", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 14, clave)
	require.EqualValues(t, "A", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 16, clave)
	require.EqualValues(t, "F", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 22, clave)
	require.EqualValues(t, "C", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 24, clave)
	require.EqualValues(t, "G", valor)
}

func TestIterarABBFueraDeRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	desde := 1
	hasta := 2
	dic.Guardar(14, "A")
	dic.Guardar(4, "B")
	dic.Guardar(22, "C")
	dic.Guardar(3, "D")
	dic.Guardar(9, "E")
	dic.Guardar(16, "F")
	dic.Guardar(24, "G")
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

}

func TestABBIteradorRangoRecorreInOrder(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	desde := 4
	hasta := 16
	dic.Guardar(14, "A")
	dic.Guardar(desde, "B")
	dic.Guardar(22, "C")
	dic.Guardar(3, "D")
	dic.Guardar(9, "E")
	dic.Guardar(hasta, "F")
	dic.Guardar(24, "G")
	iter := dic.IteradorRango(&desde, &hasta)

	clave, valor := iter.VerActual()
	require.EqualValues(t, desde, clave)
	require.EqualValues(t, "B", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 9, clave)
	require.EqualValues(t, "E", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 14, clave)
	require.EqualValues(t, "A", valor)

	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, hasta, clave)
	require.EqualValues(t, "F", valor)
}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](compararString)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](func(s1, s2 string) int {
		return 0
	})

	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](compararInt)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene unicamente su clave")
	dic := TDADiccionario.CrearABB[string, int](compararString)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](compararString)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazarDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](compararString)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](compararString)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestUnaClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](compararString)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestConValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](compararString)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestGuardarBorrarRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces")

	dic := TDADiccionario.CrearABB[int, int](compararInt)
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func TestIteradorInternoClave(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](compararString)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.True(t, dic.Pertenece(clave1))
	require.True(t, dic.Pertenece(cs[0]))
	require.True(t, dic.Pertenece(clave2))
	require.True(t, dic.Pertenece(cs[1]))
	require.True(t, dic.Pertenece(clave3))
	require.True(t, dic.Pertenece(cs[2]))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestArbolIteradorInternoRangosFuncionamiento(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	desde := 4
	hasta := 16
	dic.Guardar(14, "A")
	dic.Guardar(desde, "B")
	dic.Guardar(22, "C")
	dic.Guardar(3, "D")
	dic.Guardar(9, "E")
	dic.Guardar(hasta, "F")
	dic.Guardar(24, "G")

	var multiplicacion int = 1
	var ptr_m *int = &multiplicacion
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		*ptr_m *= clave
		return true
	})

	require.EqualValues(t, 8064, multiplicacion)

}

func TestArbolIteradorInternoFueraRango(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	desde := 50
	hasta := 100
	dic.Guardar(14, "A")
	dic.Guardar(4, "B")
	dic.Guardar(22, "C")
	dic.Guardar(3, "D")
	dic.Guardar(9, "E")
	dic.Guardar(16, "F")
	dic.Guardar(24, "G")

	var multiplicacion int = 1
	var ptr_m *int = &multiplicacion
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		*ptr_m *= clave
		return true
	})

	require.EqualValues(t, 1, multiplicacion)
}

func TestArbolIteradorInternoFueraRangosCambiados(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	desde := 100
	hasta := 50
	dic.Guardar(14, "A")
	dic.Guardar(4, "B")
	dic.Guardar(22, "C")
	dic.Guardar(3, "D")
	dic.Guardar(9, "E")
	dic.Guardar(16, "F")
	dic.Guardar(24, "G")

	var multiplicacion int = 1
	var ptr_m *int = &multiplicacion
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		*ptr_m *= clave
		return true
	})

	require.EqualValues(t, 1, multiplicacion)

}

func TestArbolIteradorInternoSinDesde(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	var desde *int
	hasta := 7
	dic.Guardar(1, "A")
	dic.Guardar(2, "B")
	dic.Guardar(3, "C")
	dic.Guardar(4, "D")
	dic.Guardar(5, "E")
	dic.Guardar(6, "F")
	dic.Guardar(hasta, "G")

	var multiplicacion int = 1
	var ptr_m *int = &multiplicacion
	dic.IterarRango(desde, &hasta, func(clave int, valor string) bool {
		*ptr_m *= clave
		return true
	})

	require.EqualValues(t, 5040, multiplicacion)
}

func TestArbolIteradorInternoSinHasta(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	var hasta *int
	desde := 1
	dic.Guardar(desde, "A")
	dic.Guardar(2, "B")
	dic.Guardar(3, "C")
	dic.Guardar(4, "D")
	dic.Guardar(5, "E")
	dic.Guardar(6, "F")
	dic.Guardar(7, "G")

	var multiplicacion int = 1
	var ptr_m *int = &multiplicacion
	dic.IterarRango(&desde, hasta, func(clave int, valor string) bool {
		*ptr_m *= clave
		return true
	})

	require.EqualValues(t, 5040, multiplicacion)
}

func TestIteradorInternoRangoValor(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente con el iterador interno con rango con y sin rango")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](compararString)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)
	factorial := 1
	ptrFactorial := &factorial
	ptrFactorial = &factorial
	dic.IterarRango(&clave4, &clave5, func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 840, factorial)

}

func TestIteradorInternoValor(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente con el iterador interno con y sin rango")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](compararString)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 5040, factorial)

}

func TestIterarRangoConCondiciónDeCorte(t *testing.T) {
	t.Log("Iteramos el diccionario hasta que se encuentre con el 7")
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	claves := []int{3, 2, 4, 0, 1, 7}
	valores := []string{"Elefante", "Gato", "Perro", "Hamster", "Camello", "Leon"}
	for j := 0; j < 6; j++ {
		dic.Guardar(claves[j], valores[j])
	}

	var zoo string
	var ptr_zoo *string = &zoo
	dic.IterarRango(&claves[3], &claves[2], func(clave int, dato string) bool {
		if clave == claves[1] {
			return false
		}
		*ptr_zoo += dato + " "
		return true
	})
	require.EqualValues(t, "Hamster Camello ", zoo)
}

func TestIteradorInternoValoresBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](compararString)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave4)
	dic.Borrar(clave2)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 630, factorial)

	factorial = 1
	ptrFactorial = &factorial
	rangoInicial := "A"
	dic.IterarRango(&rangoInicial, &clave2, func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})
	require.EqualValues(t, 210, factorial)
}

func ejecutarPruebaVol(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](compararString)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](compararString)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](compararString)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.True(t, dic.Pertenece(primero))

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.True(t, dic.Pertenece(segundo))
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.True(t, dic.Pertenece(tercero))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioOrdenadoIterarPorRango(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Pajaro"
	clave5 := "Oso"

	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	valor4 := "pio"
	valor5 := "rugido"

	claves := []string{clave1, clave2, clave3, clave4, clave5}
	valores := []string{valor1, valor2, valor3, valor4, valor5}

	dic := TDADiccionario.CrearABB[string, string](compararString)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	dic.Guardar(claves[3], valores[3])
	dic.Guardar(claves[4], valores[4])

	inicio := clave5
	fin := clave2

	iter := dic.IteradorRango(&inicio, &fin)

	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clave_menor_fin := compararString(fin, clave) >= 0
		clave_mayor_inicio := compararString(inicio, clave) <= 0
		require.True(t, clave_menor_fin)
		require.True(t, clave_mayor_inicio)
		iter.Siguiente()
	}
	require.PanicsWithValue(t, TDADiccionario.PANIC_TERMINO_ITERAR, func() { iter.VerActual() })
	require.PanicsWithValue(t, TDADiccionario.PANIC_TERMINO_ITERAR, func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](compararString)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.True(t, dic.Pertenece(primero))
	require.True(t, dic.Pertenece(segundo))
	require.True(t, dic.Pertenece(tercero))
}

func ejecutarPruebasVolIterador(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](compararString)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func TestVolIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](compararInt)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIteradorExternoRangosCruzados(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	desde := 22
	hasta := 3
	dic.Guardar(14, "A")
	dic.Guardar(4, "B")
	dic.Guardar(desde, "C")
	dic.Guardar(hasta, "D")
	dic.Guardar(9, "E")
	dic.Guardar(16, "F")
	dic.Guardar(24, "G")
	iter := dic.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoSinDesde(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	var desde *int
	hasta := 9
	dic.Guardar(14, "A")
	dic.Guardar(4, "B")
	dic.Guardar(3, "C")
	dic.Guardar(hasta, "D")
	dic.Guardar(9, "E")
	dic.Guardar(16, "F")
	dic.Guardar(24, "G")
	iter := dic.IteradorRango(desde, &hasta)
	clave, valor := iter.VerActual()
	require.EqualValues(t, 3, clave)
	require.EqualValues(t, "C", valor)
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 4, clave)
	require.EqualValues(t, "B", valor)
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 9, clave)
	require.EqualValues(t, "E", valor)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
}

func TestIteradorExternoSinHasta(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](compararInt)
	var hasta *int
	desde := 16
	dic.Guardar(14, "A")
	dic.Guardar(4, "B")
	dic.Guardar(3, "C")
	dic.Guardar(22, "D")
	dic.Guardar(9, "E")
	dic.Guardar(16, "F")
	dic.Guardar(24, "G")
	iter := dic.IteradorRango(&desde, hasta)
	clave, valor := iter.VerActual()
	require.EqualValues(t, 16, clave)
	require.EqualValues(t, "F", valor)
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 22, clave)
	require.EqualValues(t, "D", valor)
	iter.Siguiente()

	clave, valor = iter.VerActual()
	require.EqualValues(t, 24, clave)
	require.EqualValues(t, "G", valor)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
}
