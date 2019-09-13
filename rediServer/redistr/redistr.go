//redistr包实现键值对的Get、Set方法
package redistr

import (
	"fmt"
	"sync"
)

//定义一个读写锁，保障线程安全
var RediStrLock sync.RWMutex

//定义一个操作 键值对存储的对象
type RediStrObj struct {
	RediStr map[string]string
}

//创建操作键值对存储的对象
func NewRediStrObj() *RediStrObj {
	return &RediStrObj{
		RediStr: make(map[string]string),
	}
}

//将键key设定为指定的“字符串”值。
//
//如果 key 已经保存了一个值，那么这个操作会直接覆盖原来的值，并且忽略原始类型。
//当set命令执行成功之后，之前设置的过期时间都将失效
//返回值
//simple-string-reply:如果SET命令正常执行那么回返回OK，否则如果加了NX 或者 XX选项，但是没有设置条件。那么会返回nil。
func (strObj *RediStrObj) Set(key, value string) string {
	RediStrLock.Lock()
	defer RediStrLock.Unlock()

	strObj.RediStr[key] = value
	return "OK"
}

//返回key的value。
//
// 如果key不存在，返回特殊值nil。如果key的value不是string，就返回错误，因为GET只处理string类型的values。
//返回值
//simple-string-reply:key对应的value，或者nil（key不存在时）
func (strObj *RediStrObj) Get(key string) string {
	RediStrLock.RLock()
	defer RediStrLock.RUnlock()

	if value, ok := strObj.RediStr[key]; ok {
		return fmt.Sprintf("%q", value)
	}

	return "(nil)"
}


