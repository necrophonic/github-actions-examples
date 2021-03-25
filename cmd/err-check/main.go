package main

import "encoding/json"

func main() {
	j := []byte(`{"cat":"dog"}`)

	var animal map[string]interface{}
	json.Unmarshal(j, &animal)
}
