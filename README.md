## 1. Go - Echo Crud

Este projeto é uma API RESTful escrita em Go, utilizando o framework [Echo](https://echo.labstack.com/) e banco de dados SQLite. A aplicação implementa autenticação via JWT e segue uma arquitetura organizada por camadas (controller, service, repository, interfaces).

## ✨ Tecnologias e bibliotecas

- **Go (Golang)**
- **Echo v4** – framework web
- **SQLite** – banco de dados leve
- **JWT** – autenticação segura
- **bcrypt** – hash de senhas
- **godotenv** – carregamento de variáveis de ambiente

---

## 📁 Estrutura do projeto
Abaixo está a estrutura organizada do projeto, seguindo boas práticas de separação de responsabilidades:

    cmd/
    Ponto de entrada principal da aplicação (main.go).

    internal/
    Módulos internos da aplicação:

        controller/: Camada de controle (Echo handlers).

        service/: Contém as regras de negócio (lógica da aplicação).

        repository/: Responsável pelo acesso ao banco de dados (SQLite).

        database/: Scripts de migração e conexão com o banco.

        helpers/: Utilitários como geração/validação de JWT e helpers de resposta HTTP.

        interfaces/: Interfaces utilizadas para injeção de dependência entre camadas.

        middleware/: Middlewares HTTP, como autenticação JWT.

        models/: Estruturas de dados e modelos do domínio.

    dto/
    Objetos de entrada (Data Transfer Objects), como LoginRequest.

    .env
    Arquivo para configuração de variáveis de ambiente (não deve ser versionado).

    go.mod / go.sum
    Gerenciamento de dependências do projeto Go.


---

## 🔐 Autenticação

- A autenticação é feita via **JWT**.
- O token deve ser enviado no cabeçalho `Authorization`

---

## 🚀 Como rodar localmente

### 1. Clone o projeto

```bash
git clone https://github.com/seu-usuario/GinEchoCrud.git
cd GinEchoCrud
```

## 2. Crie o arquivo .env
    JWT_SECRET=sua_chave_super_secreta
    DEFAULT_PASSWORD=123456
    PORT=1323


## 3. Rode a aplicação
go run cmd/main.go

A API ficará disponível em:
http://localhost:1323

## 📦 Endpoints disponíveis:
| Método | Rota   | Autenticado | Descrição                    |
| ------ | ------ | ----------- | ---------------------------- |
| POST   | /login | ❌           | Login e geração de token JWT |
| GET    | /users | ✅           | Lista todos os usuários      |



##🧪 Futuras implementações

    Cadastro de usuários (POST /users)

    Atualização e deleção de usuários

    Testes automatizados com Testify ou GoMock

    Suporte para PostgreSQL e Docker

