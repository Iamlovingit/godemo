package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	var user User
	req := `{"name":"junneyang", "age": 88}`
	reqnew := bytes.NewBuffer([]byte(req))
	request, _ := http.NewRequest("POST", "http://192.168.3.38:8088/welcome/", reqnew)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	fmt.Printf("response === %v\n", response)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println("something wrong.")
		}
		fmt.Println(user)
	}
}
