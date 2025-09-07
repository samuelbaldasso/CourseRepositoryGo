# Plataforma de Cursos

Este projeto é uma plataforma de cursos desenvolvida em GoLang. O objetivo é fornecer uma aplicação que permita a criação, leitura, atualização e exclusão de cursos.

## Estrutura do Projeto

A estrutura do projeto é organizada da seguinte forma:

```
plataforma-cursos
├── cmd
│   └── main.go                  # Ponto de entrada da aplicação
├── internal
│   ├── controllers
│   │   ├── course_controller.go   # Controlador para gerenciar cursos
│   │   ├── user_controller.go     # Controlador para gerenciar usuários
│   ├── models
│   │   ├── course.go              # Modelo que representa um curso
│   │   ├── user.go                # Modelo que representa um usuário
│   ├── routes
│   │   └── routes.go              # Configuração das rotas da aplicação
│   ├── services
│   │   ├── course_service.go      # Lógica de negócios dos cursos
│   │   ├── user_service.go        # Lógica de negócios dos usuários
│   ├── middleware
│   │   └── auth.go                # Middleware de autenticação
├── pkg
│   └── database
│       └── db.go                  # Funções para conexão com o banco de dados
├── go.mod                         # Módulo de configuração do Go
├── go.sum                         # Somas de verificação das dependências
├── Dockerfile                     # Dockerfile para containerização do serviço
└── README.md                      # Documentação do projeto
```

## Funcionalidades

- **Gerenciamento de Cursos**: Permite criar, obter, atualizar e deletar cursos.
- **Gerenciamento de Usuários**: Permite criar, obter, atualizar e deletar usuários.
- **Conexão com Banco de Dados**: Integração com um banco de dados para persistência de dados.
- **Rotas Configuráveis**: Configuração de rotas para gerenciar requisições HTTP.
- **Autenticação**: Middleware para proteger rotas sensíveis.
- **Containerização**: Pronto para rodar como microsserviço em Docker.

## Como Executar Localmente

1. Clone o repositório:

   ```sh
   git clone <URL_DO_REPOSITORIO>
   ```

2. Navegue até o diretório do projeto:

   ```sh
   cd plataforma-cursos
   ```

3. Instale as dependências:

   ```sh
   go mod tidy
   ```

4. Execute a aplicação:

   ```sh
   go run cmd/main.go
   ```

## Como Executar com Docker

1. Construa a imagem Docker:

   ```sh
   docker build -t plataforma-cursos .
   ```

2. Rode o container:

   ```sh
   docker run -p 8080:8080 plataforma-cursos
   ```

O serviço estará disponível em [http://localhost:8080](http://localhost:8080).

## Testes

Para rodar os testes automatizados:

```sh
go test ./internal/services/...
go test ./internal/controllers/...
```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.
