package tests

import (
	"fmt"
	"log"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/usmall"
	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
)

func TparseANDadd() {
	consumer_key, _ := DataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := DataFile("secret_key")     // Секретный код пользователя

	// Авторизация
	userWC, _ := woocommerce.New(consumer_key, secret_key)

	// Проверка на авторизацию
	if ok := userWC.IsOrder(); ok != nil {
		log.Fatalln(ok)
	}

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

	// Ссылка для теста:
	// https://usmall.ru/product/3252429-faux-fur-hooded-teddy-coat-max-studio
	ProductNumber := "3252429"
	ware, _ := usmall.Ware(ProductNumber) // Получить запрос с API
	var variety bases.Variety2
	variety.Product = make([]bases.Product2, 1) // выделить память
	variety.Product[0].Cat[0].Name = "Женщины"
	variety.Product[0].Cat[0].Slug = "women"
	variety.Product[0].Cat[1].Name = "Одежда"
	variety.Product[0].Cat[1].Slug = "clothes"
	variety.Product[0].Cat[2].Name = "Пальто"
	variety.Product[0].Cat[2].Slug = "wool-pea-coats"
	variety.Product[0].Cat[3].Name = "Max Studio"
	variety.Product[0].Cat[3].Slug = "max-studio"
	usmall.WareInProduct2(&variety.Product[0], ware) // Преобразовать в домашнюю структуру

	fmt.Printf("%+v", variety.Product[0])
}
