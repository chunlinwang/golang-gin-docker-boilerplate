package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chunlinwang/golang-gin-docker-boilerplate/config"
)

type Cat struct {
	Breeds []string `json:"breeds"`
	Id     string   `json:"id"`
	Url    string   `json:"url"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
}

type CatQueryOptions struct {
	Limit int    `default:"5" form:"limit"`
	Page  int    `default:"1" form:"page"`
	Order string `default:"DESC" form:"order"`
}

func GetCats(catsOptins CatQueryOptions) ([]Cat, error) {
	client := &http.Client{}

	catApiKey := config.GetConfig("CAT_API_KEY")
	catApiBaseUrl := config.GetConfig("CAT_API_BASE_URI")
	catApiVersion := config.GetConfig("CAT_API_VERSION")

	fmt.Println(catsOptins)

	query := fmt.Sprintf("limit=%d&page=%d&order=%s", catsOptins.Limit, catsOptins.Page, catsOptins.Order)

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s%s?%s", catApiBaseUrl, catApiVersion, "/images/search", query), nil)

	req.Header.Add("x-api-key", catApiKey)

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var cats []Cat

	json.Unmarshal([]byte(body), &cats)

	fmt.Println("go routines")

	fmt.Println(cats)

	return cats, err
}
