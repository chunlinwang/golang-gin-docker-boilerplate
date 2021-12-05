package controllers

import (
	"fmt"
	"net/http"

	"github.com/chunlinwang/golang-gin-docker-boilerplate/services"
	"github.com/gin-gonic/gin"
)

type CatController struct{}	

func (this *CatController) Cats(c *gin.Context) {
	var catsOptins services.CatQueryOptions
	c.BindQuery(&catsOptins)

	cats, _ := services.GetCats(catsOptins)

	fmt.Println("controller")

	fmt.Println(cats)

	c.JSON(http.StatusOK, cats)
}
