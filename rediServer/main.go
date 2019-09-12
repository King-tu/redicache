package main

import (
	"github.com/King-tu/redicache/rediServer/server"
	"io"
	"log"
	"os"
)

func init() {
	errFile,err:=os.OpenFile("errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开日志文件失败：", err)
	}

	//设置 前缀
	log.SetPrefix("[ErrorMsg] ")
	//设置日志输出选项： 日期 时间 文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//输出到文件
	//log.SetOutput(errFile)
	//输出到文件 和 标准错误输出(屏幕)
	log.SetOutput(io.MultiWriter(os.Stderr,errFile))
}

func main() {

	RediServer := server.NewServer()
	RediServer.Server()

}
