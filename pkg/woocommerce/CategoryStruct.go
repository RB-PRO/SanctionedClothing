// Дерево категорий.
// Изначально создайте нулевую категорию:
//
//	root := woocommerce.NewCategoryes()
//
// Далее Вы можете приступать к добавлению элементов по родительсному ID:
//
//	root.Add(0, 1)
//	root.Add(0, 2)
//	root.Add(1, 3)
//
// После добавления Вы имеете возможность расспечатать дерево:
//
// root.PrintInorder("-")
package woocommerce

import (
	"errors"
	"fmt"
	"strconv"
)

// Базовая структура данных
type Node struct {
	MeCat            // Содержимое категории
	Children []*Node // Потомки категории
}

// Внутренняя структура данных
type MeCat struct {
	Id       int    `json:"-"`                // ID в системе WC
	ParentID int    `json:"parent,omitempty"` // ID родителя - используется только при загрузке категории ан WC
	Name     string `json:"name"`             // Название категории
	Slug     string `json:"slug"`             // Label
	Image    string `json:"image,omitempty"`  // Ссылка на картинку
}

// Создать новый список категорий.
func NewCategoryes() *Node {
	return &Node{Children: []*Node{}, MeCat: MeCat{Id: 0}}
}

// Добавить категорию по родительскому ID
func (root *Node) Add(parentID int, MeStruct MeCat) error {
	if root == nil {
		return errors.New("Add: Node of nil")
	}

	if parentID == 0 { // Для корневой категории
		return root.addNode(MeStruct) // Добавить Категорию в потомка
	}

	// Ищем родительскую категори
	findRoot, errorRoot := root.find(parentID)
	if errorRoot != nil {
		return errorRoot
	}

	return findRoot.addNode(MeStruct) // Добавить Категорию в потомка
}

// Выделение памяти/Добавление новой сторуктуры
func (root *Node) addNode(MeStruct MeCat) error {
	if root == nil {
		return errors.New("addNode: Node of nil")
	}

	root.Children = append(root.Children,
		&Node{
			Children: []*Node{},
			MeCat:    MeStruct})

	return nil
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

///////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////

// Поиска подкатегории по ID
func (root *Node) find(id int) (*Node, error) {
	if root == nil {
		return nil, errors.New("FindId: Пустой объект")
	}
	for _, val := range root.Children { // Цикл по потомкам
		// Если была найдена подкатегория
		if val.Id == id {
			return val, nil
		} else { // Ищем в дочерних подкатегориях
			FindVal, valError := val.find(id)
			if valError == nil {
				return FindVal, nil
			}
		}
	}
	return nil, errors.New("FindId: не найден id " + strconv.Itoa(id))
}

// Поиск подкатегории по ID. Доступно извне.
// Возвращает ссылку на значение или его булево значение
func (root *Node) FindId(id int) (*Node, bool) {
	findNode, errorFind := root.find(id)
	if errorFind != nil {
		return nil, false
	}
	return findNode, true
}

// ******************************************************************

// Поиска подкатегории по Slug
func (root *Node) findS(slug string) (*Node, error) {
	if root == nil {
		return nil, errors.New("findS: Пустой объект")
	}
	for _, val := range root.Children { // Цикл по потомкам
		// Если была найдена подкатегория
		if val.Slug == slug {
			return val, nil
		} else { // Ищем в дочерних подкатегориях
			FindVal, valError := val.findS(slug)
			if valError == nil {
				return FindVal, nil
			}
		}
	}
	return nil, errors.New("FindId: не найден slug " + slug)
}

// Поиск подкатегории по Slug. Доступно извне.
// Возвращает ссылку на значение или его булево значение
func (root *Node) FindSlug(slug string) (*Node, bool) {
	findNode, errorFind := root.findS(slug)
	if errorFind != nil {
		return nil, false
	}
	return findNode, true
}

// ******************************************************************

// Поиска подкатегории по Name
func (root *Node) findN(name string) (*Node, error) {
	if root == nil {
		return nil, errors.New("findN: Пустой объект")
	}
	for _, val := range root.Children { // Цикл по потомкам
		// Если была найдена подкатегория
		if val.Name == name {
			return val, nil
		} else { // Ищем в дочерних подкатегориях
			FindVal, valError := val.findN(name)
			if valError == nil {
				return FindVal, nil
			}
		}
	}
	return nil, errors.New("FindId: не найден name " + name)
}

// Поиск подкатегории по Name. Доступно извне.
// Возвращает ссылку на значение или его булево значение
func (root *Node) FindName(name string) (*Node, bool) {
	findNode, errorFind := root.findN(name)
	if errorFind != nil {
		return nil, false
	}
	return findNode, true
}
