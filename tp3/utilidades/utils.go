package utilidades

func EncontrarMaximo(m map[string]int) int {
	var maximo int

	for _, valor := range m {
		if valor > maximo {
			maximo = valor
		}
	}

	return maximo
}
