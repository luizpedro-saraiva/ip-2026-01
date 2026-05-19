package utils

// Importa os pacotes necessários

import (
	"log" // Usado para registrar erros no terminal
)

// Paciente representa a estrutura de dados de um paciente no sistema
// Cada campo corresponde a uma coluna da tabela "pacientes" no banco de dados
type Paciente struct {
	ID             int
	NomeCompleto   string
	CPF            string
	DataNascimento string
	Telefone       string
	Email          string
	Diagnostico    string
}

// GetPacienteByCPF busca um paciente no banco de dados pelo CPF informado
// Retorna um ponteiro para Paciente se encontrado, ou nil se não existir
func GetPacienteByCPF(cpf string) *Paciente {

	// Monta o comando SQL para buscar o paciente pelo CPF
	query := `SELECT id, nome_completo, cpf, data_nascimento, telefone, email, diagnostico
	          FROM pacientes WHERE cpf = $1`

	// Executa a busca e armazena o resultado na variável row
	row := DB.QueryRow(query, cpf)

	// Cria uma variável do tipo Paciente para guardar os dados encontrados
	var p Paciente

	// Tenta preencher a variável com os dados retornados do banco
	err := row.Scan(
		&p.ID,
		&p.NomeCompleto,
		&p.CPF,
		&p.DataNascimento,
		&p.Telefone,
		&p.Email,
		&p.Diagnostico,
	)

	if err != nil {
		// Registra no terminal se o paciente não foi encontrado ou ocorreu erro
		log.Println("Paciente não encontrado: ", err)
		return nil
	}

	// Retorna o ponteiro para o paciente encontrado
	return &p
}
