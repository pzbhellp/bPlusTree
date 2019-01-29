package Node

import (
	"github.com/pangzhanbo/bPlusTree/Data"
	"testing"
)

func TestNode_Insert(t *testing.T) {
	node := NewNode()
	for i := 0; i < 17; i++ {
		data := Data.GenData()
		temp, _ := node.Insert(data)
		if nil != temp {
			println("split node:")
			println(temp.ToJson())
		}
	}
	println(node.ToJson())
}
