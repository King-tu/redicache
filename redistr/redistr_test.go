package redistr

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	fmt.Println("请输入：")
	for {
		_, err := fmt.Scan(&cmd, &key, &value)
		if err != nil {
			//panic(err)
			fmt.Println("fmt.Scan err:", err)
			return
		}

		fmt.Println("cmd = ", cmd)
		fmt.Println("cmd = ", key)
		fmt.Println("cmd = ", value)
	}
}
