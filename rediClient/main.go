package main

import "log"

func init() {
	log.SetPrefix("[ErrorMsg] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main()  {
	RediGo()
}