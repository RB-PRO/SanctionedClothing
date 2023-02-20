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
	return &Node{Children: map[int]*Node{}, Value: MeCat{Id: 0}}

}

// Добавить категорию по родительскому ID
func (root *Node) Add(parentID, id int) error {
	if parentID == 0 { // Для корневой категории
		root.addNode(id) // Добавить Категорию в потомка
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
	root.Children[id] = &Node{Children: map[int]*Node{}, Value: MeCat{Id: id}}
}

// Поиска подкатегории по ID
func (root *Node) FindId(id int) (*Node, error) {
	for _, val := range root.Children { // Цикл по потомкам
		if val != nil {
			// Если была найдена подкатегория
			if val.Value.Id == id {
				return val, nil
			}

			// Ищем в дочерних подкатегориях
			FindVal, valError := val.FindId(id)
			if valError != nil {
				return FindVal, nil
			}
		}
	}
	return nil, errors.New("не найден " + strconv.Itoa(id) + " id")
}

// Вывод всех категорий
func (root *Node) PrintInorder(prefix string) {
	if root == nil {
		return
	}

	fmt.Println(prefix, root.Value.Id)
	for _, val := range root.Children {
		val.PrintInorder(prefix + "-")
	}
}
