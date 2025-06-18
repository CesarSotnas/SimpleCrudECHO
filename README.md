# 🛠️ GinEchoCrud

Sistema de autenticação e gerenciamento de usuários feito com Go, Echo Framework, SQLite e JWT.
Organizado em camadas (Controller, Service, Repository) com injeção de dependência, resposta padrão e segurança via token.

---

## 🚀 Tecnologias Utilizadas

- [Golang](https://golang.org/)
- [Echo Framework](https://echo.labstack.com/)
- [SQLite](https://www.sqlite.org/index.html)
- [JWT (JSON Web Tokens)](https://jwt.io/)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [godotenv](https://github.com/joho/godotenv)

---

## 🗂️ Estrutura do Projeto

GinEchoCrud/
├── cmd/ # Entrada de ferramentas auxiliares (ex: gerador de hash)
├── dto/ # Data Transfer Objects (ex: LoginRequest)
├── internal/
│ ├── controller/ # Camada de controle (Echo handlers)
│ ├── service/ # Regras de negócio (Services)
│ ├── repository/ # Acesso ao banco (SQL, queries)
│ ├── middleware/ # Middlewares (ex: JWT)
│ ├── helpers/ # Funções auxiliares (JWT, respostas padrão)
│ ├── interfaces/ # Contratos para injeção de dependência
│ ├── models/ # Structs que representam entidades
│ └── database/ # Inicialização e conexão com o banco SQLite
├── .env # Variáveis de ambiente (não versionado)
├── go.mod / go.sum # Gerenciador de dependências do Go

---

## 🔐 Autenticação com JWT

- Após realizar o `POST /login`, o usuário receberá um token.
- Esse token deve ser enviado no cabeçalho `Authorization: Bearer <token>` nas rotas protegidas.

---

## 🧪 Testando com cURL/Postman

### 🔑 Login

```
curl --request POST http://localhost:1323/login \
--header "Content-Type: application/json" \
--data '{
  "email": "email@email.com",
  "password": "senhaSegura"
}'
```

```
curl --request GET http://localhost:1323/users \
--header "Authorization: Bearer <seu_token_aqui>"
```

```
JWT_SECRET=sua_chave_super_secreta
DEFAULT_PASSWORD=123456
PORT=:1323
```

```
CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  age INTEGER,
  email TEXT UNIQUE,
  password TEXT
);
```

```
-- Inserção com senha hash
INSERT INTO users (name, age, email, password)
VALUES ("Carlos", 30, "carlos@email.com", "$2a$10$abc...xyz");
```

go run cmd/tools/hashgen.go minhaSenhaSegura


✨ Próximos passos

Implementar refresh token

Documentar API com Swagger

Adicionar testes de unidade e integração

    Substituir SQLite por PostgreSQL para produção

🧑‍💻 Autor

Desenvolvido por Cesar Santos
Feito com foco em segurança, organização e aprendizado contínuo.


