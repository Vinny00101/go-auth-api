## Descrição do Projeto

Este projeto é uma API REST desenvolvida em Go (Golang) com o framework Gin e o ORM GORM para persistência de dados em SQLite.
A API implementa funcionalidades básicas de login e registro de usuários, servindo como ponto de partida para sistemas maiores que necessitam de autenticação.

## Funcionalidades

- Registro de usuário com nome, email e senha.

- Login de usuário validando credenciais no banco.

- Migração automática de tabelas com GORM.

- Estrutura preparada para evoluir com JWT e middlewares de segurança.

## Arquitetura

- A API foi construída utilizando uma arquitetura organizada em camadas:

- Handlers (Controllers): recebem as requisições HTTP e retornam respostas.

- Services: contêm a lógica de negócio (ex: validações, regras de cadastro e login).

- Repository (Database): responsável pelo acesso ao banco de dados via GORM.

Models: definem a estrutura das entidades (ex: User).

Essa separação segue o princípio Single Responsibility (SRP) do SOLID, garantindo que cada parte do sistema tenha apenas uma responsabilidade clara. Isso facilita a manutenção, a testabilidade e a evolução do código.

## Tecnologias Utilizadas

- linguagem principal: `Go`

- framework web rápido e minimalista: `Gin`

- ORM para manipulação do banco: `GORM`

- banco de dados leve e prático: `SQLite`

### Segurança
Para garantir a segurança das senhas dos usuários, o projeto utiliza a biblioteca `bcrypt` para hash e verificação de senhas. As senhas nunca são armazenadas em texto puro no banco de dados.

Para melhorar ainda mais a segurança, recomenda-se implementar autenticação baseada em tokens JWT (JSON Web Tokens) para proteger as rotas que exigem autenticação. Isso pode ser feito adicionando middlewares que validam o token em cada requisição.

Para proteger contra ataques comuns, como SQL Injection, o uso do GORM já ajuda a mitigar esses riscos ao utilizar consultas parametrizadas. Alem disso, é importante validar e sanitizar todas as entradas do usuário.

## Próximos Passos

Implementar autenticação com JWT.

Criar middlewares para proteger rotas privadas.

Adicionar testes automatizados.

Expandir a arquitetura para suportar novos módulos (ex: perfis, permissões).