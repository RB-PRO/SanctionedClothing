package woocommerce

import (
	"errors"
	"fmt"
	"strconv"
)

// Базовая структура данных
type Node struct {
	Value    MeCat         // Содержимое категории
	Children map[int]*Node // Потомки категории
}

// Внутренняя структура данных
type MeCat struct {
	Id    int    // ID в системе WC
	Name  string // Название категории
	Slug  string // Label
	Image string // Ссылка на картинку
}

// Создать новый список категорий.
func NewCategoryes() *Node {
	n := new(Node)
	n.Children = make(map[int]*Node)
	n.Value = MeCat{}
	n.Value.Id = 0
	return n
}

// Добавить категорию по родительскому ID
func (root *Node) Add(parentID, id int) error {
	if parentID == 0 { // Для корневой категории
		root.addNode(id) //// Добавить Категорию в потомка
		return nil
	}

	// Ищем родительскую категори
	findRoot, errorRoot := root.FindId(parentID)
	if errorRoot != nil {
		return errorRoot
	}

	findRoot.addNode(id) // Добавить Категорию в потомка

	return nil
}

// Выделение памяти/Добавление новой сторуктуры
func (root *Node) addNode(id int) {
	root.Children[id] = new(Node)
	root.Children[id].Children = make(map[int]*Node)
	root.Children[id].Value = MeCat{}
	root.Children[id].Value.Id = id
}

// Поиска подкатегории по ID
func (root *Node) FindId(id int) (*Node, error) {
	for _, val := range root.Children {

		// Если была найдена подкатегория
		if val.Value.Id == id {
			return val, nil
		}
		if FindVal, valError := val.FindId(id); valError != nil {
			return FindVal, nil
		}
	}
	return nil, errors.New("не найден " + strconv.Itoa(id) + " id")
}

// Вывод всех категорий
func (t *Node) PrintInorder(prefix string) {
	if t == nil {
		return
	}

	fmt.Println(prefix, t.Value.Id)
	for _, val := range t.Children {
		val.PrintInorder(prefix + "-")
	}
}
