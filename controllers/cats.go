package controllers

import (
	"fmt"
	"net/http"

	"github.com/chunlinwang/golang-gin-docker-boilerplate/services"
	"github.com/gin-gonic/gin"
)

type CatController struct{}

type CatQueryOptions struct {
	Limit int    `default:"5" form:"limit"`
	Page  int    `default:"1" form:"page"`
	order string `default:"DESC" form:"order"`
}

func (this *CatController) Cats(c *gin.Context) {
	var catsOptins services.CatQueryOptions
	c.BindQuery(&catsOptins)

	cats, _ := services.GetCats(catsOptins)

	fmt.Println("controller")

	fmt.Println(cats)

	c.JSON(http.StatusOK, cats)
}
