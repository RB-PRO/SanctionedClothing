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
	for i := 0; i < 10; i++ {

		node := woocommerce.NewCategoryes()
		var err error
		err = node.Add(0, 1)
		if err != nil {
			fmt.Println(err)
		}
		err = node.Add(0, 2)
		if err != nil {
			fmt.Println(err)
		}
		err = node.Add(1, 3)
		if err != nil {
			fmt.Println(err)
		}
		err = node.Add(1, 4)
		if err != nil {
			fmt.Println(err)
		}
		err = node.Add(3, 5)
		if err != nil {
			fmt.Println(err)
		}
		err = node.Add(3, 6)
		if err != nil {
			fmt.Println(err)
		}
		err = node.Add(4, 55)
		if err != nil {
			fmt.Println(err)
		}
		err = node.Add(4, 66)
		if err != nil {
			fmt.Println(err)
		}
		err = node.Add(66, 43234)
		if err != nil {
			fmt.Println(err)
		}

		findNode, isFind := node.FindId(43234)
		if isFind {
			fmt.Println("!!", findNode.Id)
		}

		node.PrintInorder("-")
		fmt.Println()
	}
}
