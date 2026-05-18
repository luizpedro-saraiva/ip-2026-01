package utils

// Importa os pacotes necessários para conectar ao banco de dados

import (
	"database/sql" // Pacote padrão do Go para trabalhar com banco de dados
	"fmt"          // Usado para formatar strings
	"log"          // Usado para registrar erros no terminal
	"os"           // Usado para ler variáveis de ambiente

	_ "github.com/lib/pq" // Driver do PostgreSQL para o Go (importado apenas pelo efeito colateral)
)

// DB é a variável global que armazena a conexão com o banco de dados
// Outros arquivos do pacote utils podem usá-la diretamente
var DB *sql.DB

// ConnectToDB abre a conexão com o banco de dados PostgreSQL
// Usa as variáveis de ambiente definidas no arquivo .env
func ConnectToDB() {

	// Monta a string de conexão com os dados do banco
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Tenta abrir a conexão com o banco de dados
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		// Encerra o programa se não conseguir abrir a conexão
		log.Fatal("Erro ao abrir conexão com o banco de dados: ", err)
	}

	// Verifica se o banco está realmente acessível (ping)
	err = DB.Ping()
	if err != nil {
		// Encerra o programa se o banco não responder
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}

	// Informa no terminal que a conexão foi bem-sucedida
	fmt.Println("Conectado ao banco de dados com sucesso!")
}
