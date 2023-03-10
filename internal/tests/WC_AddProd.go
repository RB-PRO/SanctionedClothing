package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
	wc "github.com/hiscaler/woocommerce-go"
	config "github.com/hiscaler/woocommerce-go/config"
	"github.com/hiscaler/woocommerce-go/entity"
	//"honnef.co/go/tools/config"
	//"honnef.co/go/tools/config"
	//"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
	//config "github.com/hiscaler/woocommerce-go/config"
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
	fmt.Println(tagMap)

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
				Manufacturer:   "1.STATE",
				Name:           "Balloon Sleeve Crew Neck Sweater",
				FullName:       "Complete your cool-weather look with the soft and cozy 1.STATE™ Balloon Sleeve Crew Neck Sweater.",
				Link:           "/p/1-state-balloon-sleeve-crew-neck-sweater-antique-white/product/9621708/color/26216",
				Article:        "9621708",
				Cat:            bases.Cat{{"Женщины", "women"}, {"Clothing", "clothing"}, {"Sweaters", "sweaters"}, {"1.STATE", "1-state"}},
				GenderLabel:    "women",
				Specifications: map[string]string{"Length": "23 in"},
				Size:           []string{"SM", "LG", "XL"},
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
	idManuf, isFind_AttrManuf := attr.Find_id_of_name("Производитель")
	if isFind_AttrManuf != nil {
		fmt.Println("Не нашёл аттрибут Производителя")
	}
	fmt.Println("ID аттрибута Производителя", idManuf)

	//
	wooClient := wc.NewClient(c)

	fmt.Println(AddProduct(userWC, wooClient, variet.Product[0], tagMap, NodeCategoryes, idAttrColor, idAttrSize, idManuf))

	//paramAttr:=wc.Term

	// Получу все аттрибуты и сохраню в мапу их ID, где ключ - цвет
	//wild-oak antique-white
	//tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(wooClient, idAttrColor, variet.Product[0].Item["wild-oak"].ColorEng, "wild-oak")
	//fmt.Println("Для данного товара Аттрибуты цвета будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)

	// *******************************************
	/*
		paramVar := wc.CreateProductVariationRequest{
			SKU:   variet.Product[0].Article + "wild-oak",
			Image: &entity.ProductImage{Src: variet.Product[0].Item["wild-oak"].Image[0]},
		}
		itemVar, errvar := wooClient.Services.ProductVariation.Create(itemID, paramVar)
		if err != nil {
			fmt.Println(errvar)
		}
		fmt.Println(itemVar.ID)
	*/
}

