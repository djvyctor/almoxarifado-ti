# Sistema de Almoxarifado TI 📦

Sistema completo de gerenciamento de estoque e patrimônio para o setor de TI, permitindo o controle eficiente de itens, usuários, quantidades e movimentações do almoxarifado com interface moderna e intuitiva.

## 📸 Capturas de Tela

### Tela de Login

![Tela de Login](./docs/images/login.png)

### Dashboard Analítico

![Dashboard](./docs/images/dashboard.png)

### Gerenciamento de Estoque

![Estoque](./docs/images/estoque.png)

### Gerenciamento de Usuários

![Usuários](./docs/images/usuarios.png)

---

## 📋 Sobre o Projeto

Sistema web full-stack desenvolvido para gerenciar o inventário de equipamentos e materiais de TI com recursos avançados de controle patrimonial, vinculação de itens a usuários e analytics em tempo real.

## ✨ Funcionalidades

### 🔐 Autenticação
- Login seguro com JWT
- Proteção de rotas
- Persistência de sessão
- Interface com glassmorphism e animações

### 📦 Gestão de Itens
- CRUD completo de itens do estoque
- Campos: nome, categoria, patrimônio (5 dígitos), valor monetário, quantidade, descrição, usuário vinculado
- **10 categorias predefinidas** com código de cores:
  - Periféricos (azul)
  - Hardware (roxo)
  - Cabos e Adaptadores (amarelo)
  - Monitores (ciano)
  - Computadores (índigo)
  - Notebooks (rosa)
  - Impressoras (laranja)
  - Rede (verde)
  - Armazenamento (teal)
  - Outros (cinza)
- Autocomplete de categorias em tempo real
- Autocomplete de usuários para vinculação
- Validação de patrimônio (5 dígitos numéricos)
- Formatação automática de valores monetários (R$)
- Descrições com placeholders dinâmicos por categoria
- Capitalização automática do nome

### 👥 Gestão de Usuários
- CRUD completo de usuários
- Campos: nome, email (opcional), departamento
- **12 departamentos predefinidos**:
  - Administrativo, RH, TI, Telemarketing
  - Gerentes, Lider, Segurança, Marketing
  - Faxina, Operacional, Estrategico, Recepção
- Autocomplete de departamentos
- Filtro por departamento
- Busca por nome, email ou departamento
- **Exclusão em cascata**: ao deletar usuário, remove todos itens vinculados
- Modal de confirmação mostrando quantos itens serão excluídos
- Capitalização automática da primeira letra do nome

### 📊 Dashboard Analítico (Power BI Style)
- **Total de Itens**: soma das quantidades de todos os itens
- **Estoque Baixo**: alerta de itens com quantidade < 5
- **Valor Total Estimado**: cálculo automático (valor × quantidade)
- **Distribuição por Categoria**: barras horizontais com:
  - Gradientes coloridos por categoria
  - Percentuais calculados automaticamente
  - Quantidade de unidades por categoria
  - Ordenação por quantidade (maior → menor)

### 🔍 Filtros e Buscas
- Busca em tempo real por nome/categoria nos itens
- Filtro dropdown de categorias no estoque
- Busca por nome/email/departamento nos usuários
- Filtro dropdown de departamentos

### 💾 Outros Recursos
- Modal customizado para exclusão com animações
- Feedback visual (loading, erros, sucesso)
- Design responsivo
- Tema consistente (roxo/índigo)
- Tailwind CSS v4 com animações
- Validações client-side e server-side

---

## 🛠️ Tecnologias Utilizadas

### Backend
- **Go 1.25.5** - Linguagem principal
- **PostgreSQL 16** - Banco de dados relacional
- **JWT (golang-jwt/jwt/v5)** - Autenticação com tokens
- **bcrypt** - Hash seguro de senhas
- **golang-migrate** - Gerenciamento de migrations
- **go-playground/validator v10** - Validação de structs
- **rs/cors** - Configuração CORS
- **lib/pq** - Driver PostgreSQL

### Frontend
- **Vue.js 3.5.26** - Framework JavaScript (Composition API)
- **Vite 7.3.1** - Build tool e dev server
- **Pinia 3.0.4** - Gerenciamento de estado
- **Vue Router 4.6.4** - Roteamento SPA
- **Axios 1.13.2** - Cliente HTTP
- **Tailwind CSS v4.1.18** - Framework CSS
- **PostCSS** - Processamento CSS

### DevOps
- **Docker** - Containerização
- **Docker Compose** - Orquestração de containers
- **Health checks** - Monitoramento de serviços

---

## 📁 Estrutura do Projeto

