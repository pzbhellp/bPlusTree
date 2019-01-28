package Node

import (
	"encoding/json"
	"fmt"
	"github.com/pangzhanbo/bPlusTree/Data"
)

//基础的节点信息

type Node struct {
	MaxSize int
	Size    int
	Data    []*Data.Data
	Next    *Node
}

func (n *Node) Insert(data *Data.Data) error {
	if err := n.optimisticInsert(data); nil != err {
		return n.pessimisticInsert(data)
	}
	return nil
}

func (n *Node) optimisticInsert(data *Data.Data) error {
	if !n.canInsert() {
		return fmt.Errorf("Can't insert data. Node:%v ", n)
	}
	n.Data = Data.Sort(append(n.Data, data))
	n.Size++
	return nil
}

func (n *Node) pessimisticInsert(data *Data.Data) error {

	return nil
}

func (n *Node) Split() []*Node {
	return nil
}

func (n *Node) canInsert() bool {
	return n.Size < n.MaxSize
}

func (n *Node) Sort() {
	Data.Sort(n.Data)
}

func (n *Node) ToJson() string {
	ret, err := json.Marshal(n)
	if nil != err {
		return ""
	}
	return string(ret)
}
