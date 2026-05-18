# Servidor HTTP de Saúde com GO

## Visão Geral

Este projeto é um servidor HTTP desenvolvido em GoLang no contexto de **saúde**, inspirado no projeto modelo [servidorHTTP](https://github.com/IP-2025-1/servidorHTTP).

Permite o **cadastro, consulta, atualização e exclusão de pacientes** (CRUD completo).
Utiliza **PostgreSQL** como banco de dados e fornece uma interface web simples em HTML/CSS.

---

## Estrutura do Projeto

```
servidorSaude/
├── .env.example
├── .gitignore
├── docker-compose.yml.example
├── go.mod
├── app/
│   ├── main.go
│   ├── handlers/
│   │   ├── helloHandler.go
│   │   ├── formHandler.go
│   │   ├── loginHandler.go
│   │   ├── updateAccountHandler.go
│   │   └── deleteAccountHandler.go
│   └── utils/
│       ├── connectToDB.go
│       ├── createPaciente.go
│       ├── getPacienteByCPF.go
│       ├── updatePaciente.go
│       └── deletePaciente.go
└── static/
    ├── index.html
    ├── forms/
    │   ├── createAccount.html
    │   ├── login.html
    │   ├── updateAccount.html
    │   └── deleteAccount.html
    └── styles/
        ├── index.style.css
        ├── createAccount.style.css
        ├── login.style.css
        ├── updateAccount.style.css
        └── deleteAccount.style.css
```

---

## Configuração do Ambiente

### Pré-requisitos

1. **GoLang** instalado na máquina
2. **PostgreSQL** instalado ou rodando via Docker
3. **Docker** (opcional)

### Passos para Configuração

1. Clone o repositório:
```
git clone <URL_DO_REPOSITORIO>
cd servidorSaude
```

2. Crie o arquivo `.env` a partir do exemplo:
```
cp .env.example .env
```
Edite o `.env` com suas credenciais do banco.

3. Se usar Docker, suba o banco:
```
cp docker-compose.yml.example docker-compose.yml
docker compose up -d
```

4. Crie a tabela no banco de dados:
```sql
CREATE TABLE pacientes (
    id              SERIAL PRIMARY KEY,
    nome_completo   VARCHAR(150) NOT NULL,
    cpf             VARCHAR(14)  NOT NULL UNIQUE,
    data_nascimento DATE         NOT NULL,
    telefone        VARCHAR(20),
    email           VARCHAR(150),
    diagnostico     TEXT,
    criado_em       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

5. Instale as dependências:
```
go mod tidy
```

---

## Executando o Projeto

```
go run app/main.go
```

O servidor estará disponível no endereço exibido no terminal, exemplo:
```
Servidor rodando em: http://192.168.1.10:3000/
```

---

## Rotas Principais

| Rota              | Descrição                        |
|-------------------|----------------------------------|
| `/`               | Página inicial                   |
| `/hello`          | Testa se o servidor está rodando |
| `/form`           | Cadastra um novo paciente        |
| `/login`          | Busca paciente pelo CPF          |
| `/updateAccount`  | Atualiza dados do paciente       |
| `/deleteAccount`  | Remove o paciente pelo CPF       |

### Formulários (páginas estáticas)

| Página                         | Descrição                  |
|--------------------------------|----------------------------|
| `/forms/createAccount.html`    | Formulário de cadastro     |
| `/forms/login.html`            | Formulário de busca        |
| `/forms/updateAccount.html`    | Formulário de atualização  |
| `/forms/deleteAccount.html`    | Formulário de exclusão     |

---

## Tecnologias Utilizadas

- **GoLang** — servidor HTTP nativo (`net/http`)
- **PostgreSQL** — banco de dados relacional
- **github.com/lib/pq** — driver PostgreSQL para Go
- **HTML5 + CSS3** — interface web estática
