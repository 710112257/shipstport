package main

import (
	"net/http"
)


func handle(res http.ResponseWriter,req *http.Request){
	if req.URL.String()=="/"{
		res.Write([]byte("hello"))
		return
	}
	res.Write([]byte("bye"))
}

func main(){
	http.HandleFunc("/",handle)
	http.ListenAndServe(":8080",nil)
}