package bPlusTree

import (
	"github.com/pangzhanbo/bPlusTree/Data"
	"github.com/pangzhanbo/bPlusTree/Node"
)

type Tree interface {
	Insert(data *Data.Data) (err error)
}

type TreeImpl struct {
	RootNode *Node.NodeBase
}

func NewTree() *TreeImpl {
	return &TreeImpl{
		RootNode: Node.NewNode(),
	}
}

func (t *TreeImpl) Find(key *Data.Key) *Data.Data {
	return nil
}

func (t *TreeImpl) FindNode(key *Data.Key) *Node.NodeBase {

	return nil
}
