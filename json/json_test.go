package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Person struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Hobbies   []string  `json:"hobbies"`
	Addresses []Address `json:"addresses"`
}

func TestJson(t *testing.T) {
	bytes, err := json.Marshal(Person{
		Name:    "John",
		Age:     20,
		Hobbies: []string{"Coding", "Football"},
		Addresses: []Address{
			{City: "Jakarta", Street: "Jl. Kebon Jeruk"},
		},
	})

	if err != nil {
		panic(err)
	}

	t.Log(string(bytes))
	assert.JSONEq(t, `{"name":"John","age":20,"hobbies":["Coding","Football"],"addresses":[{"city":"Jakarta","street":"Jl. Kebon Jeruk"}]}`, string(bytes))

}

func TestDecode(t *testing.T) {
	jsonStr := `{"name":"John","age":20}`
	person := Person{}

	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		panic(err)
	}

	t.Log(person)
	assert.Equal(t, "John", person.Name)
	assert.Equal(t, 20, person.Age)
	assert.Nil(t, person.Hobbies)
	assert.Nil(t, person.Addresses)
}

func TestMap(t *testing.T) {
	bytes, _ := json.Marshal(map[string]any{"name": "John", "age": "20"})

	assert.JSONEq(t, `{"age":"20","name":"John"}`, string(bytes))

	jsonMap := `{"name":"John","age":"20"}`
	jsonBytes := []byte(jsonMap)

	var res map[string]any
	err := json.Unmarshal(jsonBytes, &res)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "John", res["name"])
	assert.Equal(t, "20", res["age"])
}

func TestDecoder(t *testing.T) {
	reader, _ := os.Open("sample.json")

	decoder := json.NewDecoder(reader)

	var res map[string]any

	decoder.Decode(&res)

	t.Log(res)
}

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("result.json")

	encoder := json.NewEncoder(writer)

	encoder.Encode(map[string]any{"name": "John", "age": "20"})
}
