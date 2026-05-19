package utils

// Importa os pacotes necessários para conectar ao banco de dados

import (
	"bufio"  // Usado para ler o arquivo .env linha a linha
	"database/sql" // Pacote padrão do Go para trabalhar com banco de dados
	"fmt"          // Usado para formatar strings
	"log"          // Usado para registrar erros no terminal
	"os"           // Usado para ler variáveis de ambiente e arquivos
	"strings"      // Usado para manipular strings ao processar o .env

	_ "github.com/lib/pq" // Driver do PostgreSQL para o Go (importado apenas pelo efeito colateral)
)

// DB é a variável global que armazena a conexão com o banco de dados
// Outros arquivos do pacote utils podem usá-la diretamente
var DB *sql.DB

// carregarEnv lê o arquivo .env e define as variáveis de ambiente
// que ainda não estejam definidas no sistema operacional
func carregarEnv() {
	arquivo, err := os.Open(".env")
	if err != nil {
		// Se o .env não existir, assume que as variáveis já estão no ambiente
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())

		// Ignora linhas vazias e comentários
		if linha == "" || strings.HasPrefix(linha, "#") {
			continue
		}

		// Divide a linha em chave e valor pelo primeiro "="
		partes := strings.SplitN(linha, "=", 2)
		if len(partes) != 2 {
			continue
		}

		chave := strings.TrimSpace(partes[0])
		valor := strings.TrimSpace(partes[1])

		// Define a variável apenas se ela ainda não estiver definida no sistema
		if os.Getenv(chave) == "" {
			os.Setenv(chave, valor)
		}
	}
}

// ConnectToDB abre a conexão com o banco de dados PostgreSQL
// Usa as variáveis de ambiente definidas no arquivo .env
func ConnectToDB() {

	// Carrega o arquivo .env antes de ler as variáveis
	carregarEnv()

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
