package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{
 "Name": "oscar",
 "Age": 18,
 "Verified": "false"
}`
	var obj map[string]interface{}
	json.Unmarshal([]byte(jsonData), &obj)
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))
}
