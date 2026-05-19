package handlers

// Importa os pacotes necessários

import (
	"fmt" // Usado para escrever respostas e formatar texto

	"net/http" // Usado para lidar com requisições e respostas HTTP

	"servidorSaude/app/utils" // Importa as funções utilitárias do projeto
)

// FormHandler processa o formulário de cadastro de um novo paciente

// Corresponde à rota "/form" definida no main.go

func FormHandler(w http.ResponseWriter, r *http.Request) {

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

	// Lê cada campo enviado pelo formulário

	nome := r.FormValue("nome_completo")

	cpf := r.FormValue("cpf")

	dataNascimento := r.FormValue("data_nascimento")

	telefone := r.FormValue("telefone")

	email := r.FormValue("email")

	diagnostico := r.FormValue("diagnostico")

	// Chama a função utilitária para salvar o paciente no banco de dados

	// e verifica se ocorreu algum erro

	if err := utils.CreatePaciente(nome, cpf, dataNascimento, telefone, email, diagnostico); err != nil {

		// Retorna erro 500 e informa o usuário caso o cadastro falhe

		http.Error(w, "Erro ao cadastrar paciente. Tente novamente.", http.StatusInternalServerError)

		return

	}

	// Exibe uma mensagem de sucesso no navegador

	fmt.Fprintf(w, `

        <h2>✅ Paciente cadastrado com sucesso!</h2>

        <p><b>Nome:</b> %s</p>

        <p><b>CPF:</b> %s</p>

        <br><a href="/">← Voltar ao início</a>

    `, nome, cpf)

}
