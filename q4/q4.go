package q4

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.


import "fmt"

func encontrarCidadeDestino(caminhos [][]string) (string, error) {
	cidadesEntrada := make(map[string]struct{})
	cidadesSaida := make(map[string]struct{})

	for _, caminho := range caminhos {
		cidadesEntrada[caminho[0]] = struct{}{}
		cidadesSaida[caminho[1]] = struct{}{}
	}

	for _, caminho := range caminhos {
		if _, ok := cidadesEntrada[caminho[1]]; !ok {
			return caminho[1], nil
		}
	}

	return "", fmt.Errorf("nenhuma cidade de destino encontrada")
}

func main() {
	caminhos := [][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}

	cidadeDestino, err := encontrarCidadeDestino(caminhos)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Cidade de Destino:", cidadeDestino)
}
