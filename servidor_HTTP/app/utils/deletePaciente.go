package utils

// Importa os pacotes necessários

import (
	"fmt" // Usado para formatar mensagens
	"log" // Usado para registrar erros
)

// DeletePaciente remove um paciente do banco de dados pelo CPF informado
func DeletePaciente(cpf string) {

	// Monta o comando SQL para deletar o paciente com o CPF informado
	query := `DELETE FROM pacientes WHERE cpf = $1`

	// Executa o comando SQL passando o CPF como argumento
	_, err := DB.Exec(query, cpf)
	if err != nil {
		// Registra o erro no terminal se a exclusão falhar
		log.Println("Erro ao deletar paciente: ", err)
		return
	}

	// Informa no terminal que o paciente foi removido com sucesso
	fmt.Println("Paciente deletado com sucesso!")
}
