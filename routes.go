package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func configurarRotas(r *gin.Engine) {

	v1 := r.Group("/api/v1")

	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now(),
			"version":   "1.0.0",
		})
	})

	v1.POST("/salas", criarSala)
	v1.GET("/salas", listarSalas)
	v1.GET("/salas/:id/agenda", consultarAgendaSala)

	v1.POST("/alunos", criarAluno)
	v1.GET("/alunos", listarAlunos)
	v1.GET("/alunos/:id", buscarAluno)

	v1.POST("/turmas", criarTurma)
	v1.GET("/turmas", listarTurmas)

	v1.POST("/turmas/:id/alunos", matricularAluno)
	v1.GET("/turmas/:id/alunos", listarAlunosTurma)

	v1.POST("/turmas/:id/alocar", alocarSala)
}
