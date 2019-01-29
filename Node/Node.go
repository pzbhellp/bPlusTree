package Node

import (
	"encoding/json"
	"fmt"
	"github.com/pangzhanbo/bPlusTree/Data"
	"github.com/pangzhanbo/bPlusTree/config"
	"math"
)

//基础的节点信息

type Node interface {
	Insert(data *Data.Data) (*NodeBase, error)
	OptimisticInsert(data *Data.Data) error
	PessimisticInsert(data *Data.Data) (splitNode *NodeBase, err error)
	Split() *NodeBase
	Sort()
	ToJson() string
	Find(key *Data.Key) *Data.Data
}

type NodeBase struct {
	MaxSize  int
	Size     int
	Data     []*Data.Data
	NextNode *NodeBase
	UpNode   *NodeBase
}

func NewNode() *NodeBase {
	return &NodeBase{
		MaxSize:  config.NODE_MAX_SIZE,
		Size:     0,
		NextNode: nil,
		UpNode:   nil,
	}
}

func (n *NodeBase) Insert(data *Data.Data) (*NodeBase, error) {
	if err := n.OptimisticInsert(data); nil != err {
		return n.PessimisticInsert(data)
	}
	return nil, nil
}

func (n *NodeBase) OptimisticInsert(data *Data.Data) error {
	if !n.canInsert() {
		return fmt.Errorf("Can't insert data. Node:%v ", n)
	}
	n.Data = Data.Sort(append(n.Data, data))
	n.Size++
	return nil
}

func (n *NodeBase) PessimisticInsert(data *Data.Data) (splitNode *NodeBase, err error) {
	if !n.canInsert() {
		splitNode = n.Split()
	}

	if data.Key.GE(splitNode.Data[0].Key) {
		err = splitNode.OptimisticInsert(data)
	} else {
		err = n.OptimisticInsert(data)
	}

	return
}

func (n *NodeBase) Split() *NodeBase {
	next := NewNode()

	fillNum := int(math.Ceil(float64(n.MaxSize) * config.FILL_FACTOR))

	next.Size = fillNum
	next.Data = n.Data[fillNum:]

	n.Data = n.Data[:fillNum]
	n.Size = n.MaxSize - fillNum

	if nil != n.NextNode {
		next.NextNode = n.NextNode
	}
	if n.UpNode != nil {
		next.UpNode = n.UpNode
	}
	n.NextNode = next

	return next
}

func (n *NodeBase) canInsert() bool {
	return n.Size < n.MaxSize
}

func (n *NodeBase) Sort() {
	Data.Sort(n.Data)
}

func (n *NodeBase) ToJson() string {
	ret, err := json.Marshal(n)
	if nil != err {
		return ""
	}
	return string(ret)
}

func (n *NodeBase) Find(key *Data.Key) *Data.Data {
	if n.Size <= 0 {
		return nil
	}
	return Data.Find(n.Data, key)
}

func (n *NodeBase) FindLast(key *Data.Key) *Data.Data {
	if n.Size <= 0 {
		return nil
	}
}