func AddProduct(userWC *woocommerce.User, wooC *wc.WooCommerce, product bases.Product2, tagMap map[string]int, NodeCategoryes *woocommerce.Node, idAttrColor int, idAttrSize int, idManuf int) error {

	ManufrId, ManufName, ManufSlug := AddAttr(wooC, idAttrColor, "Производитель", product.Manufacturer)
	fmt.Println("Для данного товара Аттрибуты Производителя:", ManufrId, ManufName, ManufSlug)

	// Создать категории для товаров и получить её ID
	idCat, errorAddCat := userWC.AddCat(NodeCategoryes, product.Cat)
	if errorAddCat != nil {
		fmt.Println("Error IDCAT")
	}
	fmt.Println("ID категории", idCat)

	// Создаём аттрибуты товара для цвета
	for key := range product.Item {
		tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(wooC, idAttrColor, product.Item[key].ColorEng, key)
		fmt.Println("Для данного товара Аттрибуты цвета будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)
	}
	// Создаём аттрибуты товара для Размера
	for _, valSize := range product.Size {
		tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(wooC, idAttrSize, valSize, bases.FormingColorEng(valSize))
		fmt.Println("Для данного товара Аттрибуты размера будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)
	}

	// Собираем гендер для загрузки в теги товара
	idGender, isGenderSlug := bases.GenderBook(product.GenderLabel)
	if !isGenderSlug {
		fmt.Println("Не найден гендер.", idGender)
	}
	fmt.Println("Гендр:", idGender)

	// Создаём массив цветов с полными назвавниями
	var colors []string
	for _, colorSet := range product.Item {
		colors = append(colors, colorSet.ColorEng)
	}
	// Сделаю массив со всеми изображениями
	imageInput := make([]entity.ProductImage, 0)
	var chet int
	for _, colorItemValue := range product.Item {
		for indexImage, valueImage := range colorItemValue.Image {
			if chet == 0 {
				imageInput = append(imageInput, entity.ProductImage{
					Src:  valueImage,
					Name: valueImage + strconv.Itoa(indexImage) + ".jpg",
					Alt:  valueImage + strconv.Itoa(indexImage),
				})
			}
			imageInput = append(imageInput, entity.ProductImage{
				Src:  valueImage,
				Name: valueImage + strconv.Itoa(indexImage) + ".jpg",
				Alt:  valueImage + strconv.Itoa(indexImage),
			})
			chet++
		}
	}
	fmt.Println(product.GenderLabel)
	// Структура с исходным товаром
	paramVariableProduct := wc.CreateProductRequest{
		Name:             product.Name,
		Type:             "variable",
		SKU:              product.Article,
		Description:      product.Description.Eng,
		Tags:             []entity.ProductTag{{Name: idGender, Slug: product.GenderLabel}},
		ShortDescription: product.FullName,
		RegularPrice:     228.0,
		Slug:             bases.FormingColorEng(product.Name),

		Images: imageInput,

		Categories: []entity.ProductCategory{{ID: idCat}},

		Attributes: []entity.ProductAttribute{
			{
				ID:      idManuf,
				Options: []string{product.Manufacturer},
				Visible: true,
			},
			{
				ID:        idAttrColor,
				Variation: true,
				Visible:   true,
				Options:   colors,
			},
			{
				ID:        idAttrSize,
				Variation: true,
				Visible:   true,
				Options:   product.Size,
			},
		},
	}

	//asd := entity.ProductVariation{}

	item, errorItem := wooC.Services.Product.Create(paramVariableProduct)
	if errorItem != nil {
		log.Fatal(errorItem)
	}
	itemID := item.ID
	fmt.Println("Done itemID", itemID)

	// Вариационные товары
	for colorKey, colorItemValue := range product.Item {
		/*
			// Массив картинок. Но WC не позволяет загрузить картинки в вариационный товар
			imageInput := make([]entity.ProductImage, 0)
			for indexImage, valueImage := range colorItemValue.Image {
				imageInput = append(imageInput, entity.ProductImage{
					Src:  valueImage,
					Name: valueImage + strconv.Itoa(indexImage) + ".jpg",
					Alt:  valueImage + strconv.Itoa(indexImage),
				})
			}
		*/
		itemVar, errvar := wooC.Services.ProductVariation.Create(itemID, wc.CreateProductVariationRequest{
			SKU:          product.Article + colorKey,
			RegularPrice: colorItemValue.Price,
			Description:  "Цвет: " + colorItemValue.ColorEng + "\n" + product.Description.Eng,
			Image: &entity.ProductImage{
				Src:  colorItemValue.Image[0],
				Name: colorItemValue.ColorEng + ".jpg",
				Alt:  colorItemValue.ColorEng,
			},
			//Images: imageInput,
		})
		if errvar != nil {
			fmt.Println(errvar)
		}
		fmt.Println("Add variation product", itemVar.ID)
	}

	PostSmartImageErr := userWC.PostSmartImage(itemID)
	if PostSmartImageErr != nil {
		fmt.Println(PostSmartImageErr)
	}

	return nil
}

func AddAttr(wooClient *wc.WooCommerce, idAttrColor int, newName, NewSlug string) (tecalAttrId int, tecalAttrName string, tecalAttrSlug string) {
	items, total, _, _, _ := wooClient.Services.ProductAttributeTerm.All(idAttrColor, wc.ProductAttributeTermsQueryParaTerms{Slug: NewSlug})
	// Если такого цвета не существует, то создаём его
	if total == 0 {
		AttributeTermCreate, errorCreate := wooClient.Services.ProductAttributeTerm.Create(idAttrColor, wc.CreateProductAttributeTermRequest{
			Name:        newName,
			Slug:        NewSlug,
			Description: "Создано автоматически при загрузке товара",
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
