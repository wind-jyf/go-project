package main

import (
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type Person struct {
    Name string `json:"name"`
    Age string `json:"age"`
}

type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data Person `json:"data"`
}

func cors(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")  // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
        w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
        w.Header().Add("Access-Control-Allow-Credentials", "true") //设置为true，允许ajax异步请求带cookie信息
        w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //允许请求方法
        w.Header().Set("content-type", "application/json;charset=UTF-8")             //返回数据格式是json
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        f(w, r)
    }
}

func main() {
	http.HandleFunc("/", cors(ExampleHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("method:", r.Method) // method: POST
// 	param := &struct {
// 		Username string `json:"Username"`
//     	age string `json:"age"`
// 	}{}
   
//    // 通过json解析器解析参数
//    json.NewDecoder(r.Body).Decode(param)
//    fmt.Println(fmt.Sprintf("%#v", param)) 
//    fmt.Println("Username:", param.age)
// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(r.Body)

	// 用于存放参数key=value数据
	var params map[string]string

	// 解析参数 存入map
	decoder.Decode(&params)
   	res := Response{
		0,
		"success",
		Person{
			params["Username"],
			params["age"],
		},
	}
	json.NewEncoder(w).Encode(res)
}
