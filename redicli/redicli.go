package redicli

import (
	"bufio"
	"fmt"
	"os"
	"github.com/King-tu/redicache/redipersist"
	"strings"
)

const USAGE = `USAGE:
	set key value
	get key
	del key [key...]
集合：
	sadd key member [member...]
	srem key member [member...]
	sismember key member 
	SINTER key [key ...]
	SUNION key [key ...]
	KEYS pattern
`

const (
	IP   = "192.168.5.131"
	Port = "6379"
)

func RediGo(rc *redipersist.RediCache) {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print(IP + ":" + Port + "> ")

	var ret string

	for {
		inputInfo, err := inputReader.ReadString('\n')
		if err != nil {
			//panic(err)
			fmt.Println("inputReader.ReadString err:", err)
			return
		}
		//按空格切割
		strs := strings.Fields(inputInfo)
		length := len(strs)
		if length < 1 {
			//fmt.Println("参数错误, 请重新输入:")
			//fmt.Println(USAGE)
			fmt.Print(IP + ":" + Port + "> ")
			continue
		}
		//匹配strs[0]命令，不区分大小写
		switch strings.ToLower(strs[0]) {
		case "set":
			if length != 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'set' command")
				fmt.Print(IP + ":" + Port + "> ")
				//fmt.Println(USAGE)
				continue
			}
			ret = rc.RediStrObj.Set(strs[1], strs[2])

		case "get":
			if length != 2 {
				fmt.Println("(error) ERR wrong number of arguments for 'get' command")
				fmt.Print(IP + ":" + Port + "> ")
				//fmt.Println(USAGE)
				continue
			}
			ret = rc.RediStrObj.Get(strs[1])

		case "del":
			if length < 2 {
				fmt.Println("(error) ERR wrong number of arguments for 'del' command")
				fmt.Print(IP + ":" + Port + "> ")
				//fmt.Println(USAGE)
				continue
			}
			ret = rc.Del(strs[1:])

		case "sadd":
			if length < 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'sadd' command")
				fmt.Print(IP + ":" + Port + "> ")
				continue
			}
			ret = rc.RediSetObj.Sadd(strs[1], strs[2:])

		case "srem":
			if length < 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'srem' command")
				fmt.Print(IP + ":" + Port + "> ")
				continue
			}
			ret = rc.RediSetObj.Srem(strs[1], strs[2:])

		case "smembers":
			if length != 2 {
				fmt.Println("(error) ERR wrong number of arguments for 'smembers' command")
				fmt.Print(IP + ":" + Port + "> ")
				continue
			}
			ret = rc.RediSetObj.Smembers(strs[1])

		case "sismember":
			if length != 3 {
				fmt.Println("(error) ERR wrong number of arguments for 'sismember' command")
				fmt.Print(IP + ":" + Port + "> ")
				continue
			}
			ret = rc.RediSetObj.Sismember(strs[1], strs[2])

		case "sinter":
			if length < 2 {
				fmt.Println("(error) ERR wrong number of arguments for 'sinter' command")
				fmt.Print(IP + ":" + Port + "> ")
				continue
			}
			ret = rc.RediSetObj.Sinter(strs[1:])

		case "sunion":
			if length < 2 {
				fmt.Println("(error) ERR wrong number of arguments for 'sunion' command")
				fmt.Print(IP + ":" + Port + "> ")
				continue
			}
			ret = rc.RediSetObj.Sunion(strs[1:])

		case "save":
			ret = rc.SaveToFile()

		case "exit", "quit":
			rc.SaveToFile()
			return

		case "keys":
			if length != 2 {
				fmt.Println("(error) ERR wrong number of arguments for 'keys' command")
				fmt.Print(IP + ":" + Port + "> ")
				continue
			}
			ret = rc.GetKeys(strs[1])

		default:
			//fmt.Println("参数错误, 请重新输入:")
			//fmt.Println(USAGE)
			fmt.Printf("(error) ERR unknown command '%s\n'", strs[0])
			fmt.Print(IP + ":" + Port + "> ")
			continue
		}

		fmt.Println(ret)
		fmt.Print(IP + ":" + Port + "> ")
	}
}
