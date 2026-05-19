package utils

// Importa os pacotes necessários

import (
	"fmt" // Usado para formatar mensagens
	"log" // Usado para registrar erros
)

// DeletePaciente remove um paciente do banco de dados pelo CPF informado
// Retorna um erro caso a exclusão falhe ou o paciente não seja encontrado
func DeletePaciente(cpf string) error {

	// Monta o comando SQL para deletar o paciente com o CPF informado
	query := `DELETE FROM pacientes WHERE cpf = $1`

	// Executa o comando SQL passando o CPF como argumento
	resultado, err := DB.Exec(query, cpf)
	if err != nil {
		// Registra o erro no terminal e retorna para o handler
		log.Println("Erro ao deletar paciente: ", err)
		return fmt.Errorf("erro ao deletar paciente: %w", err)
	}

	// Verifica quantas linhas foram afetadas pelo DELETE
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar exclusão: %w", err)
	}

	// Se nenhuma linha foi afetada, o CPF não existia no banco
	if linhasAfetadas == 0 {
		return fmt.Errorf("paciente com CPF %s não encontrado", cpf)
	}

	// Informa no terminal que o paciente foi removido com sucesso
	fmt.Println("Paciente deletado com sucesso!")
	return nil
}
