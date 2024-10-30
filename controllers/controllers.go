package controllers

import (
	"gin-api/database"
	"gin-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
	
)

func Index(ctx *gin.Context){
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	ctx.HTML(200,"index.html",gin.H{
		"alunos":alunos,
	})
}

func ListAll(ctx *gin.Context){
	var aluno []models.Aluno
	database.DB.Find(&aluno)
	ctx.JSON(200,aluno)
}
func ListById(ctx *gin.Context){
	var aluno models.Aluno
	id := ctx.Params.ByName("id")
	database.DB.First(&aluno,id)
	
	if aluno.ID == 0{
		ctx.JSON(404,gin.H{
			"NOT FOUND!":"Aluno Não Encontrado",
		})
		return
	}
	ctx.JSON(200,aluno)

}
func DeleteAluno(ctx * gin.Context){
	id := ctx.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno,id)
	if aluno.ID == 0{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"Not Found":"Aluno Inexistente",
		})
	    return
	}

	database.DB.Delete(&aluno,id)
	ctx.JSON(200,gin.H{
		"aluno":&aluno,
		"message":"Deletado Com Suceso",
	})
}
func NewAluno(ctx *gin.Context){
	var aluno models.Aluno
	if err:= ctx.ShouldBindJSON(&aluno); err !=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
		return 
	}
	if err:= models.Validator(&aluno); err !=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
		return 
	}
	database.DB.Create(&aluno)
	ctx.JSON(200,aluno)
}

func EditAluno(ctx *gin.Context){
	var aluno models.Aluno
    id := ctx.Params.ByName("id")
	database.DB.First(&aluno,id)

	if err := ctx.ShouldBindJSON(&aluno); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
	    return
	} 
	database.DB.Model(&aluno).UpdateColumns(aluno)
	ctx.JSON(200,gin.H{
		"aluno":aluno,
		"status":"Editado Com Sucesso",
	})

}
func FindByCpf(ctx *gin.Context){
	var aluno models.Aluno
	cpf := ctx.Params.ByName("cpf")

	database.DB.Where(models.Aluno{CPF:cpf}).First(&aluno)

	if aluno.ID == 0{
		ctx.JSON(404,gin.H{
			"NOT FOUND!":"Aluno Não Encontrado",
		})
		return
	}
	ctx.JSON(200,aluno)
}

func NotFound(ctx *gin.Context){
	ctx.HTML(404,"404.html",nil)
}