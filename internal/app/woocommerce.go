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
	var errAdd error
	root := woocommerce.NewCategoryes()

	errAdd = root.Add(0, 1)
	if errAdd != nil {
		fmt.Println(errAdd)
	}
	errAdd = root.Add(0, 2)
	if errAdd != nil {
		fmt.Println(errAdd)
	}
	errAdd = root.Add(0, 5)
	if errAdd != nil {
		fmt.Println(errAdd)
	}

	errAdd = root.Add(1, 3)
	if errAdd != nil {
		fmt.Println(errAdd)
	}
	errAdd = root.Add(1, 4)
	if errAdd != nil {
		fmt.Println(errAdd)
	}

	//root.Add(5, 22)
	//root.Add(3, 9)

	root.PrintInorder("-")

}
