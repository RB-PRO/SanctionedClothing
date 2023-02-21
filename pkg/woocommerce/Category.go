package woocommerce

import (
	"errors"
	"fmt"
)

// Базовая структура данных
type Node struct {
	MeCat            // Содержимое категории
	Children []*Node // Потомки категории
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
	return &Node{Children: []*Node{}, MeCat: MeCat{Id: 0}}

}

// Добавить категорию по родительскому ID
func (root *Node) Add(parentID, id int) error {
	if root == nil {
		return errors.New("Add: Node of nil")
	}

	if parentID == 0 { // Для корневой категории
		return root.addNode(id) // Добавить Категорию в потомка
	}

	// Ищем родительскую категори
	findRoot, errorRoot := root.find(parentID)
	if errorRoot != nil {
		return errorRoot
	}

	return findRoot.addNode(id) // Добавить Категорию в потомка
}

// Выделение памяти/Добавление новой сторуктуры
func (root *Node) addNode(id int) error {
	if root == nil {
		return errors.New("addNode: Node of nil")
	}

	root.Children = append(root.Children,
		&Node{
			Children: []*Node{},
			MeCat:    MeCat{Id: id}})

	return nil
}

// Поиск подкатегории по ID. Доступно извне.
// Возвращает ссылку на значение или его булево значение
func (root *Node) FindId(id int) (*Node, bool) {
	findNode, _ := root.find(id)
	if findNode == nil {
		return nil, false
	}
	return findNode, false
}

// Поиска подкатегории по ID
func (root *Node) find(id int) (*Node, error) {
	if root == nil {
		return nil, errors.New("FindId: Пустой объект")
	}
	for _, val := range root.Children { // Цикл по потомкам
		// Если была найдена подкатегория
		if val.Id == id {
			return val, nil
		}

		// Ищем в дочерних подкатегориях
		FindVal, valError := val.find(id)
		if valError == nil {
			return FindVal, nil
		}
	}
	return nil, errors.New("FindId: не найден id")
}

// Вывод всех категорий
func (root *Node) PrintInorder(prefix string) {
	if root == nil {
		return
	}

	fmt.Println(prefix, root.Id)
	for _, val := range root.Children {
		val.PrintInorder(prefix + "-")
	}
}
