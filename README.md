# COMO RODAR O PROJETO

- É necessário ter o docker e docker compose instalados, assim como conseguir rodar shell scripts no seu sistema operacional com sudo ou como super user.

1. Clone o repositório
2. Entre na pasta do projeto
3. Rode o comando `sudo ./up.sh` para inicializar os containers pelo docker compose
4. Rode o comando `sudo ./down.sh` para parar os containers pelo docker compose

# Rotas

- GET / - Rota de healthcheck
- GET /create - Rota de criação de usuário, recebe os parametros `email`, `password`, `cpf_cnpj`, `name` e `user_type`, onde `user_type` pode ser `store` ou `person` retornando um objeto com `message` informando a resposta da requisição. Exemplo de requisição: `curl --location --request POST 'http://127.0.0.1:8080/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "usuario@email.com",
    "password": "passwor1d@",
    "cpf_cnpj": "4213423",
    "name": "asd",
    "user_type": "person"
}'`

- POST /login - Rota de login, recebe os parametros `email` e `password` retornando `message` com o status da requisição e `token` com o token de acesso que deve ser informado no header `Authorization`. Exemplo de requisição: `curl --location --request POST 'http://127.0.0.1:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "person2@email.com",
    "password": "passwor1d@"
}'`

- GET /balance - Rota de saldo, recebe o token de acesso no header `Authorization` e retorna `message` com o status da requisição e `balance` com o saldo. Exemplo de requisição: `curl --location --request GET 'http://localhost:8080/balance' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywidXNlcl90eXBlIjoicGVyc29uIiwiZXhwIjoxNjc1MDg2Njk2fQ.jCkSV-ydo3xUn5rpUyLfcCVL9Lxyh2qoZ92kDo8fSuY'`

- POST /transaction - Rota de transação, recebe o token de acesso no header `Authorization`, `value` e `payee` e retorna `message` com o status da requisição. Exemplo de requisição: `curl --location --request POST 'http://127.0.0.1:8080/transaction' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywidXNlcl90eXBlIjoicGVyc29uIiwiZXhwIjoxNjc1MTI0NjA2fQ._COzsVii50xy-VrZtXctUoKy4UofPBTV6hQKUtNttiw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "value": 0.99,
    "payee": 2
}'`
