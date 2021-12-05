package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/chunlinwang/golang-gin-docker-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	v1 := r.Group("/v1")

	v1.Use()
	{
		testController := new(controllers.TestController)
		v1.GET("/json", testController.JsonApi)
		v1.GET("/jsonPureApi", testController.JsonPureApi)

		catController := new(controllers.CatController)
		v1.GET("/cats", catController.Cats)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome gin app!")
	})

	r.GET("/test", func(c *gin.Context) {
		client := &http.Client{}

		jsonPath := "https://raw.githubusercontent.com/chunlinwang/golang-excerces/main/example.json"
		req, _ := http.NewRequest("GET", jsonPath, nil)

		resp, _ := client.Do(req)

		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var decodeJson []map[string]interface{}

		json.Unmarshal(body, &decodeJson)

		c.PureJSON(200, decodeJson)
	})

	return r
}
