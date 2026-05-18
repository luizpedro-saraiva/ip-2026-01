package handlers

// Importa os pacotes necessários

import (
	"fmt"      // Usado para escrever a resposta no terminal
	"net/http" // Usado para lidar com requisições e respostas HTTP
)

// HelloHandler responde à rota "/hello" com uma mensagem simples
// É usado para testar se o servidor está funcionando corretamente
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	// Verifica se a URL acessada é exatamente "/hello"
	if r.URL.Path != "/hello" {
		// Retorna erro 404 caso a rota não exista
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Verifica se o método da requisição é GET
	if r.Method != "GET" {
		// Retorna erro 405 caso o método não seja GET
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Escreve a resposta no terminal (log) e na tela do navegador
	fmt.Fprintf(w, "Olá! O servidor de saúde está funcionando!")
}
