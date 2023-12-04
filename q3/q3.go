package q3

//O cozinheiro Remy preparou uma refeição para si mesmo e, enquanto almoçava, decidiu assistir a um vídeo no RataTube. No
//entanto, ele tem um tempo limitado de 𝑡 segundos para o almoço. Ele pediu a sua ajuda para escolher o vídeo ideal.
//
//O RataTube possui um feed de 𝑛 vídeos, cada um representado por uma estrutura de vídeo contendo informações sobre sua
//duração em segundos e o nível de entretenimento. O feed é inicialmente aberto no primeiro vídeo, e Remy pode pular para
//o próximo vídeo em 1 segundo (caso exista). Ele pode pular vídeos quantas vezes desejar, inclusive não pular nenhum.
//
//Sua tarefa é auxiliar Remy a escolher um vídeo que ele possa abrir e assistir dentro do tempo disponível, 𝑡 segundos. Se
//houver vários vídeos que se encaixam nessa condição, ele deseja escolher o vídeo com o maior nível de entretenimento.
//Retorne qualquer vídeo apropriado ou exiba um erro caso não haja um vídeo adequado dentro do tempo disponível.


import "fmt"

type Participant struct {
	Name string
	Role string
}

func calcularNumeroMaximoEquipes(participantes []Participant) int {
	contagemProgramadores := 0
	contagemMatematicos := 0

	for _, participante := range participantes {
		if participante.Role == "Programmer" {
			contagemProgramadores++
		} else if participante.Role == "Mathematician" {
			contagemMatematicos++
		}
	}

	numeroMaximoEquipes := 0
	for contagemProgramadores > 0 && contagemMatematicos > 0 {
		contagemProgramadores--
		contagemMatematicos--
		numeroMaximoEquipes++
	}

	return numeroMaximoEquipes
}
