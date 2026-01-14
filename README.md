# Sistema de Almoxarifado TI 📦

Sistema completo de gerenciamento de estoque para o setor de TI, permitindo o controle eficiente de itens, quantidades e movimentações do almoxarifado.

## Sobre o Projeto

Este é um sistema web full-stack desenvolvido para gerenciar o inventário de equipamentos e materiais de TI. O sistema oferece funcionalidades completas de CRUD (Create, Read, Update, Delete) para controle de:

- **Cadastro de itens** (nome, quantidade, descrição, categoria)
- **Controle de entrada e saída** de produtos
- **Gerenciamento de estoque** em tempo real
- **Sistema de autenticação JWT** para administradores
- **Validação de dados** com regras de negócio
- **API RESTful** com CORS configurado

## Tecnologias Utilizadas

### Backend
- **Go 1.25.5** - Linguagem principal do backend
- **PostgreSQL 16** - Banco de dados relacional
- **JWT (golang-jwt/jwt/v5)** - Autenticação baseada em tokens
- **bcrypt** - Hash seguro de senhas
- **golang-migrate** - Gerenciamento de migrations
- **go-playground/validator** - Validação de structs
- **CORS** - Configuração de cross-origin

### DevOps
- **Docker** - Containerização da aplicação
- **Docker Compose** - Orquestração dos containers
- **Health checks** - Monitoramento de serviços

## Estrutura do Projeto

```
almoxarifado-ti/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go              # Entrypoint da aplicação
│   ├── config/
│   │   └── config.go                # Configurações da aplicação
│   ├── internal/
│   │   ├── database/
│   │   │   ├── postgres.go          # Conexão com PostgreSQL
│   │   │   ├── seed.go              # Seed de admin padrão
│   │   │   └── migrations/          # Migrations SQL
│   │   │       ├── 001_create_itens_table.sql
│   │   │       └── 002_create_admins_table.sql
│   │   ├── handlers/
│   │   │   ├── auth_handler.go      # Handlers de autenticação
│   │   │   └── item_handler.go      # Handlers de itens
│   │   ├── middleware/
│   │   │   └── auth.go              # Middleware JWT
│   │   ├── models/
│   │   │   ├── admin.go             # Model de Admin
│   │   │   └── item.go              # Model de Item
│   │   ├── repositories/
│   │   │   ├── admin_repository.go  # Acesso ao BD (admins)
│   │   │   └── item_repository.go   # Acesso ao BD (itens)
│   │   ├── services/
│   │   │   ├── auth_service.go      # Lógica de autenticação
│   │   │   └── item_service.go      # Lógica de negócio (itens)
│   │   └── utils/
│   │       └── validator.go         # Validador customizado
│   ├── dockerfile
│   ├── generate_hash.go             # Utilitário para gerar hash de senhas
│   └── go.mod
├── docker-compose.yml
└── README.md
```

## Funcionalidades

### Autenticação
- ✅ Login com email e senha
- ✅ JWT com expiração de 24 horas
- ✅ Hash bcrypt para senhas
- ✅ Middleware de autenticação em rotas protegidas
- ✅ Admin padrão criado automaticamente no primeiro boot

### Gerenciamento de Itens
- ✅ Criar novo item (nome, descrição, quantidade, categoria)
- ✅ Listar todos os itens
- ✅ Buscar item por ID
- ✅ Atualizar informações do item
- ✅ Deletar item
- ✅ Validação de campos obrigatórios
- ✅ Controle de quantidade (mínimo 0)

### Infraestrutura
- ✅ API RESTful com padrão de arquitetura limpa
- ✅ Migrations automáticas no startup
- ✅ Seed de admin padrão
- ✅ Health checks no PostgreSQL
- ✅ Sistema 100% containerizado com Docker
- ✅ CORS configurado para desenvolvimento

## API Endpoints

### Autenticação
```
POST /login
Body: {
  "email": "string",
  "password": "string"
}
Response: {
  "token": "jwt_token_here"
}
```

### Itens (Rotas protegidas - requerem JWT)
```
GET    /itens           # Listar todos os itens
GET    /itens/{id}      # Buscar item por ID
POST   /itens           # Criar novo item
PUT    /itens/{id}      # Atualizar item
DELETE /itens/{id}      # Deletar item

Body (POST/PUT): {
  "nome": "string",          # obrigatório
  "descricao": "string",
  "quantidade": int,         # obrigatório, >= 0
  "categoria": "string"      # obrigatório
}
```

### Header de Autenticação
```
Authorization: Bearer {jwt_token}
```

## Como Executar com Docker

### Pré-requisitos
- Docker
- Docker Compose

### Passo a passo

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
cd almoxarifado-ti
```

2. Crie um arquivo `.env` na raiz do projeto:
```env
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin123
POSTGRES_DB=almoxarifado

ADMIN_DEFAULT_EMAIL=admin@example.com
ADMIN_DEFAULT_PASSWORD=admin123

JWT_SECRET=sua_chave_secreta_muito_segura_aqui_min_32_chars
```

3. Execute os containers:
```bash
docker-compose up -d
```

4. A API estará disponível em:
```
http://localhost:8080
```

5. Para visualizar os logs:
```bash
docker-compose logs -f api
```

6. Para parar os containers:
```bash
docker-compose down
```

## Desenvolvimento Local (sem Docker)

### Pré-requisitos
- Go 1.25.5 ou superior
- PostgreSQL 16

### Setup

1. Configure as variáveis de ambiente ou crie um arquivo `.env`

2. Instale as dependências:
```bash
cd backend
go mod download
```

3. Execute a aplicação:
```bash
go run cmd/api/main.go
```

## Segurança

- Senhas armazenadas com hash bcrypt (custo 10)
- JWT com secret configurável via variável de ambiente
- Validação de entrada em todos os endpoints
- Middleware de autenticação em rotas protegidas
- CORS configurado (ajustar para produção)

## Ferramentas Úteis

### Gerar hash de senha
```bash
cd backend
go run generate_hash.go
```

## Observações

- O admin padrão é criado automaticamente no primeiro boot
- As migrations são executadas automaticamente
- O banco de dados persiste os dados em um volume Docker
- A API roda na porta 8080
- O PostgreSQL roda na porta 5432
