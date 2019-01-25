package main

import "net/http"


func wsHanlder(w http.ResponseWriter, r *Request) {
	w.Write([]byte("hello"))
	_ = r
}


func main() {
	http.HandleFunc("/ws",wsHanlder)

	http.ListenAndServe("0.0.0.0:7777",nil)
}