```
almoxarifado-ti/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go              # Entrypoint da aplicação
│   ├── config/
│   │   └── config.go                # Configurações (env vars)
│   ├── internal/
│   │   ├── database/
│   │   │   ├── postgres.go          # Conexão PostgreSQL
│   │   │   ├── seed.go              # Seed admin padrão
│   │   │   └── migrations/
│   │   │       ├── 001_create_itens_table.sql
│   │   │       ├── 002_create_admins_table.sql
│   │   │       └── 003_create_users_table.sql
│   │   ├── handlers/
│   │   │   ├── auth_handler.go      # Handlers autenticação
│   │   │   ├── item_handler.go      # Handlers itens
│   │   │   └── user_handler.go      # Handlers usuários
│   │   ├── middleware/
│   │   │   ├── auth.go              # Middleware JWT
│   │   │   └── rate_limiter.go      # Rate limiting
│   │   ├── models/
│   │   │   ├── admin.go             # Model Admin
│   │   │   ├── item.go              # Model Item
│   │   │   └── user.go              # Model User
│   │   ├── repositories/
│   │   │   ├── admin_repository.go  # Repo Admin
│   │   │   ├── item_repository.go   # Repo Item (CRUD)
│   │   │   └── user_repository.go   # Repo User (CRUD + cascata)
│   │   ├── services/
│   │   │   ├── auth_service.go      # Service autenticação
│   │   │   ├── item_service.go      # Service itens
│   │   │   └── user_service.go      # Service usuários
│   │   └── utils/
│   │       └── validator.go         # Validações customizadas
│   ├── dockerfile
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── assets/
│   │   │   ├── base.css
│   │   │   └── main.css
│   │   ├── components/
│   │   │   └── AppLayout.vue        # Layout com sidebar
│   │   ├── config/
│   │   │   └── api.js               # Config Axios
│   │   ├── router/
│   │   │   └── index.js             # Rotas Vue Router
│   │   ├── services/
│   │   │   └── api.js               # Instância Axios
│   │   ├── stores/
│   │   │   ├── auth.js              # Store autenticação
│   │   │   ├── items.js             # Store itens
│   │   │   └── users.js             # Store usuários
│   │   ├── views/
│   │   │   ├── AboutView.vue
│   │   │   ├── HomeView.vue         # Dashboard
│   │   │   ├── ItemsView.vue        # Gestão estoque
│   │   │   ├── LoginView.vue        # Tela login
│   │   │   └── UsersView.vue        # Gestão usuários
│   │   ├── App.vue
│   │   └── main.js
│   ├── eslint.config.js
│   ├── index.html
│   ├── jsconfig.json
│   ├── package.json
│   ├── postcss.config.js
│   ├── tailwind.config.js
│   └── vite.config.js
├── docker-compose.yml
└── README.md
```

---

## 🚀 Como Executar o Projeto

### Pré-requisitos
- Docker e Docker Compose instalados
- Portas 8080 (backend) e 5173 (frontend) disponíveis

### Passo a Passo

1. **Clone o repositório**
```bash
git clone <url-do-repositorio>
cd almoxarifado-ti
```

2. **Configure as variáveis de ambiente (opcional)**

Crie um arquivo `.env` na pasta `backend/` (ou use as variáveis padrão):
```env
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha_segura
DB_NAME=almoxarifado

# JWT
JWT_SECRET=sua_chave_jwt_secreta_minimo_32_caracteres

# Admin padrão
ADMIN_DEFAULT_EMAIL=seu_email@exemplo.com
ADMIN_DEFAULT_PASSWORD=sua_senha_forte_aqui

# App
APP_PORT=8080
```

3. **Inicie os containers**
```bash
docker-compose up -d
```

Isso irá:
- ✅ Criar o container PostgreSQL
- ✅ Executar as migrations automaticamente
- ✅ Criar o admin padrão
- ✅ Iniciar o backend na porta 8080

4. **Instale as dependências do frontend**
```bash
cd frontend
npm install
```

5. **Inicie o servidor de desenvolvimento**
```bash
npm run dev
```

6. **Acesse o sistema**
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- Login padrão:
  - **Email**: (definido no arquivo .env)
  - **Senha**: (definida no arquivo .env)

---

## 📊 Banco de Dados

### Schema

**Tabela `items`**
```sql
- id (UUID, PK)
- name (TEXT, NOT NULL)
- category (TEXT, NOT NULL)
- patrimony (TEXT, opcional)
- assigned_to (TEXT, opcional)
- description (TEXT, opcional)
- value (DECIMAL(10,2), opcional)
- quantity (INTEGER, NOT NULL, CHECK >= 0)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

**Tabela `users`**
```sql
- id (UUID, PK)
- name (TEXT, NOT NULL)
- email (TEXT, UNIQUE, opcional)
- department (TEXT, opcional)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

