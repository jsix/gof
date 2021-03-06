package tree

import (
	_ "fmt"
)

func List2Tree(nodeList []FlatNode, nodeFn func(node *TreeNode)) (rootNode *TreeNode) {
	for i, k := range nodeList {
		if k.ID == 0 {
			rootNode = k.Node()
			nodeList = append(nodeList[:i], nodeList[i+1:]...)
			break
		}
	}

	if rootNode == nil {
		rootNode = &TreeNode{
			id:       0,
			Title:    "根节点",
			Value:    "",
			Url:      "",
			Icon:     "",
			Expand:   true,
			Children: nil}
	}
	walkTree(rootNode, nodeList, nodeFn)
	return rootNode
}

func walkTree(node *TreeNode, nodeList []FlatNode, nodeFn func(node *TreeNode)) {
	node.Children = []*TreeNode{}
	for _, v := range nodeList {
		if v.Pid == node.id {
			n := v.Node()
			if nodeFn != nil {
				nodeFn(n)
			}
			node.Children = append(node.Children, n)
			walkTree(n, nodeList, nodeFn)
		}
	}
}
