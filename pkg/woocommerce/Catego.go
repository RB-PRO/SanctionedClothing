package woocommerce

import (
	"fmt"

	"github.com/mrsinham/catego"
)

type NodeSource struct {
	Id    int    // ID в системе WC
	Name  string // Название категории
	Slug  string // Label
	Image string // Ссылка на картинку
}

func NewNodeSourse() *NodeSource {
	return &NodeSource{}
}
func (node *NodeSource) Next() bool {
	return false
}

func (node *NodeSource) Get() (current catego.ID, parent catego.ID, err error) {
	return catego.ID(1), catego.ID(2), nil
}

// Вывод всех категорий
func (root *NodeSource) PrintInorder(prefix string) {
	fmt.Println(prefix, root.Id)
	for _, val := range root.Children {
		val.PrintInorder(prefix + "-")
	}
}
