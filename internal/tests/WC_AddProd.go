package tests

import (
	"fmt"
	"log"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
)

func AddProd() {
	consumer_key, _ := DataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := DataFile("secret_key")     // Секретный код пользователя
	//yandexToken, _ := DataFile("yandexToken")   // Секретный код пользователя

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
	fmt.Printf("%#+v", tags)

	// Создать Мапу тэгов
	tagMap := woocommerce.MapTags(tags)
	fmt.Println("Найденные теги из словаря", tagMap)

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

	// Создать тестовый товар
	variet := bases.Variety2{
		[]bases.Product2{
			bases.Product2{
				Name:           "Balloon Sleeve Crew Neck Sweater",
				FullName:       "Complete your cool-weather look with the soft and cozy 1.STATE™ Balloon Sleeve Crew Neck Sweater.",
				Link:           "/p/1-state-balloon-sleeve-crew-neck-sweater-antique-white/product/9621708/color/26216",
				Article:        "9621708",
				Cat:            bases.Cat{{"Женщины", "women"}, {"Clothing", "clothing"}, {"Sweaters", "sweaters"}, {"1.STATE", "1-state"}},
				GenderLabel:    "women",
				Specifications: map[string]string{"Length": "23 in"},
				Description: struct {
					Eng string
					Rus string
				}{Eng: `Complete your cool-weather look with the soft and cozy 1.STATE™ Balloon Sleeve Crew Neck Sweater.
				SKU: #9621708
				Pull-over design with ribbed crew neckline.
				Long balloon sleeves with elongated, ribbed cuffs.
				Classic fit with straight hemline.
				73% acrylic, 24% polyester, 3% spandex.
				Hand wash, dry flat.
				Imported.
				Product measurements were taken using size SM. Please note that measurements may vary by size.
				 Length: 23 in
				Complete your cool-weather look with the soft and cozy 1.STATE™ Balloon Sleeve Crew Neck Sweater.
				SKU: #9621708
				Pull-over design with ribbed crew neckline.
				Long balloon sleeves with elongated, ribbed cuffs.
				Classic fit with straight hemline.
				73% acrylic, 24% polyester, 3% spandex.
				Hand wash, dry flat.
				Imported.
				Product measurements were taken using size SM. Please note that measurements may vary by size.
				 Length: 23 in`},
				Item: map[string]bases.ProdParam{
					"wild-oak": bases.ProdParam{
						Link:     "/product/9621708/color/836781",
						ColorEng: "wild-oak",
						Price:    42.0,
						Size:     []string{"SM", "LG", "XL"},
						Image:    []string{"https://m.media-amazon.com/images/I/91GJ2hRcTeL.jpg", "https://m.media-amazon.com/images/I/91WQzGVObeL.jpg", "https://m.media-amazon.com/images/I/913KXCLH1lL.jpg", "https://m.media-amazon.com/images/I/71a8c4Fw+uL.jpg"},
					},
					"antique-white": bases.ProdParam{
						Link:     "/product/9621708/color/26216",
						ColorEng: "antique-white",
						Price:    31.58,
						Size:     []string{"SM", "LG", "XL"},
						Image:    []string{"https://m.media-amazon.com/images/I/71Mf94kDFvL.jpg", "https://m.media-amazon.com/images/I/71EOOcBc+bL.jpg", "https://m.media-amazon.com/images/I/81PeCItuTmL.jpg", "https://m.media-amazon.com/images/I/71+cz20ouIL.jpg"},
					},
				},
			},
		},
	}

	// Создать категории для товаров и получить её ID
	idCat, errorAddCat := userWC.AddCat(NodeCategoryes, variet.Product[0].Cat)
	if errorAddCat != nil {
		fmt.Println("Error IDCAT")
	}

	idGender, isGenderSlug := bases.GenderBook(variet.Product[0].GenderLabel)
	if isGenderSlug {
		// Создать структуру добавления товара
		prodWC := woocommerce.Product2ProductWC(variet.Product[0], idCat, tagMap[idGender])

		// Добавление товара
		errorAddProd := userWC.AddProduct_WC(prodWC)
		if errorAddProd != nil {
			fmt.Println(errorAddProd)
		}
	}
}
