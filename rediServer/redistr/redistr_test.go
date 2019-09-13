package redistr

import (
	"fmt"
	"testing"
)

var rediStrObj *RediStrObj
var key = "k1"
var val = "v1"

func TestNewRediStrObj(t *testing.T) {
	rediStrObj = NewRediStrObj()
	fmt.Println("rediStrObj = ", rediStrObj)
}


func TestRediStrObj_Set(t *testing.T) {
	ret := rediStrObj.Set(key, val)
	fmt.Println("rediStrObj.Set: ", ret)
}

func TestRediStrObj_Get(t *testing.T) {
	ret := rediStrObj.Get(key)
	fmt.Println("rediStrObj.Get: ", ret)
}
