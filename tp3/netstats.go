package main

import (
	"fmt"
	"netstats/grafo"
)

/* func main() {
	grafo := grafo.CrearGrafo()
	grafo.AgregarVertice("Hola")
	grafo.AgregarVertice("Soy")
	grafo.AgregarVertice("Nicor")
	grafo.AgregarVertice("Que tal")
	grafo.AgregarArista("Hola", "Nicor")
	grafo.AgregarArista("Soy", "Nicor")
	grafo.AgregarArista("Nicor", "Que tal")

	fmt.Println(grafo.ObtenerAdyacentes("Hola"))
	fmt.Println(grafo.ObtenerVertices())
	fmt.Println(grafo.ExisteArista("Hola", "Nicor"))
	fmt.Println(grafo.ExisteVertice("Soy"))

	grafo.SacarVertice("Soy")
	grafo.SacarArista("Hola", "Nicor")

	fmt.Println(grafo.ObtenerAdyacentes("Hola"))
	fmt.Println(grafo.ObtenerVertices())
	fmt.Println(grafo.ExisteArista("Hola", "Nicor"))
	fmt.Println(grafo.ExisteVertice("Soy"))
} */

func main() {
	grafoMarvel := grafo.CrearGrafo()

	// Agregar personajes como vértices
	grafoMarvel.AgregarVertice("Iron Man")
	grafoMarvel.AgregarVertice("Captain America")
	grafoMarvel.AgregarVertice("Thor")
	grafoMarvel.AgregarVertice("Hulk")
	grafoMarvel.AgregarVertice("Black Widow")
	grafoMarvel.AgregarVertice("Spider-Man")

	// Agregar relaciones entre personajes como aristas
	grafoMarvel.AgregarArista("Iron Man", "Captain America")
	grafoMarvel.AgregarArista("Captain America", "Thor")
	grafoMarvel.AgregarArista("Thor", "Hulk")
	grafoMarvel.AgregarArista("Hulk", "Black Widow")
	grafoMarvel.AgregarArista("Black Widow", "Iron Man")
	grafoMarvel.AgregarArista("Spider-Man", "Iron Man")

	// Consultar relaciones y vértices
	fmt.Println("Adyacentes a 'Iron Man':", grafoMarvel.ObtenerAdyacentes("Iron Man"))
	fmt.Println("Todos los personajes:", grafoMarvel.ObtenerVertices())
	fmt.Println("¿Existe relación entre 'Iron Man' y 'Captain America'?", grafoMarvel.ExisteArista("Iron Man", "Captain America"))
	fmt.Println("¿Existe el personaje 'Black Panther'?", grafoMarvel.ExisteVertice("Black Panther"))

	// Realizar algunas eliminaciones
	grafoMarvel.SacarVertice("Spider-Man")
	grafoMarvel.SacarArista("Thor", "Hulk")

	// Consultar después de eliminaciones
	fmt.Println("Adyacentes a 'Thor' después de eliminar la arista con 'Hulk':", grafoMarvel.ObtenerAdyacentes("Thor"))
	fmt.Println("Todos los personajes después de eliminar 'Spider-Man':", grafoMarvel.ObtenerVertices())
}
