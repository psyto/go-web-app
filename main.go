package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	myJson := `[
		{
			"first_name": "Hiroyuki",
			"last_name": "Saito",
			"age": 53
		},
		{
			"first_name": "Amane",
			"last_name": "Saito",
			"age": 5
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Mirai"
	m1.LastName = "Saito"
	m1.Age = 11

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Mayumi"
	m2.LastName = "Saito"
	m2.Age = 44

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")

	if err != nil {
		log.Println("Error marshalling json", err)
	}

	fmt.Println(string(newJson))
}
