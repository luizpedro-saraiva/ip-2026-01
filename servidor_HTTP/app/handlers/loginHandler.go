package handlers

// Importa os pacotes necessários

import (
	"fmt"                     // Usado para escrever respostas formatadas
	"net/http"                // Usado para lidar com requisições e respostas HTTP
	"servidorSaude/app/utils" // Importa as funções utilitárias do projeto
)

// LoginHandler processa o formulário de busca de paciente pelo CPF
// Corresponde à rota "/login" definida no main.go
// No contexto de saúde, funciona como a consulta/busca do paciente
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Verifica se o método da requisição é POST (envio de formulário)
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Faz o parse dos dados enviados pelo formulário HTML
	err := r.ParseForm()
	if err != nil {
		// Retorna erro 400 se os dados do formulário forem inválidos
		http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
		return
	}

	// Lê o CPF enviado pelo formulário
	cpf := r.FormValue("cpf")

	// Busca o paciente no banco de dados pelo CPF
	paciente := utils.GetPacienteByCPF(cpf)

	// Verifica se o paciente foi encontrado
	if paciente == nil {
		fmt.Fprintf(w, "Paciente com CPF %s não encontrado.", cpf)
		return
	}

	// Exibe os dados do paciente encontrado no navegador
	fmt.Fprintf(w, `
		<h2>Dados do Paciente</h2>
		<p><b>Nome:</b> %s</p>
		<p><b>CPF:</b> %s</p>
		<p><b>Data de Nascimento:</b> %s</p>
		<p><b>Telefone:</b> %s</p>
		<p><b>E-mail:</b> %s</p>
		<p><b>Diagnóstico:</b> %s</p>
		<br><a href="/">Voltar ao início</a>
	`,
		paciente.NomeCompleto,
		paciente.CPF,
		paciente.DataNascimento,
		paciente.Telefone,
		paciente.Email,
		paciente.Diagnostico,
	)
}
