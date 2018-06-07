package EPlan

import ()

type OutputNode struct {
	Inputs   []*OutputNode
	NodeType string
	Location string
}

func NewOutputNodeFromENode(node ENode) *OutputNode {
	loc := node.GetLocation()
	res := &OutputNode{
		NodeType: node.GetNodeType().String(),
		Location: (&loc).GetURL(),
	}
	return res

}

func EPlanOutput(nodes []ENode) *OutputNode {
	if len(nodes) <= 0 {
		return nil
	}
	onodes := []*OutputNode{}
	nodeMap := map[string]int{}

	for i, node := range nodes {
		onode := NewOutputNodeFromENode(node)
		onodes = append(onodes, onode)
		nodeMap[onode.Location] = i
	}

	for i, node := range nodes {
		for _, loc := range node.GetInputs() {
			ad := loc.GetAddress()
			j := nodeMap[ad]
			onodes[i].Inputs = append(onodes[i].Inputs, onodes[j])
		}
	}
	return onodes[len(onodes)-1]
}
