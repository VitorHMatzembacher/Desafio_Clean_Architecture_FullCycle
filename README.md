# Desafio Clean Architecture - Orders

Esta aplicação implementa um use case de listagem e criação de pedidos (**orders**) seguindo os princípios de *Clean Architecture*, expondo:

- **REST API** (HTTP/JSON)
- **GraphQL API**
- **gRPC API**

Toda a infraestrutura roda com Docker Compose, bastando um único comando para levantar banco, migrations e serviços.

---

## 🚀 Pré-requisitos

- Docker e Docker Compose instalados na máquina
- (Opcional) VS Code com extensão HTTP Client para testar os arquivos `api/orders.http`

---

## 🔧 Configuração

1. Clone este repositório:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd Desafio_Clean_Architecture_FullCycle
   ```

2. As variáveis de ambiente já estão definidas em `.env` com valores padrões para PostgreSQL:
   ```dotenv
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=postgres
   POSTGRES_DB=orders

   HTTP_PORT=8080
   GRPC_PORT=50051

   DB_HOST=db
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=orders
   ```

---

## 🐳 Subindo a aplicação

Para subir a aplicação pela primeira vez (ou sempre que quiser recriar o banco e reexecutar as migrations), rode:
```bash
docker-compose down -v   # opcional na primeira vez ou se precisar reaplicar migrations
docker-compose up
```

Se você já subiu uma vez e não alterou as migrations, basta:
```bash
docker-compose up
```

- **PostgreSQL**: container `orders_postgres`, porta **5432**
- **Adminer**: em http://localhost:8081 (usuário/senha: `postgres` / `postgres`, database: `orders`)
- **App (REST/GraphQL)**: em http://localhost:8080
- **gRPC**: porta **50051**

---

## 📦 Estrutura de pastas

```
├── cmd/ordersystem        # entrypoint da aplicação
├── configs                # carregamento de configurações e conexão ao DB
├── db/migrations          # arquivos SQL para criação de tabelas
├── internal/
│   ├── entity             # definição de entity Order e interface de repositório
│   ├── infra/
│   │   ├── database       # implementação SQL de OrderRepository
│   │   ├── web            # handlers REST + configuração de rotas
│   │   ├── graph          # resolvers e schema GraphQL
│   │   └── grpc           # serviço gRPC e protobufs
│   └── usecase           # regras de negócio (CreateOrder, ListOrders)
├── api/orders.http        # exemplos de requisições REST e GraphQL
├── Dockerfile
├── docker-compose.yaml
└── README.md
```

---

## 🛠️ Endpoints REST

Veja exemplos em `api/orders.http` ou diretamente:

- **POST /order**
  - Cria um pedido
  - **Body**: `{ "price": float, "tax": float }`
  - **Resposta**: `201 Created` com o objeto Order (contendo `id`, `price`, `tax`, `finalPrice`)

- **GET /order**
  - Lista todos os pedidos
  - **Resposta**: `200 OK` com array de Orders

---

## 🧩 GraphQL

- **Endpoint**: `POST /query` (envia-se uma requisição JSON com o campo `query` contendo a string da query GraphQL)

  **Exemplo de corpo**:
  ```json
  { "query": "{ listOrders { id price tax finalPrice } }" }
  ```

*Mais adiante, você encontrará também o GraphQL Playground como uma opção extra.*

---

## 📡 gRPC

- **Endpoint**: `localhost:50051`
- **Serviço**: `OrderService`
  - `CreateOrder(CreateOrderRequest) returns (OrderResponse)`
  - `ListOrders(ListOrdersRequest) returns (ListOrdersResponse)`


---

## ✅ Testes

- **API HTTP**: use `api/orders.http` no VS Code (clique em "Send Request").
- **Adminer**: confirme a tabela `orders` em http://localhost:8081.
- **gRPC**: teste com `grpcurl -plaintext` se desejar.

---

## 📄 Licença

Este projeto está disponível sob a [MIT License](LICENSE).
