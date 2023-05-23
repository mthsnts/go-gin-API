package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mthsnts/go-gin-API/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.POST("/alunos/new", controllers.CriaNovoAluno)
	r.DELETE("/alunos/delete/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/edit/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RouteNotFound)
	r.Run()
}
