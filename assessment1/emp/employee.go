package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Age    int    `json:"Age"`
	Salary int    `json:"Salary"`
}

func main() {
	e := Employee{1, "xyz", 20, 250000}
	emp, err := json.Marshal(e)
	if err != nil {
		return
	}
	fmt.Println(string(emp))

}
