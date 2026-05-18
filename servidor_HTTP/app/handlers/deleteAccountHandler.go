package handlers

// Importa os pacotes necessários

import (
	"fmt"                     // Usado para escrever respostas formatadas
	"net/http"                // Usado para lidar com requisições e respostas HTTP
	"servidorSaude/app/utils" // Importa as funções utilitárias do projeto
)

// DeleteAccountHandler processa o formulário de exclusão de um paciente
// Corresponde à rota "/deleteAccount" definida no main.go
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {

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

	// Chama a função utilitária para deletar o paciente do banco de dados
	utils.DeletePaciente(cpf)

	// Exibe mensagem de confirmação no navegador
	fmt.Fprintf(w, "Paciente com CPF %s removido do sistema.", cpf)
}
