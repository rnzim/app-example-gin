package routes

import (
	"gin-api/controllers"
	
	"github.com/gin-gonic/gin"
)

func HandleRequest()  {
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")
	app.Static("/assets","./assets")
	app.GET("/index",controllers.Index)
	app.GET("/",controllers.ListAll)
	app.GET("/:id",controllers.ListById)
	app.DELETE("/deleta/:id",controllers.DeleteAluno)
	app.POST("/aluno",controllers.NewAluno)
	app.PATCH("/aluno/:id",controllers.EditAluno)
	app.NoRoute(controllers.NotFound)
	app.GET("/aluno/cpf/:cpf",controllers.FindByCpf)
	app.Run(":30001")
}