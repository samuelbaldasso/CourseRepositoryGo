# Plataforma de Cursos

Este projeto é uma plataforma de cursos desenvolvida em GoLang. O objetivo é fornecer uma aplicação que permita a criação, leitura, atualização e exclusão de cursos.

## Estrutura do Projeto

```text
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
├── docker-compose.yml             # Orquestração de containers para app e banco
├── db_init.sql                    # Script de inicialização do banco
└── README.md                      # Documentação do projeto
```

## Funcionalidades

- **Gerenciamento de Cursos**: Permite criar, obter, atualizar e deletar cursos.
- **Gerenciamento de Usuários**: Permite criar, obter, atualizar e deletar usuários.
- **Conexão com Banco de Dados**: Integração com Postgres para persistência de dados.
- **Rotas Configuráveis**: Configuração de rotas para gerenciar requisições HTTP.
- **Autenticação**: Middleware para proteger rotas sensíveis.
- **Tratamento Global de Erros**: Exception handler global para respostas padronizadas.
- **Injeção de Dependências**: Utiliza Google Wire para DI em toda a aplicação e testes.
- **Testes Automatizados**: Testes de services e controllers integrados ao DI, sem setup manual de banco.
- **Containerização**: Pronto para rodar como microsserviço em Docker e Docker Compose.

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

4. Crie um arquivo `.env` na raiz do projeto com as variáveis necessárias (exemplo no repositório).

5. Execute o banco de dados Postgres (pode usar o docker-compose):

   ```sh
   docker-compose up -d db
   ```

6. Execute a aplicação:

   ```sh
   go run cmd/main.go
   ```

## Como Executar com Docker Compose

1. Suba toda a stack (app + banco):

   ```sh
   docker-compose up --build
   ```

O serviço estará disponível em [http://localhost:8080](http://localhost:8080).

## Testes

Os testes automatizados de services e controllers usam DI e inicialização automática do banco via o método init do pacote `internal/services/test_helpers.go`.

Para rodar todos os testes:

```sh
go test ./internal/services/...
go test ./internal/controllers/...
```

Não é necessário setup manual de banco ou variáveis de ambiente nos testes: tudo é carregado automaticamente.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.
