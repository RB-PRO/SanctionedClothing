package woocommerce

import (
	"fmt"
)

type MeCat struct {
	Id    int    // ID в системе WC
	Name  string // Название категории
	Slug  string // Label
	Image string // Ссылка на картинку
}

type NodeCategory struct {
	MeCat
	children []*NodeCategory
}

func (root *NodeCategory) Add(parentId, id int) {
	nodeTable := map[int]*NodeCategory{}
	fmt.Printf("add: id=%v parentId=%v\n", id, parentId)

	node := &NodeCategory{MeCat: MeCat{Id: id}, children: []*NodeCategory{}}
	if parentId == 0 {
		root = node
	} else {

		parent, ok := nodeTable[parentId]
		if !ok {
			fmt.Printf("add: parentId=%v: not found\n", parentId)
			return
		}

		parent.children = append(parent.children, node)
	}

	nodeTable[id] = node
}

func (root *NodeCategory) ShowNode(prefix string) {
	if prefix == "" {
		fmt.Printf("%v\n\n", root.Id)
	} else {
		fmt.Printf("%v %v\n\n", prefix, root.Id)
	}
	for _, n := range root.children {
		n.ShowNode(prefix + "--")
	}
}

func (root *NodeCategory) Show() {
	if root == nil {
		fmt.Printf("show: root node not found\n")
		return
	}
	fmt.Printf("RESULT:\n")
	root.ShowNode("")
}
