
# Processo seletivo Loc Pay

O projeto consiste em uma api que envolve recebedores e operações para antecipações de alugueis com uma taxa de 3%.


## Informações importantes: 

⚠️ Este projeto está configurado para conectar em um banco de teste do Supabase, com dados fictícios apenas para demonstração.

Como não foi especificado nas instruções para que fizesse uma rota para criar recebedores, criei no proprio banco de dados um recebedor teste que será usado para validar todas as rotas. Seu id é:

```
4fc2f2fd-a0a3-4f13-befa-2fa952a27a29
```

Também deixei um postman com todas as rotas prontas para testar caso prefiram para ter uma vizualização melhor:

```
https://planetary-eclipse-621566.postman.co/workspace/locpay~245be6ae-ba52-4720-a2b3-09bf40b32f33/request/28160737-50e08f9f-0456-45f1-b2d7-bad50f8dc089
```

## Tecnologias utilizadas:

- Go(Golang)
- Gin Framework
- Supabase

## Pré Requisitos:

#### Go:
O link abaixo tem um tutoriais para Linux, Mac e Windows

```
https://go.dev/doc/install
```

## Sobre as rotas

| Rota | Retorno |
|-----------|------------|
| `POST /api/v1/operations` | Retorna o DTO da operação para não expor dados sensíveis e o id da operação para continuar os testes |
| `POST /api/v1/operations/:id/confirm` | Retorna os status da operação e uma mensagem de confirmação |
| `GET /api/v1/operations` | Retorna os dados completos da operação |
| `GET /api/v1/receivers/:id` | Retorna saldo, nome e histórico de operações do recebedor |




## Como testar


#### Clone o repositório:

```
git clone https://github.com/Nadoutti/locpay_ps.git
```
#### Instale as dependencias:

```
go mod tidy
```

#### Crie um arquivo .env e preencha com as chaves anexadas no email: 

```
# Servidor
PORT=8080

# Conexão com supabase
SUPABASE_URL=https://xxxx.supabase.co
SUPABASE_KEY="SUA_CHAVE_AQUI"

```

#### Rode o projeto

```
go run ./cmd/server/main.go
```

