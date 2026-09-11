# SGA - Sistema de Gestão de Alocação

API REST desenvolvida em Go utilizando o framework Gin para gerenciamento de salas, alunos, turmas, matrículas e alocação de horários.

O sistema foi desenvolvido para atender às regras de negócio relacionadas à capacidade das salas, matrícula de alunos e prevenção de conflitos de horário.

---

## Tecnologias utilizadas

- Go
- Gin
- API REST
- Armazenamento em memória

---

## Estrutura do projeto

```text
APLICAÇÃO GO/
│
├── main.go
├── models.go
├── routes.go
├── storage.go
├── aluno_h.go
├── sala_h.go
├── turma_h.go
├── util.go
├── go.mod
├── go.sum
└── .gitignore
