package redipersist

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/King-tu/redicache/rediServer/rediset"
	"github.com/King-tu/redicache/rediServer/redistr"
	"regexp"
)

const (
	REDI_FLIE_NAME = "dump.rdb"
)

type RediCache struct {
	RediStrObj *redistr.RediStrObj
	RediSetObj *rediset.RediSetObj
}

//创建一个RediCache对象
func NewRediCache() *RediCache {
	rc := &RediCache{
		RediStrObj: redistr.NewRediStrObj(),
		RediSetObj: rediset.NewRediSetObj(),
	}
	rc.LoadFromFile()
	return rc
}

func (rc *RediCache) GetKeys(pattern string) (keys string) {
	if pattern == "*" {
		//TODO
		//fmt.Println("pattern = ", pattern)
		//“.” 匹配任意一个 字符。
		pattern = "."
	}
	regExp := regexp.MustCompile(pattern)

	for key, _ := range rc.RediStrObj.RediStr {
		if regExp.MatchString(key) {
			keys += fmt.Sprintf("%q\n", key)
		}
	}

	for key, _ := range rc.RediSetObj.RediSet {
		if regExp.MatchString(key) {
			keys += fmt.Sprintf("%q\n", key)
		}
	}

	if len(keys) == 0 {
		return "(empty list or set)"
	}

	return
}

//删除指定的一批keys，如果删除中的某些key不存在，则直接忽略。
//返回值
//integer-reply： 被删除的keys的数量
func (rc *RediCache) Del(keys []string) string {

	if len(keys) == 0 {
		return "(error) ERR wrong number of arguments for 'del' command"
	}

	var count int
	for _, key := range keys {
		if _, ok := rc.RediStrObj.RediStr[key]; ok {
			//加锁
			redistr.RediStrLock.Lock()
			delete(rc.RediStrObj.RediStr, key)
			//解锁
			redistr.RediStrLock.Unlock()
			count++
		}

		if _, ok := rc.RediSetObj.RediSet[key]; ok {
			//加锁
			rediset.RediSetLock.Lock()
			delete(rc.RediStrObj.RediStr, key)
			//解锁
			rediset.RediSetLock.Unlock()
			count++
		}
	}
	return fmt.Sprintf("(integer) %d", count)
}


//保存缓存数据至磁盘文件
func (rc *RediCache) SaveToFile() string {
	////测试代码
	//fmt.Println("SaveToFile: ", rc.RediStrObj, rc.RediSetObj)
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(rc)
	if err != nil {
		log.Println("encoder.Encode err: ", err)
		return ""
	}

	err = ioutil.WriteFile(REDI_FLIE_NAME, buf.Bytes(), 0600)
	if err != nil {
		log.Println("ioutil.WriteFile err: ", err)
		return ""
	}

	return "OK"
}

//从磁盘文件加载数据
func (rc *RediCache) LoadFromFile() {
	if IsFileNotExist(REDI_FLIE_NAME) {
		log.Println(REDI_FLIE_NAME, " 文件不存在")
		return
	}
	rdbBuf, err := ioutil.ReadFile(REDI_FLIE_NAME)
	if err != nil {
		log.Println("ioutil.ReadFile err: ", err)
		return
	}

	decoder := gob.NewDecoder(bytes.NewReader(rdbBuf))
	err = decoder.Decode(rc)
	if err != nil {
		log.Println("decoder.Decode err: ", err)
		return
	}
	//测试代码
	//fmt.Println("LoadFromFile: ", rc.RediStrObj, rc.RediSetObj)
}

//判断文件是否存在
func IsFileNotExist(fileName string) bool {
	_, err := os.Stat(fileName)
	//测试代码
	//fmt.Println(os.IsExist(err))	//存在:false	//不存在：false
	//fmt.Println(os.IsNotExist(err))	//存在:false	//不存在：true
	return os.IsNotExist(err)
}
