package handlers

// Importa os pacotes necessários

import (
	"fmt"                     // Usado para escrever respostas formatadas
	"net/http"                // Usado para lidar com requisições e respostas HTTP
	"servidorSaude/app/utils" // Importa as funções utilitárias do projeto
)

// UpdateAccountHandler processa o formulário de atualização dos dados do paciente
// Corresponde à rota "/updateAccount" definida no main.go
func UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {

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

	// Lê os campos enviados pelo formulário de atualização
	cpf := r.FormValue("cpf")
	telefone := r.FormValue("telefone")
	email := r.FormValue("email")
	diagnostico := r.FormValue("diagnostico")

	// Chama a função utilitária para atualizar o paciente no banco de dados
	// e verifica se ocorreu algum erro
	if err := utils.UpdatePaciente(cpf, telefone, email, diagnostico); err != nil {
		// Retorna erro 500 e informa o usuário caso a atualização falhe
		http.Error(w, "Erro ao atualizar paciente. Tente novamente.", http.StatusInternalServerError)
		return
	}

	// Exibe mensagem de confirmação no navegador
	fmt.Fprintf(w, `
		<h2>✅ Dados atualizados com sucesso!</h2>
		<p><b>CPF:</b> %s</p>
		<br><a href="/">← Voltar ao início</a>
	`, cpf)
}
