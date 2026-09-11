package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func criarAluno(c *gin.Context) {

	var novoAluno Aluno

	if err := c.ShouldBindJSON(&novoAluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	if novoAluno.Nome == "" ||
		novoAluno.Matricula == "" ||
		novoAluno.Email == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Nome, matrícula e e-mail são obrigatórios",
		})
		return
	}

	for _, aluno := range alunos {
		if aluno.Matricula == novoAluno.Matricula {
			c.JSON(http.StatusConflict, gin.H{
				"erro": "Aluno já cadastrado",
			})
			return
		}
	}

	alunos = append(alunos, novoAluno)

	c.JSON(http.StatusCreated, novoAluno)
}

func listarAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, alunos)
}

func buscarAluno(c *gin.Context) {

	id := c.Param("id")

	for _, aluno := range alunos {
		if aluno.Matricula == id {
			c.JSON(http.StatusOK, aluno)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"erro": "Aluno não encontrado",
	})
}
