package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Demo struct {
	Age      int    `json:"Age"`
	Name     string `json:"Name"`
	Verified bool   `json:"Verified"`
}

func main() {
	var data Demo
	//var obj map[string]interface{}
	arg := os.Args[1]
	jsonCon := fmt.Sprintf("`%s`", arg)
	fmt.Println(jsonCon)
	err := json.Unmarshal([]byte(arg), &data)
	fmt.Println(err)

	fmt.Println("Age : ", data.Age)
	fmt.Println("Name : ", data.Name)
	fmt.Println("Verified : ", data.Verified)
}
