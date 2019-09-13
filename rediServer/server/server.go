//server包搭建了一个简单的tcp服务器，监听客户端连接
package server

import (
	"fmt"
	"github.com/King-tu/redicache/rediServer/conf"
	"github.com/King-tu/redicache/rediServer/redipersist"
	"log"
	"net"
	"strings"
	"time"
)

//定义服务器操作对象
type Server struct {
	RC *redipersist.RediCache
}

//创建服务器操作对象
func NewServer() *Server {
	return &Server{
		RC: redipersist.NewRediCache(),
	}
}

//启动服务器，监听客户端连接
func (s *Server) Server() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", conf.ADDRESS)
	if err != nil {
		log.Println(err)
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("等待客户端连接: ")

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Println("tcpListener.AcceptTCP err: ", err)
			continue
		}

		go func() {
			defer tcpConn.Close()
			remoteAddr := tcpConn.RemoteAddr().String()
			fmt.Printf("客户端 %s 已连接...\n", remoteAddr)

			for {
				buf := make([]byte, 512)
				cnt, err := tcpConn.Read(buf)
				if err != nil {
					log.Printf("客户端 %s 已断开连接。\n", remoteAddr)
					return
				}

				//转换为字符串
				cliData := string(buf[:cnt])
				//调用业务处理函数
				retInfo := s.Handle(cliData)

				cnt, err = tcpConn.Write([]byte(retInfo))
				if err != nil {
					log.Println("tcpConn.Write err: ", err)
					continue
				}
			}
		}()
	}
}

//处理客户端命令
func (s *Server) Handle(cliData string) string {

	//按空格切割
	strs := strings.Fields(cliData)
	length := len(strs)
	if length < 1 {
		return "命令不能为空"
	}

	errMsg := fmt.Sprintf("(error) ERR wrong number of arguments for '%s' command", strs[0])
	//匹配strs[0]命令，不区分大小写
	switch strings.ToLower(strs[0]) {
	case "set":
		//校验
		if length != 3 {
			return errMsg
		}
		return s.RC.RediStrObj.Set(strs[1], strs[2])

	case "get":
		if length != 2 {
			return errMsg
		}
		return s.RC.RediStrObj.Get(strs[1])

	case "del":
		if length < 2 {
			return errMsg
		}
		return s.RC.Del(strs[1:])

	case "sadd":
		if length < 3 {
			return errMsg
		}
		return s.RC.RediSetObj.Sadd(strs[1], strs[2:])

	case "srem":
		if length < 3 {
			return errMsg
		}
		return s.RC.RediSetObj.Srem(strs[1], strs[2:])

	case "smembers":
		if length != 2 {
			return errMsg
		}
		return s.RC.RediSetObj.Smembers(strs[1])

	case "sismember":
		if length != 3 {
			return errMsg
		}
		return s.RC.RediSetObj.Sismember(strs[1], strs[2])

	case "sinter":
		if length < 2 {
			return errMsg
		}
		return s.RC.RediSetObj.Sinter(strs[1:])

	case "sunion":
		if length < 2 {
			return errMsg
		}
		return s.RC.RediSetObj.Sunion(strs[1:])

	case "save", "exit", "quit":
		return s.RC.SaveToFile()

	case "keys":
		if length != 2 {
			return errMsg
		}
		return s.RC.GetKeys(strs[1])

	default:
		return fmt.Sprintf("(error) ERR unknown command '%s'", strs[0])
	}
}

//定时保存缓存数据
func (s *Server) TimeToSave()  {
	ticker := time.NewTicker(time.Second * conf.DURATION)
	for {
		<- ticker.C
		s.RC.SaveToFile()
	}
}