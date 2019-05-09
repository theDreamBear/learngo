package main

import (
	"fmt"
	"gomcpack/npc"
	"notifyproxy/core/handle"
	"time"
)

var server npc.Server

func initServer() {
	server.Addr = ":10804"
	server.Handler = new(handle.Handler)
	server.ReadTimeout = time.Second * 30
	server.WriteTimeout = time.Second * 30
	server.ErrorLog = nil

}

func main() {
	initServer()
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
