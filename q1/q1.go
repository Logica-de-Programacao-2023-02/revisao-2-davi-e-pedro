package q1

//Na loja de animais à venda, existem alguns tipos de ração disponíveis para compra, sendo eles:
//
//* Ração para cachorro
//* Ração para gato
//* Ração universal
//
//O estoque dessas rações está representado em um mapa, onde a chave é o nome da ração e o valor correspondente é a
//quantidade disponível em estoque.
//
//Polycarp possui 𝑥 cães e 𝑦 gatos. Gostaríamos de determinar se é possível para ele comprar comida suficiente para todos
//os seus animais na loja. Cada um dos seus cães e gatos deve receber um pacote de ração adequado para sua espécie.


import "fmt"

func verificarPossibilidadeCompra(qtdCachorros, qtdGatos int, estoque map[string]int) bool {
	// Verifica se há rações suficientes para os cachorros
	estoqueCachorro, okCachorro := estoque["dog"]
	if !okCachorro || estoqueCachorro < qtdCachorros {
		return false
	}

	// Verifica se há rações suficientes para os gatos
	estoqueGato, okGato := estoque["cat"]
	if !okGato || estoqueGato < qtdGatos {
		return false
	}

	// Calcula a quantidade total necessária
	qtdTotalNecessaria := qtdCachorros + qtdGatos

	// Verifica se há rações universais suficientes para complementar
	estoqueUniversal, okUniversal := estoque["universal"]
	return okUniversal && estoqueUniversal >= qtdTotalNecessaria
}
