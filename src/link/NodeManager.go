package link

import (
	"ignition-link/src/system"
)

type NodeManager struct {
	Nodes []Node
}

func (this *NodeManager) RegisterNode(node Node) {
	node.UID = system.RandomUint64()

	this.Nodes = append(this.Nodes, node)
}

func (this *NodeManager) GetNode(uid uint64) *Node {
	for i := 0; i < len(this.Nodes); i++ {
		n := this.Nodes[i]

		if n.UID == uid {
			return &n
		}
	}

	return nil
}
