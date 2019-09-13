package rediset

import (
	"fmt"
	"testing"
)

func TestNewRediSetObj(t *testing.T) {
	fmt.Println("rediSet: ", NewRediSetObj())
}

func TestRediSetObj_Sadd(t *testing.T) {
	rediSet := NewRediSetObj()
	key := "k1"
	strs := []string{"v1", "v2", "v3", "v4"}
	ret := rediSet.Sadd(key, strs)
	fmt.Println(ret)
}

func TestRediSetObj_Srem(t *testing.T) {
	rediSet := NewRediSetObj()
	key := "k1"
	strs := []string{"v1", "v2", "v3", "v4"}
	rediSet.Sadd(key, strs)
	ret := rediSet.Srem(key, strs[:2])
	fmt.Println(ret)
}

func TestRediSetObj_Sinter(t *testing.T) {
	rediSet := NewRediSetObj()
	key1 := "k1"
	strs1 := []string{"v1", "v2", "v3", "v4"}
	key2 := "k2"
	strs2 := []string{"v3", "v4", "v5", "v6"}
	rediSet.Sadd(key1, strs1)
	rediSet.Sadd(key2, strs2)
	ret := rediSet.Sinter([]string{key1, key2})

	//key3 := "k3"
	//ret := rediSet.Sinter([]string{key1, key2, key3})
	fmt.Println(ret)
}

func TestRediSetObj_Sunion(t *testing.T) {
	rediSet := NewRediSetObj()
	key1 := "k1"
	strs1 := []string{"v1", "v2", "v3", "v4"}
	key2 := "k2"
	strs2 := []string{"v3", "v4", "v5", "v6"}
	rediSet.Sadd(key1, strs1)
	rediSet.Sadd(key2, strs2)
	//ret := rediSet.Sunion([]string{key1, key2})

	key3 := "k3"
	ret := rediSet.Sunion([]string{key1, key2, key3})
	fmt.Println(ret)
}
