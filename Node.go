package bPlusTree

import (
	"encoding/json"
	"fmt"
)

//基础的节点信息

type Node struct {
	MaxSize int
	Size int
	Data []*Data
}

func (n *Node) Insert (data *Data) error {
	if !n.canInsert() {
		return fmt.Errorf("Can't insert.Node:[%#v] ",n)
	}
	n.Data = append(n.Data,data)
	n.Sort()
	return nil
}

func (n *Node) canInsert () bool {
	return n.Size < n.MaxSize
}

func (n *Node) Sort () {
	Sort(n.Data)
}


func (n *Node) ToJson () string {
	ret,err := json.Marshal(n)
	if nil != err {
		return ""
	}
	return string(ret)
}
