package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
	wc "github.com/hiscaler/woocommerce-go"
	config "github.com/hiscaler/woocommerce-go/config"
	"github.com/hiscaler/woocommerce-go/entity"
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

	// Создать Мапу тэгов
	tagMap := woocommerce.MapTags(tags)

	// Получить дерево категорий
	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		log.Fatalln(errPLC)
	}

	// Дерево категорий - Формирование внутренней структуры
	NodeCategoryes := woocommerce.NewCategoryes()
	for _, categ := range plc.Category {
		addingCategory := woocommerce.MeCat{
			Id:   categ.ID,
			Name: categ.Name,
			Slug: categ.Slug,
		}
		NodeCategoryes.Add(categ.Parent, addingCategory)
	}

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
				 Length: 23 in`},
				Item: map[string]bases.ProdParam{
					"wild-oak": bases.ProdParam{
						Link:     "/product/9621708/color/836781",
						ColorEng: "Wild Oak",
						Price:    42.0,
						Size:     []string{"SM", "LG", "XL"},
						Image:    []string{"https://m.media-amazon.com/images/I/91GJ2hRcTeL.jpg", "https://m.media-amazon.com/images/I/91WQzGVObeL.jpg", "https://m.media-amazon.com/images/I/913KXCLH1lL.jpg", "https://m.media-amazon.com/images/I/71a8c4Fw+uL.jpg"},
					},
					"antique-white": bases.ProdParam{
						Link:     "/product/9621708/color/26216",
						ColorEng: "Antique White",
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
	fmt.Println("ID категории", idCat)

	// Аттрибуты
	attr, errAttr := userWC.ProductsAttributes()
	if errAttr != nil {
		log.Fatalln(errAttr)
	}
	idAttrColor, isFind_AttrColor := attr.Find_id_of_name("Цвет")
	if isFind_AttrColor != nil {
		fmt.Println("Не нашёл аттрибут Цвета")
	}
	fmt.Println("ID аттрибута Цвета", idAttrColor)
	idAttrSize, isFind_AttrSize := attr.Find_id_of_name("Размер")
	if isFind_AttrSize != nil {
		fmt.Println("Не нашёл аттрибут Размера")
	}
	fmt.Println("ID аттрибута Размера", idAttrSize)

	// Собираем
	idGender, isGenderSlug := bases.GenderBook(variet.Product[0].GenderLabel)
	if !isGenderSlug {
		fmt.Println("Не найден гендер.", idGender)
	}

	/*
		// Создать структуру добавления товара
		prodWC := woocommerce.Product2ProductWC(variet.Product[0], idCat, tagMap[idGender])

		// Добавление товара
		errorAddProd := userWC.AddProduct_WC(prodWC)
		if errorAddProd != nil {
			fmt.Println(errorAddProd)
		}
	*/

	// **************************************
	// Новое добавление товара

	// Read you config
	b, err := os.ReadFile("config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	wooClient := wc.NewClient(c)

	//paramAttr:=wc.Term

	// Получу все аттрибуты и сохраню в мапу их ID, где ключ - цвет
	//wild-oak antique-white
	//tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(wooClient, idAttrColor, variet.Product[0].Item["wild-oak"].ColorEng, "wild-oak")
	//fmt.Println("Для данного товара Аттрибуты цвета будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)

	// Нужно за раз с категориями загружать.
	paramVariableProduct := wc.CreateProductRequest{
		Name:             variet.Product[0].Name,
		Type:             "variable",
		SKU:              variet.Product[0].Article,
		Description:      variet.Product[0].Description.Eng,
		Tags:             []entity.ProductTag{{ID: tagMap[variet.Product[0].GenderLabel], Slug: variet.Product[0].GenderLabel}},
		ShortDescription: variet.Product[0].FullName,
		RegularPrice:     228.0,

		Categories: []entity.ProductCategory{{ID: idCat}},

		MetaData: []entity.Meta{
			entity.Meta{
				Key:   "Цвет",
				Value: "Wild Oak",
			},
			entity.Meta{
				Key:   "Цвет",
				Value: "wild-oak",
			},
			entity.Meta{
				Key:   "Размер",
				Value: "S",
			},
		},
	}

	item, errorItem := wooClient.Services.Product.Create(paramVariableProduct)
	if errorItem != nil {
		log.Fatal(errorItem)
	}
	itemID := item.ID
	fmt.Println("Done itemID", itemID)

	// *******************************************

	//wild-oak
	//params:=wooClient.Services.ProductVariation.

}

func AddAttr(wooClient *wc.WooCommerce, idAttrColor int, newName, NewSlug string) (tecalAttrId int, tecalAttrName string, tecalAttrSlug string) {
	items, total, _, _, _ := wooClient.Services.ProductAttributeTerm.All(idAttrColor, wc.ProductAttributeTermsQueryParaTerms{Search: newName})
	//totalPages, isLastPage, ProductAttributeTermAll
	// Если такого цвета не существует, то создаём его
	if total == 0 {
		AttributeTermCreate, errorCreate := wooClient.Services.ProductAttributeTerm.Create(idAttrColor, wc.CreateProductAttributeTermRequest{
			Name:        newName,
			Slug:        NewSlug,
			Description: "Создано автоматически",
		})
		if errorCreate != nil {
			fmt.Println(errorCreate)
		}
		tecalAttrId = AttributeTermCreate.ID
		tecalAttrName = AttributeTermCreate.Name
		tecalAttrSlug = AttributeTermCreate.Slug
	} else {
		//fmt.Println("total", total)
		//fmt.Println("totalPages", totalPages)
		//fmt.Println("isLastPage", isLastPage)
		//fmt.Println("ProductAttributeTermAll", ProductAttributeTermAll)
		tecalAttrId = items[0].ID
		tecalAttrName = items[0].Name
		tecalAttrSlug = items[0].Slug
	}
	return tecalAttrId, tecalAttrName, tecalAttrSlug
}
