package main

import (
	"ShortUrl/app"
	"log"
	"net/http"
)

func main() {
	go func() {
		mux:=http.NewServeMux()
		mux.HandleFunc("/",app.LongToShort)
		if err:=http.ListenAndServe(":8081",mux);err!=nil {
			log.Fatalf(err.Error())
		}

	}()

	go func() {
		mux:=http.NewServeMux()
		mux.HandleFunc("/",app.RedirectLongUrl)
		if err:=http.ListenAndServe(":8082",mux);err!=nil {
			log.Fatalf(err.Error())
		}
	}()
	ch:=make(chan int)
	<-ch
}