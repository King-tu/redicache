package main

import (
	"github.com/King-tu/redicache/rediClient/redicli"
	"log"
)

func init() {
	log.SetPrefix("[ErrorMsg] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main()  {
	redicli.RediGo()
}