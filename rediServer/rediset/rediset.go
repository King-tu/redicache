//rediset包提供一些操作集合的方法
package rediset

import (
	"fmt"
	"sync"
)

//定义读写锁
var RediSetLock sync.RWMutex

//定义集合的操作对象
type RediSetObj struct {
	RediSet map[string]map[string]struct{}
}

//创建集合的操作对象
func NewRediSetObj() *RediSetObj {
	return &RediSetObj{
		RediSet: make(map[string]map[string]struct{}),
	}
}

//添加一个或多个指定的member元素到集合的 key中.
//
// 指定的一个或者多个元素member 如果已经在集合key中存在则忽略.如果集合key 不存在，则新建集合key,并添加member元素到集合key中.
//如果key 的类型不是集合则返回错误.
//返回值
//integer-reply:返回新成功添加到集合里元素的数量，不包括已经存在于集合中的元素.
func (rsObj *RediSetObj) Sadd(key string, members []string) string {
	if len(members) == 0 {
		return "(error) ERR wrong number of arguments for 'sadd' command"
	}

	RediSetLock.Lock()
	defer RediSetLock.Unlock()

	memMap := make(map[string]struct{})
	var count int
	for _, member := range members {
		memMap[member] = struct{}{}
		count++
	}
	rsObj.RediSet[key] = memMap

	return fmt.Sprintf("(integer) %d", count)
}

//在key集合中移除指定的元素.
//
// 如果指定的元素不是key集合中的元素则忽略 如果key集合不存在则被视为一个空的集合，该命令返回0.
//如果key的类型不是一个集合,则返回错误.
//返回值
//integer-reply:从集合中移除元素的个数，不包括不存在的成员.
func (rsObj *RediSetObj) Srem(key string, members []string) string {
	var count int

	RediSetLock.Lock()
	defer RediSetLock.Unlock()

	if memMap, ok := rsObj.RediSet[key]; ok {
		for _, member := range members {
			delete(memMap, member)
			count++
		}
	}
	return fmt.Sprintf("(integer) %d", count)
}

//返回key集合所有的元素.
//
//该命令的作用与使用一个参数的SINTER 命令作用相同.
//返回值
//array-reply:集合中的所有元素.
func (rsObj *RediSetObj) Smembers(key string) (members string) {
	RediSetLock.RLock()
	defer RediSetLock.RUnlock()

	return PrintSet(rsObj.RediSet[key])
}


//返回成员 member 是否是存储的集合 key的成员.
//
//返回值
//integer-reply,详细说明:
//如果member元素是集合key的成员，则返回1
//如果member元素不是key的成员，或者集合key不存在，则返回0
func (rsObj *RediSetObj) Sismember(key, member string) string {
	var count int

	RediSetLock.RLock()
	defer RediSetLock.RUnlock()

	if memMap, ok := rsObj.RediSet[key]; ok {
		if _, ok := memMap[member]; ok {
			count++
		}
	}
	return fmt.Sprintf("(integer) %d", count)
}


//返回指定所有的集合的成员的交集.
//
//如果key不存在则被认为是一个空的集合,当给定的集合为空的时候,结果也为空.(一个集合为空，结果一直为空).
//返回值
//array-reply: 结果集成员的列表.
func (rsObj *RediSetObj) Sinter(keys []string) (members string) {
	if _, ok := rsObj.RediSet[keys[0]]; !ok {
		members = "(empty list or set)"
		return
	}

	RediSetLock.RLock()
	defer RediSetLock.RUnlock()

	interSet := rsObj.RediSet[keys[0]]
	for _, key := range keys[1:] {
		tmpSet := make(map[string]struct{})
		if _, ok := rsObj.RediSet[key]; !ok {
			members = "(empty list or set)"
			return
		}

		if len(interSet) < len(rsObj.RediSet[key]) {
			for mem, _ := range interSet {
				if _, ok := rsObj.RediSet[key][mem]; ok {
					tmpSet[mem] = struct{}{}
				}
			}
		} else {
			for mem, _ := range rsObj.RediSet[key] {
				if _, ok := interSet[mem]; ok {
					tmpSet[mem] = struct{}{}
				}
			}
		}
		interSet = tmpSet
	}
	return PrintSet(interSet)
}

//返回给定的多个集合的并集中的所有成员
//
//不存在的key可以认为是空的集合
//返回值
//array-reply:并集的成员列表
func (rsObj *RediSetObj) Sunion(keys []string) (members string) {
	unionSet := make(map[string]struct{})

	RediSetLock.RLock()
	defer RediSetLock.RUnlock()

	for _, key := range keys {
		if memMap, ok := rsObj.RediSet[key]; ok {
			for mem, _ := range memMap {
				unionSet[mem] = struct{}{}
			}
		}
	}
	return PrintSet(unionSet)
}

//打印集合元素
func PrintSet(set map[string]struct{}) (members string) {

	if len(set) == 0 {
		members = "(empty list or set)"
	}

	for member, _ := range set {
		members += fmt.Sprintf("%q\n", member)
	}
	return
}
