package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const USAGE =
`USAGE: <command> <key> [option]
	set key value
	get key
	del key [key...]

	sadd key member [member...]
	srem key member [member...]
	sismember key member 
	sinter key [key ...]
	sunion key [key ...]
	keys pattern
`

const (
	IP   = "127.0.0.1"
	Port = "6399"
)

func RediGo() {

	address := fmt.Sprint(IP +":"+ Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("成功连接至服务器", conn.RemoteAddr().String())
	fmt.Println(USAGE)

	inputReader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print(IP + ":" + Port + "> ")
		inputInfo, err := inputReader.ReadString('\n')
		if err != nil {
			log.Println("inputReader.ReadString err:", err)
			continue
		}
		//测试
		//fmt.Printf("inputInfo = %q\n", inputInfo)
		if inputInfo == "\r\n" || inputInfo == "\n" {
			continue
		}

		//Write to server
		cnt, err := conn.Write([]byte(inputInfo))
		if err != nil {
			log.Fatalln("net.Write err: ", err)
		}

		//Read from server
		buf := make([]byte, 512)
		cnt, err = conn.Read(buf)
		if err != nil {
			log.Fatalln("net.Read err: ", err)
		}
		fmt.Println(string(buf[:cnt]))
	}
}
