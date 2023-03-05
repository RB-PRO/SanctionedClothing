package tests

import (
	"fmt"
	"log"

	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
)

func AddProd() {
	consumer_key, _ := DataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := DataFile("secret_key")     // Секретный код пользователя
	yandexToken, _ := DataFile("yandexToken")   // Секретный код пользователя

	// Авторизация
	userWC, _ := woocommerce.New(consumer_key, secret_key)

	// Проверка на авторизацию
	if ok := userWC.IsOrder(); ok != nil {
		log.Fatalln(ok)
	}

	// Получить тэги
	tags, tagsError := userWC.AllTags_WC()
	if tagsError != nil {
		log.Fatalln(tagsError)
	}
	fmt.Println(tags)

	// Получить дерево категорий
	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		log.Fatalln(errPLC)
	}

	// Дерево категорий
	NodeCategoryes := woocommerce.NewCategoryes()
	for _, categ := range plc.Category {
		addingCategory := woocommerce.MeCat{
			Id:   categ.ID,
			Name: categ.Name,
			Slug: categ.Slug,
		}
		NodeCategoryes.Add(categ.Parent, addingCategory)
	}
	NodeCategoryes.PrintInorder("-") // печать категорий

	// Добавление товара

}
