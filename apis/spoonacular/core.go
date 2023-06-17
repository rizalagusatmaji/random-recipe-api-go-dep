package spoonacular

import (
	"io/ioutil"
	"net/http"
)

func GetRandomRecipes() (Recipe, error) {
	//
	url := "https://api.spoonacular.com/recipes/random"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return Recipe{}, err
	}
	//secret variable
	req.Header.Add("x-api-key", "ff56e33bbf144ac5967a9b61b1bb58f9")

	res, err := client.Do(req)
	if err != nil {
		return Recipe{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Recipe{}, err
	}

	randomRecipesResponse, err := UnmarshalRandomRecipesResponse(body)
	if err != nil {
		return Recipe{}, err
	}

	return randomRecipesResponse.Recipes[0], nil
}
