package main

func countingSortEdades(arreglo []Persona) {
	var minEdad int = arreglo[0].edad
	var maxEdad int = arreglo[0].edad
	for i := 1; i < len(arreglo); i++ {
		if arreglo[i].edad > maxEdad {
			maxEdad = arreglo[i].edad
		}
		if arreglo[i].edad < minEdad {
			minEdad = arreglo[i].edad
		}
	}
	var rango int = (maxEdad - minEdad) - 1
	var contador [rango]int
	for i := 0; i < rango; i++ {
		contador[obtenerIndiceEdad(arreglo[i].edad, minEdad)]
	}

}

func obtenerIndiceEdad(edad, min int) int {
	return edad - min
}

func countingSortNacionalidades(arreglo []Persona) {

}

func radixSort(arreglo []Persona) {
	countingSortNacionalidades(arreglo)
	countingSortEdades(arreglo)
}

type Persona struct {
	nombreCompleto string
	edad           int
	nacionalidad   string
}

func main() {

	var nombreArray [10]int
	nombreArray[0] = 1

}