**Tabela `admins`**
```sql
- id (UUID, PK)
- email (TEXT, UNIQUE, NOT NULL)
- password_hash (TEXT, NOT NULL)
- created_at (TIMESTAMP)
- updated_at (TIMESTAMP)
```

---

## 🔌 Endpoints da API

### Autenticação
- `POST /auth/login` - Login de administrador

### Itens
- `GET /items` - Lista todos os itens
- `POST /items` - Cria novo item
- `GET /items/:id` - Busca item por ID
- `PUT /items/:id` - Atualiza item
- `DELETE /items/:id` - Remove item

### Usuários
- `GET /users` - Lista todos os usuários
- `POST /users` - Cria novo usuário
- `GET /users/:id` - Busca usuário por ID
- `PUT /users/:id` - Atualiza usuário
- `DELETE /users/:id` - Remove usuário (e itens vinculados)
- `GET /users/linked-items?name=` - Conta itens vinculados ao usuário

### Health Check
- `GET /health` - Status da API
- `GET /health/db` - Status do banco de dados

> **Nota**: Todas as rotas (exceto `/auth/login` e `/health*`) requerem autenticação JWT via header `Authorization: Bearer <token>`

---

## 🎨 Paleta de Cores (Tailwind)

- **Primária**: Purple (roxo) - `purple-600`, `purple-700`
- **Secundária**: Indigo (índigo) - `indigo-600`, `indigo-700`
- **Neutro**: Slate (cinza) - `slate-50` até `slate-900`
- **Sucesso**: Green - `green-50`, `green-700`
- **Erro**: Red - `red-50`, `red-600`
- **Alerta**: Amber - `amber-50`, `amber-600`

---

## 🔒 Segurança

- ✅ Senhas criptografadas com bcrypt (custo 10)
- ✅ Autenticação JWT com expiração de 24h
- ✅ Middleware de autenticação em rotas protegidas
- ✅ Rate limiting no endpoint de login
- ✅ CORS configurado para origens específicas
- ✅ Validação de dados no backend e frontend
- ✅ SQL injection prevenido (prepared statements)
- ✅ NULL handling correto em campos opcionais

---

## 📦 Comandos Úteis

### Docker
```bash
# Parar containers
docker-compose down

# Remover volumes (limpa banco de dados)
docker-compose down -v

# Rebuild sem cache
docker-compose build --no-cache

# Ver logs do backend
docker logs almoxarifado_api

# Ver logs do postgres
docker logs almoxarifado_postgres

# Acessar PostgreSQL
docker exec -it almoxarifado_postgres psql -U <seu_usuario> -d almoxarifado
```

### Frontend
```bash
# Desenvolvimento
npm run dev

# Build para produção
npm run build

# Preview da build
npm run preview

# Lint
npm run lint
```

### Backend
```bash
# Gerar hash de senha
go run generate_hash.go

# Testes (se implementados)
go test ./...

# Build manual
go build -o api ./cmd/api
```

---

## 🐛 Troubleshooting

### Backend não inicia
- Verifique se as portas 8080 e 5432 estão disponíveis
- Confira as variáveis de ambiente no `.env`
- Veja os logs: `docker logs almoxarifado_api`

### Erro de migração
```bash
# Limpar banco e recriar
docker-compose down -v
docker-compose up -d
```

### Frontend não conecta ao backend
- Verifique se o backend está rodando na porta 8080
- Confirme a URL da API em `frontend/src/config/api.js`
- Verifique as configurações CORS no backend

### Erro de autenticação
- Limpe o localStorage do navegador (F12 → Application → Local Storage)
- Verifique se o token JWT não expirou (24h)
- Confirme as credenciais de login

---

## 📝 Licença

Este projeto é de uso interno e proprietário.

---

## 👥 Autor

Desenvolvido para gerenciamento eficiente de almoxarifado TI.

---

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'feat: adiciona nova feature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

### Padrão de Commits
- `feat:` Nova funcionalidade
- `fix:` Correção de bug
- `docs:` Documentação
- `style:` Formatação, ponto e vírgula faltando, etc
- `refactor:` Refatoração de código
- `test:` Adicionando testes
- `chore:` Atualização de tarefas de build, configurações, etc

---

**🎯 Sistema desenvolvido com foco em usabilidade, performance e manutenibilidade.**
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
POSTGRES_USER=seu_usuario
POSTGRES_PASSWORD=sua_senha_segura
POSTGRES_DB=almoxarifado

ADMIN_DEFAULT_EMAIL=seu_email@exemplo.com
ADMIN_DEFAULT_PASSWORD=sua_senha_forte_aqui

JWT_SECRET=sua_chave_jwt_secreta_minimo_32_caracteres
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
