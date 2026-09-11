package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func criarSala(c *gin.Context) {

	var novaSala Sala

	if err := c.ShouldBindJSON(&novaSala); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	if novaSala.ID == "" || novaSala.Nome == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID e nome são obrigatórios",
		})
		return
	}

	if novaSala.Capacidade <= 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "A capacidade deve ser maior que zero",
		})
		return
	}

	for _, sala := range salas {

		if sala.ID == novaSala.ID {

			c.JSON(http.StatusConflict, gin.H{
				"erro": "Já existe uma sala com esse ID",
			})
			return
		}
	}

	novaSala.Ativa = true

	salas = append(salas, novaSala)

	c.JSON(http.StatusCreated, novaSala)
}

func listarSalas(c *gin.Context) {

	c.JSON(http.StatusOK, salas)
}

func consultarAgendaSala(c *gin.Context) {

	salaID := c.Param("id")

	salaExiste := false

	for _, sala := range salas {

		if sala.ID == salaID {

			salaExiste = true
			break
		}
	}

	if !salaExiste {

		c.JSON(http.StatusNotFound, gin.H{
			"erro": "Sala não encontrada",
		})
		return
	}

	agenda := []gin.H{}

	for _, turma := range turmas {

		if turma.Alocacao != nil &&
			turma.Alocacao.SalaID == salaID {

			agenda = append(agenda, gin.H{
				"turma_id":    turma.ID,
				"turma":       turma.Nome,
				"disciplina":  turma.Disciplina,
				"dia_semana":  turma.Alocacao.DiaSemana,
				"hora_inicio": turma.Alocacao.HoraInicio,
				"hora_fim":    turma.Alocacao.HoraFim,
			})
		}
	}

	c.JSON(http.StatusOK, agenda)
}
