package main

type TypeRecurso string

const (
	Projetor         TypeRecurso = "Projetor"
	Janela           TypeRecurso = "Janela"
	AlarmeDeIncendio TypeRecurso = "Alarme de Incendio"
	Computador       TypeRecurso = "Computador"
	ArCondicionado   TypeRecurso = "Ar Condicionado"
)

type Aluno struct {
	Nome      string `json:"nome"`
	Matricula string `json:"matricula"`
	Email     string `json:"email"`
}

type Sala struct {
	ID         string        `json:"id"`
	Nome       string        `json:"nome"`
	Capacidade int           `json:"capacidade"`
	Recursos   []TypeRecurso `json:"recursos"`
	Ativa      bool          `json:"ativa"`
}

type Alocacao struct {
	SalaID     string `json:"sala_id"`
	DiaSemana  string `json:"dia_semana"`
	HoraInicio string `json:"hora_inicio"`
	HoraFim    string `json:"hora_fim"`
}

type Turma struct {
	ID         string    `json:"id"`
	Nome       string    `json:"nome"`
	Disciplina string    `json:"disciplina"`
	Professor  string    `json:"professor"`
	QntdAlunos int       `json:"qntdAlunos"`
	Alunos     []Aluno   `json:"alunos"`
	Alocacao   *Alocacao `json:"alocacao,omitempty"`
	Ativa      bool      `json:"ativa"`
}

type MatriculaRequest struct {
	AlunoID string `json:"aluno_id"`
}
