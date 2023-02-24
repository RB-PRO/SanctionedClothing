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

	// Ссылка на товар:
	// https://usmall.ru/product/3252429-faux-fur-hooded-teddy-coat-max-studio
	URL_product := "/product/3252429-faux-fur-hooded-teddy-coat-max-studio"
	CatalogsCatUsmall := usmall.CatalogsCat(usmall.URL + URL_product) // Получить каталог товара
	ProductNumber, _ := usmall.CodeOfLink(URL_product)                // Вычленить код товараw
	ware, _ := usmall.Ware(ProductNumber)                             // Получить запрос с API

	var variety bases.Variety2
	variety.Product = make([]bases.Product2, 1) // выделить память
	variety.Product[0].Cat = CatalogsCatUsmall  // Заполняем структуру категорий

	usmall.WareInProduct2(&variety.Product[0], ware) // Преобразовать в домашнюю структуру
	var CatIDcreate int                              // ID новой или старой категории
	for i := 0; i < 4; i++ {
		findNode, findNodeBool := NodeCategoryes.FindSlug(variety.Product[0].Cat[i].Slug)
		if !findNodeBool { // Если категория не добавлена

			// Добавляем в дерево категорий
			NodeCategoryes.Add(CatIDcreate, woocommerce.MeCat{Id: CatIDcreate, Name: variety.Product[0].Cat[i].Name, Slug: variety.Product[0].Cat[i].Slug})

			// То добавляем её в WC
			cat := woocommerce.MeCat{
				Name:     variety.Product[0].Cat[i].Name,
				Slug:     variety.Product[0].Cat[i].Slug,
				ParentID: CatIDcreate,
			}

			// Добавить категорию на WP
			var ParentError error
			CatIDcreate, ParentError = userWC.AddCat_WC(cat)
			if ParentError != nil {
				fmt.Println(ParentError)
			}

		} else {
			CatIDcreate = findNode.Id
		}
	}

	fmt.Println("ID новой актуальной категории товара - ", CatIDcreate)
	//fmt.Printf("%+v", variety.Product[0])

	tagId := woocommerce.FindIdTagSlug(tags, variety.Product[0].GenderLabel)

	fmt.Println("tagId", tagId)
	// Добавление товара
	prodWC := woocommerce.Product2ProductWC(variety.Product[0], CatIDcreate, tagId) // конфертирование товара
	errorADD := userWC.AddProduct_WC(prodWC)                                        // Добавляем товар
	if errorADD != nil {
		fmt.Println(errorADD)
	}
}
