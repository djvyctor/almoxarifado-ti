# Sistema de Almoxarifado TI 📦

Sistema completo de gerenciamento de estoque para o setor de TI, permitindo o controle eficiente de itens, quantidades e movimentações do almoxarifado.

##  Sobre o Projeto

Este é um sistema web full-stack desenvolvido para gerenciar o inventário de equipamentos e materiais de TI. O sistema oferece funcionalidades completas de CRUD (Create, Read, Update, Delete) para controle de:

- **Cadastro de itens** (nome, quantidade, descrição)
- **Controle de entrada e saída** de produtos
- **Gerenciamento de estoque** em tempo real
- **Sistema de autenticação** para administradores
- **Dashboard** para visualização dos dados

##  Tecnologias Utilizadas

### Backend
- **Go (Golang)** - Linguagem principal do backend
- **PostgreSQL** - Banco de dados relacional
- **JWT** - Autenticação baseada em tokens

### Frontend
- Aplicação web moderna e responsiva

### DevOps
- **Docker** - Containerização da aplicação
- **Docker Compose** - Orquestração dos containers

##  Estrutura do Projeto

```
almoxarifado-ti/
├── backend/          # API REST em Go
│   ├── cmd/         # Entrypoint da aplicação
│   ├── internal/    # Lógica de negócio
│   │   ├── handlers/     # Controllers HTTP
│   │   ├── services/     # Regras de negócio
│   │   ├── repositories/ # Acesso ao banco de dados
│   │   ├── models/       # Estruturas de dados
│   │   └── middleware/   # Middlewares (auth, cors, etc)
│   └── config/      # Configurações
├── frontend/        # Interface do usuário
└── docker/          # Arquivos Docker
```

##  Funcionalidades

- ✅ Autenticação de administradores
- ✅ CRUD completo de itens
- ✅ Controle de quantidade em estoque
- ✅ Histórico de movimentações
- ✅ API RESTful documentada
- ✅ Sistema containerizado com Docker
