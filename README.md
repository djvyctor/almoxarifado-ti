# Corvi.com

Sistema de gerenciamento de inventário e controle de estoque.

## Tecnologias

### Backend
- **Go** (1.25.5)
- **PostgreSQL** 16.11
- **JWT** para autenticação
- **CSRF** protection
- **Docker** & Docker Compose

### Frontend
- **Vue.js 3** (Composition API)
- **Vite** (Build tool)
- **Tailwind CSS v4**
- **Pinia** (State management)
- **Vue Router**

## Estrutura

```
├── backend/          # API em Go
│   ├── cmd/         # Entry points
│   ├── internal/    # Código da aplicação
│   └── config/      # Configurações
├── frontend/        # Interface Vue.js
│   └── src/
│       ├── views/   # Páginas
│       ├── stores/  # Gerenciamento de estado
│       └── services/# Serviços de API
└── docker-compose.yml
```

## Como rodar

```bash
# Subir containers
docker-compose up -d

# Frontend (http://localhost:5173)
cd frontend
npm install
npm run dev

# Backend (http://localhost:8080)
cd backend
go run cmd/api/main.go
```

## Acesso padrão

- **Email:** admin@corvi.com
- **Senha:** admin123

---