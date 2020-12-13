package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func startHTTPServer() {
	http.HandleFunc("/", welcome)
	//http.HandleFunc("/jsontest", jsontest)
	err := http.ListenAndServe("0.0.0.0:8088", nil)
	if err != nil {
		fmt.Printf("something error, %s\n", err.Error())
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("something error! %s\n", err.Error())
	} else {
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			fmt.Printf("something error! %s\n", err.Error())
		} else {
			user.Age += 100
			fmt.Println(user)
			ret, _ := json.Marshal(user)
			fmt.Fprintln(w, string(ret))
		}
	}
}
