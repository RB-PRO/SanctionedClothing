package app

import (
	"fmt"
	"log"

	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
	"github.com/mrsinham/catego"
)

func RunWoocommerce() {
	consumer_key, _ := DataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := DataFile("secret_key")     // Секретный код пользователя

	// Авторизация
	userWC, _ := woocommerce.New(consumer_key, secret_key)

	// Проверка на авторизацию
	if ok := userWC.IsOrder(); ok != nil {
		log.Fatalln(ok)
	}

	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		log.Fatalln(errPLC)
	}

	fmt.Println("len(plc)", len(plc.Category))
	for _, categ := range plc.Category {
		if categ.Name == "one" {
			fmt.Println(categ.ID, categ.Name, categ.Slug)
			fmt.Printf("%+v", categ)
		}
		if categ.Name == "two" {
			fmt.Println(categ.ID, categ.Name, categ.Slug)
			fmt.Printf("%+v", categ)
		}
		if categ.Name == "tree" {
			fmt.Println(categ.ID, categ.Name, categ.Slug)
			fmt.Printf("%+v", categ)
		}
	}

	//var cat woocommerce.MeCat

}
func Cats() {
	cs := woocommerce.NewNodeSourse()
	nodes, err := catego.NewTree(cs)
	if err != nil {
		// catch err
	}

	nodes.Add(catego.ID(1), catego.ID(0))
	nodes.Add(catego.ID(2), catego.ID(0))
	nodes.Add(catego.ID(3), catego.ID(0))
	nodes.Add(catego.ID(4), catego.ID(0))

	nodes.Add(catego.ID(31), catego.ID(1))
	nodes.Add(catego.ID(41), catego.ID(1))

	nodes.Add(catego.ID(21), catego.ID(5))

	//nodes.Get(catego.ID(0))

	fmt.Println(nodes.Get(catego.ID(1)))
	childNode, _ := nodes.Get(catego.ID(1))
	for _, val := range childNode.Children {
		fmt.Println(val.ID)
	}
}
