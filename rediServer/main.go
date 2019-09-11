package main

import (
	"github.com/King-tu/redicache/rediServer/server"
	"log"
)

func init() {
	log.SetPrefix("[ErrorMsg] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	RediServer := server.NewServer()
	RediServer.Server()

}
