# 🚀 Desafio: Construa uma API em Go para Gestão de Tarefas

Um desafio completo para praticar Go, APIs RESTful e conceitos modernos de desenvolvimento.

## 📋 Nível 1: CRUD Básico

### Requisitos:
- **Rotas**:
  - `POST /tasks` - Criar nova tarefa
  - `GET /tasks` - Listar todas tarefas (com filtro por status)
  - `GET /tasks/:id` - Buscar tarefa por ID
  - `PUT /tasks/:id` - Atualizar tarefa
  - `DELETE /tasks/:id` - Remover tarefa

### Estrutura:
```go
type Task struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"` // "pending" ou "completed"
    CreatedAt   time.Time `json:"created_at"`
}
Tecnologias:
Framework: Gin ou Echo

Armazenamento: In-memory (map/slice) ou SQLite

🔐 Nível 2: Autenticação JWT
Requisitos:
POST /register - Registrar novo usuário

POST /login - Autenticar e receber token

Todas rotas /tasks devem exigir token JWT

Segurança:
Hash de senha com bcrypt

Middleware de autenticação

🗃️ Nível 3: Banco de Dados Real
Requisitos:
Usar PostgreSQL ou MySQL

Implementar paginação (page, limit)

Migrações de banco de dados

Pacotes recomendados:
GORM ou sqlx

goose ou migrate para migrações

🧪 Nível 4: Testes e Otimização
Requisitos:
Testes unitários com httptest

Testes de integração com banco de dados

Implementar cache com Redis

Logs estruturados com zap

🎯 Nível 5: Extras (Opcional)
Desafios avançados:
Notificações em tempo real com WebSocket

Dockerize a aplicação e banco de dados

Documentação com Swagger

Deploy em cloud (Render/Heroku)

📂 Estrutura Recomendada
Copy
/projeto
├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   └── middleware/
├── pkg/
│   ├── auth/
│   └── database/
├── migrations/
├── Dockerfile
└── README.md
⏳ Entrega Esperada
Código fonte no GitHub

README com instruções de execução

Coleção Postman/Insomnia

(Opcional) Link para API deployada

💡 Dicas
Comece simples e vá incrementando

Documente seu progresso

Considere usar Makefile para comandos repetitivos

Dificuldade: ⭐⭐⭐☆ (Intermediário)
Tempo estimado: 6-10 horas

Pronto para o desafio? 🚀

