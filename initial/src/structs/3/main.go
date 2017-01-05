package main

import (
	"fmt"
	"encoding/json"
)

type Car struct {
	Name  string `json:"name"`
	Year  int `json:"year"`
	Color string `json:"color"`
	private string
	NotShowJson string `json:"-"`
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}


func main() {
	car := Car{"Gol", 2017, "Yellow", "Privado", "NotShow"}

	result, _ := json.Marshal(car)
	fmt.Println(result)
	fmt.Println(string(result))

	var carResult Car
	data := []byte(`{"Name":"Gol", "Year": 2017, "Color": "Black"}`)
	json.Unmarshal(data, &carResult)
	fmt.Println(carResult.Name, carResult.Year, carResult.Color)


	stringJsonCar := `{"name": "HB20", "year": 2015, "color": "black"}`
	resCar := Car{}
	json.Unmarshal([]byte(stringJsonCar), &resCar)
	fmt.Println(resCar)

	strFruit := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(strFruit), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}