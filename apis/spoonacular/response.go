package spoonacular

import "encoding/json"

func UnmarshalRandomRecipesResponse(data []byte) (RandomRecipesResponse, error) {
	var r RandomRecipesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RandomRecipesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RandomRecipesResponse struct {
	Recipes []Recipe `json:"recipes"`
}

type Recipe struct {
	PricePerServing     float64              `json:"pricePerServing"`
	ExtendedIngredients []ExtendedIngredient `json:"extendedIngredients"`
	Title               string               `json:"title"`
}

type ExtendedIngredient struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Unit   string  `json:"unit"`
}

type Measures struct {
	Us     Metric `json:"us"`
	Metric Metric `json:"metric"`
}

type Metric struct {
	Amount    float64 `json:"amount"`
	UnitShort string  `json:"unitShort"`
	UnitLong  string  `json:"unitLong"`
}
