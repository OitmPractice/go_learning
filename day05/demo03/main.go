package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testUnMarshalStructAndMap()
}

func testStruct() {
	monster := Monster{
		Name:     "猪八戒",
		Age:      500,
		Birthday: "0001.01.01",
		Sal:      80000,
		Skill:    "吃睡",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testMap() {
	// interface{} 这样写代表值可以是任意类型
	var params map[string]interface{}
	params = make(map[string]interface{})
	params["name"] = "oitm"
	params["age"] = 18
	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	slice = make([]map[string]interface{}, 0)

	for i := 0; i < 2; i++ {
		var param1 = make(map[string]interface{})
		param1["name"] = fmt.Sprintf("test%d", i)
		param1["age"] = fmt.Sprintf("%d", i+10)
		slice = append(slice, param1)
	}

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func testUnMarshalStructAndMap() {
	//反序列化字符串为结构体或者map
	str := "{\"name\":\"猪八戒\",\"age\":500,\"birthday\":\"0001.01.01\",\"sal\":80000,\"skill\":\"吃睡\"}"
	monster := Monster{}
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("monstaer == %v \n", monster)

	var params map[string]interface{} = make(map[string]interface{})
	err = json.Unmarshal([]byte(str), &params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("params == %v \n", params["name"])

}
