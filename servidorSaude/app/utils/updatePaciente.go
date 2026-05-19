package utils

// Importa os pacotes necessários

import (
	"fmt" // Usado para formatar mensagens
	"log" // Usado para registrar erros
)

// UpdatePaciente atualiza os dados de um paciente no banco de dados
// Identifica o paciente pelo CPF e atualiza telefone, email e diagnostico
// Retorna um erro caso a atualização falhe, ou nil se for bem-sucedida
func UpdatePaciente(cpf string, telefone string, email string, diagnostico string) error {

	// Monta o comando SQL para atualizar os dados do paciente
	query := `UPDATE pacientes
	          SET telefone = $1, email = $2, diagnostico = $3
	          WHERE cpf = $4`

	// Executa o comando SQL com os novos valores
	_, err := DB.Exec(query, telefone, email, diagnostico, cpf)
	if err != nil {
		// Registra o erro no terminal e retorna para o handler
		log.Println("Erro ao atualizar paciente: ", err)
		return fmt.Errorf("erro ao atualizar paciente: %w", err)
	}

	// Informa no terminal que os dados foram atualizados com sucesso
	fmt.Println("Paciente atualizado com sucesso!")
	return nil
}
