package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func criarTurma(c *gin.Context) {

	var novaTurma Turma

	if err := c.ShouldBindJSON(&novaTurma); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	if novaTurma.ID == "" ||
		novaTurma.Nome == "" ||
		novaTurma.Disciplina == "" ||
		novaTurma.Professor == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID, nome, disciplina e professor são obrigatórios",
		})
		return
	}

	for _, turma := range turmas {

		if turma.ID == novaTurma.ID {

			c.JSON(http.StatusConflict, gin.H{
				"erro": "Já existe uma turma com esse ID",
			})
			return
		}
	}

	novaTurma.QntdAlunos = 0
	novaTurma.Alunos = []Aluno{}
	novaTurma.Alocacao = nil
	novaTurma.Ativa = true

	turmas = append(turmas, novaTurma)

	c.JSON(http.StatusCreated, novaTurma)
}

func listarTurmas(c *gin.Context) {

	c.JSON(http.StatusOK, turmas)
}

func matricularAluno(c *gin.Context) {

	turmaID := c.Param("id")

	var dados MatriculaRequest

	if err := c.ShouldBindJSON(&dados); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	var alunoEncontrado Aluno

	alunoExiste := false

	for _, aluno := range alunos {

		if aluno.Matricula == dados.AlunoID {

			alunoEncontrado = aluno
			alunoExiste = true
			break
		}
	}

	if !alunoExiste {

		c.JSON(http.StatusNotFound, gin.H{
			"erro": "Aluno não encontrado",
		})
		return
	}

	for i := range turmas {

		if turmas[i].ID == turmaID {

			// Verifica se o aluno já está matriculado
			for _, aluno := range turmas[i].Alunos {

				if aluno.Matricula == dados.AlunoID {

					c.JSON(http.StatusConflict, gin.H{
						"erro": "Aluno já matriculado nessa turma",
					})
					return
				}
			}

			// Se a turma já possui sala, verifica capacidade
			if turmas[i].Alocacao != nil {

				for _, sala := range salas {

					if sala.ID == turmas[i].Alocacao.SalaID {

						if len(turmas[i].Alunos)+1 > sala.Capacidade {

							c.JSON(
								http.StatusUnprocessableEntity,
								gin.H{
									"erro": "Capacidade da sala insuficiente",
								},
							)
							return
						}
					}
				}
			}

			// Verifica conflito de horário do aluno
			if turmas[i].Alocacao != nil {

				for _, outraTurma := range turmas {

					if outraTurma.ID == turmaID ||
						outraTurma.Alocacao == nil {

						continue
					}

					alunoEstaNaOutraTurma := false

					for _, aluno := range outraTurma.Alunos {

						if aluno.Matricula == dados.AlunoID {

							alunoEstaNaOutraTurma = true
							break
						}
					}

					if !alunoEstaNaOutraTurma {

						continue
					}

					if outraTurma.Alocacao.DiaSemana ==
						turmas[i].Alocacao.DiaSemana &&
						horariosSobrepostos(
							turmas[i].Alocacao.HoraInicio,
							turmas[i].Alocacao.HoraFim,
							outraTurma.Alocacao.HoraInicio,
							outraTurma.Alocacao.HoraFim,
						) {

						c.JSON(http.StatusConflict, gin.H{
							"erro": "Aluno possui conflito de horário com outra turma",
						})
						return
					}
				}
			}

			// Adiciona aluno na turma
			turmas[i].Alunos = append(
				turmas[i].Alunos,
				alunoEncontrado,
			)

			turmas[i].QntdAlunos =
				len(turmas[i].Alunos)

			c.JSON(
				http.StatusCreated,
				turmas[i],
			)

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"erro": "Turma não encontrada",
	})
}

func listarAlunosTurma(c *gin.Context) {

	turmaID := c.Param("id")

	for _, turma := range turmas {

		if turma.ID == turmaID {

			c.JSON(
				http.StatusOK,
				turma.Alunos,
			)

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"erro": "Turma não encontrada",
	})
}

