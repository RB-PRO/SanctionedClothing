package app

import (
	"fmt"
	"log"

	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
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

	fmt.Println()
	var root *woocommerce.NodeCategory
	root.Add(0, 1)
	root.Add(1, 1)
	root.Add(1, 2)
	root.Add(1, 3)

	root.Add(2, 1)
	root.Add(2, 3)

	root.Show()
	root.ShowNode("-")
}
