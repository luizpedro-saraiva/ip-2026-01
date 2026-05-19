package utils

// Importa os pacotes necessários

import (
	"fmt" // Usado para formatar mensagens
	"log" // Usado para registrar erros
)

// CreatePaciente insere um novo paciente na tabela "pacientes" do banco de dados
// Recebe todos os dados do formulário de cadastro como parâmetros
// Retorna um erro caso a inserção falhe, ou nil se for bem-sucedida
func CreatePaciente(nome string, cpf string, dataNascimento string, telefone string, email string, diagnostico string) error {

	// Monta o comando SQL para inserir o paciente no banco
	// Os $1, $2, $3... são os lugares onde os valores serão colocados com segurança
	query := `INSERT INTO pacientes (nome_completo, cpf, data_nascimento, telefone, email, diagnostico)
	          VALUES ($1, $2, $3, $4, $5, $6)`

	// Executa o comando SQL passando os valores como argumentos
	_, err := DB.Exec(query, nome, cpf, dataNascimento, telefone, email, diagnostico)
	if err != nil {
		// Registra o erro no terminal e retorna o erro para o handler
		log.Println("Erro ao cadastrar paciente: ", err)
		return fmt.Errorf("erro ao cadastrar paciente: %w", err)
	}

	// Informa no terminal que o paciente foi cadastrado com sucesso
	fmt.Println("Paciente cadastrado com sucesso!")
	return nil

}