func alocarSala(c *gin.Context) {

	turmaID := c.Param("id")

	var novaAlocacao Alocacao

	if err := c.ShouldBindJSON(&novaAlocacao); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	// Verifica campos obrigatórios
	if novaAlocacao.SalaID == "" ||
		novaAlocacao.DiaSemana == "" ||
		novaAlocacao.HoraInicio == "" ||
		novaAlocacao.HoraFim == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Sala, dia da semana, horário de início e horário de término são obrigatórios",
		})
		return
	}

	// Validação dos horários
	inicio, errInicio :=
		time.Parse(
			"15:04",
			novaAlocacao.HoraInicio,
		)

	fim, errFim :=
		time.Parse(
			"15:04",
			novaAlocacao.HoraFim,
		)

	if errInicio != nil ||
		errFim != nil ||
		!inicio.Before(fim) {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Horário inválido",
		})
		return
	}

	// Procura a sala
	var salaEncontrada *Sala

	for i := range salas {

		if salas[i].ID ==
			novaAlocacao.SalaID {

			salaEncontrada = &salas[i]
			break
		}
	}

	if salaEncontrada == nil {

		c.JSON(http.StatusNotFound, gin.H{
			"erro": "Sala não encontrada",
		})
		return
	}

	// Sala precisa estar ativa
	if !salaEncontrada.Ativa {

		c.JSON(http.StatusConflict, gin.H{
			"erro": "Sala não está ativa",
		})
		return
	}

	// Procura a turma
	turmaIndex := -1

	for i := range turmas {

		if turmas[i].ID == turmaID {

			turmaIndex = i
			break
		}
	}

	if turmaIndex == -1 {

		c.JSON(http.StatusNotFound, gin.H{
			"erro": "Turma não encontrada",
		})
		return
	}

	// Turma precisa estar ativa
	if !turmas[turmaIndex].Ativa {

		c.JSON(http.StatusConflict, gin.H{
			"erro": "Turma não está ativa",
		})
		return
	}

	// Capacidade da sala
	if len(turmas[turmaIndex].Alunos) >
		salaEncontrada.Capacidade {

		c.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"erro": "Capacidade da sala insuficiente",
			},
		)

		return
	}

	// Verifica conflito de sala
	for _, turma := range turmas {

		if turma.ID == turmaID ||
			turma.Alocacao == nil {

			continue
		}

		if turma.Alocacao.SalaID ==
			novaAlocacao.SalaID &&
			turma.Alocacao.DiaSemana ==
				novaAlocacao.DiaSemana &&
			horariosSobrepostos(
				novaAlocacao.HoraInicio,
				novaAlocacao.HoraFim,
				turma.Alocacao.HoraInicio,
				turma.Alocacao.HoraFim,
			) {

			c.JSON(http.StatusConflict, gin.H{
				"erro": "Conflito de horário na sala",
			})

			return
		}
	}

	// Verifica conflito dos alunos
	for _, outraTurma := range turmas {

		if outraTurma.ID == turmaID ||
			outraTurma.Alocacao == nil {

			continue
		}

		if outraTurma.Alocacao.DiaSemana !=
			novaAlocacao.DiaSemana {

			continue
		}

		if !horariosSobrepostos(
			novaAlocacao.HoraInicio,
			novaAlocacao.HoraFim,
			outraTurma.Alocacao.HoraInicio,
			outraTurma.Alocacao.HoraFim,
		) {

			continue
		}

		for _, alunoNovaTurma := range turmas[turmaIndex].Alunos {

			for _, alunoOutraTurma := range outraTurma.Alunos {

				if alunoNovaTurma.Matricula ==
					alunoOutraTurma.Matricula {

					c.JSON(
						http.StatusConflict,
						gin.H{
							"erro": "Aluno possui conflito de horário com outra turma",
						},
					)

					return
				}
			}
		}
	}

	// Realiza a alocação
	turmas[turmaIndex].Alocacao =
		&novaAlocacao

	c.JSON(
		http.StatusOK,
		turmas[turmaIndex],
	)
}
