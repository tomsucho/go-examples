package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type data struct {
	S string `json:"blah"`
	V int    `json:"num"`
}

func main() {

	if strB, err := json.Marshal([]int{1, 2, 3}); err == nil {
		fmt.Println(strB)
	} else {
		fmt.Println(err)
	}

	if str, err := json.Marshal(data{S: "test", V: 1}); err == nil {
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(data{S: "test", V: 1})
		fmt.Println(string(str))
	} else {
		fmt.Println(err)
	}

	d := []byte(`{"name": "John", "numbers": [{"age": 42}, {"height": 182}]}`)
	var d_str map[string]interface{}
	if err := json.Unmarshal(d, &d_str); err == nil {
		fmt.Println(d_str)
	} else {
		fmt.Println(err)
	}
}
