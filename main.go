package main

import (
	"log"
	"github.com/King-tu/redicache/redicli"
	"github.com/King-tu/redicache/redipersist"
	"time"
)

func init() {
	log.SetPrefix("[ErrorMsg] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	rc := redipersist.NewRediCache()

		go func() {
		ticker := time.NewTicker(time.Second * 1)
		for {
			<- ticker.C
			rc.SaveToFile()
		}
	}()

	redicli.RediGo(rc)
}
