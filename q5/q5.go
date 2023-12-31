package q5

//Um novo serviço de e-mail, chamado "CEUBdesk", será inaugurado no CEUB em um futuro próximo. A administração do
//site quer lançar o projeto o mais rápido possível, por isso eles pedem a sua ajuda. Você é sugerido(a) a implementar o
//protótipo do sistema de registro do site. O sistema deve funcionar com o seguinte princípio.
//
//Cada vez que um novo usuário deseja se registrar, ele envia ao sistema uma solicitação com o seu nome. Se esse nome
//ainda não existe no banco de dados do sistema, ele é inserido no banco de dados e o usuário recebe uma resposta "OK",
//confirmando o registro bem-sucedido. Se o nome já existe no banco de dados do sistema, o sistema cria um novo nome de
//usuário, envia-o ao usuário como sugestão e também insere a sugestão no banco de dados. O novo nome é formado pela
//seguinte regra. Números, começando com 1, são anexados um após o outro ao nome (name1, name2, ...), entre esses números,
//o menor `i` é encontrado de forma que namei ainda não exista no banco de dados.


import "fmt"

func registrarUsuarios(nomes []string) []string {
	bancoDeDados := make(map[string]int)
	resultado := make([]string, 0)

	for _, nome := range nomes {
		if ocorrencias, existe := bancoDeDados[nome]; !existe {
			bancoDeDados[nome] = 1
			resultado = append(resultado, "OK")
		} else {
			sugestao := fmt.Sprintf("%s%d", nome, ocorrencias)
			bancoDeDados[nome]++
			bancoDeDados[sugestao] = 1
			resultado = append(resultado, sugestao)
		}
	}

	return resultado
}
