// Отдельно вынесенный пакет для загрузки товаров на WordPress
// Использует кривую библиотеку.
// Кривая она из-за того, что некоторые параметры не соотносятся с документацией Woocommerce.
// Для упрощения написания кода, локально исправил некоторые строки в скаченной библиотеке. В идеале локально развернуть библиотеку и провести необходимые манипуляции.
package wcprod

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
	wc "github.com/hiscaler/woocommerce-go"
	config "github.com/hiscaler/woocommerce-go/config"
	"github.com/hiscaler/woocommerce-go/entity"
)

// Созовая структура, которая объединяет в себе все необходимые данные для работы с библиотекой и для загрузки товаров
type WcAdd struct {
	NodeCategoryes *woocommerce.Node      // Дерево категорий собственной разработки
	UserWC         *woocommerce.User      // структура пользователя из своей библиотеки
	Tags           []woocommerce.Tag      // Массив тегов, которые присутствуют в WordPress
	TagMap         map[string]int         // Мапа тегов. Вообще бы её вывести отсюда нахрен
	Sttr           woocommerce.Attributes // Структура аттрибутов, которые лежат на WP

	// ID аттрибутов в WordPress.
	IdAttrColor int
	IdAttrSize  int
	IdManuf     int

	WooClient *wc.WooCommerce // Клиент пользовательской библиотеки, с помощью которой добавляю товар
}

// Инициализации базовой структуры загрузки товара
func New() (*WcAdd, error) {
	// Клиент от сторонней библиотеки(пользовательской)
	b, err := os.ReadFile("config_test.json")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Read config error: %s", err.Error()))

	}
	var c config.Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}
	wooClient := wc.NewClient(c)

	// Мой клиент
	consumer_key, _ := DataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := DataFile("secret_key")     // Секретный код пользователя

	userWC, _ := woocommerce.New(consumer_key, secret_key) // Авторизация
	if okErr := userWC.IsOrder(); okErr != nil {           // Проверка на авторизацию
		return nil, okErr
	}

	// Теги
	tags, tagsError := userWC.AllTags_WC()
	if tagsError != nil {
		return nil, tagsError
	}

	// Создать Мапу тэгов
	tagMap := woocommerce.MapTags(tags)

	// Получить дерево категорий
	plc, errPLC := userWC.ProductsCategories()
	if errPLC != nil {
		return nil, errPLC
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

	// Аттрибуты
	attr, errAttr := userWC.ProductsAttributes()
	if errAttr != nil {
		return nil, errAttr
	}
	idAttrColor, isFind_AttrColor := attr.Find_id_of_name("Цвет")
	if isFind_AttrColor != nil {
		return nil, isFind_AttrColor
	}
	idAttrSize, isFind_AttrSize := attr.Find_id_of_name("Размер")
	if isFind_AttrSize != nil {
		return nil, isFind_AttrSize
	}
	idManuf, isFind_AttrManuf := attr.Find_id_of_name("Производитель")
	if isFind_AttrManuf != nil {
		return nil, isFind_AttrManuf
	}

	return &WcAdd{
		WooClient:      wooClient,
		UserWC:         userWC,
		Tags:           tags,
		TagMap:         tagMap,
		NodeCategoryes: NodeCategoryes,
		IdAttrColor:    idAttrColor,
		IdAttrSize:     idAttrSize,
		IdManuf:        idManuf,
	}, nil
}

// Функция добавления товара
func (woo *WcAdd) AddProduct(product bases.Product2) error {

	ManufrId, ManufName, ManufSlug := AddAttr(woo.WooClient, woo.IdAttrColor, "Производитель", product.Manufacturer)
	fmt.Println("Для данного товара Аттрибуты Производителя:", ManufrId, ManufName, ManufSlug)

	// Создать категории для товаров и получить её ID
	idCat, errorAddCat := woo.UserWC.AddCat(woo.NodeCategoryes, product.Cat)
	if errorAddCat != nil {
		fmt.Println("Error IDCAT")
	}
	fmt.Println("ID категории", idCat)

	// Создаём аттрибуты товара для цвета
	for key := range product.Item {
		tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(woo.WooClient, woo.IdAttrColor, product.Item[key].ColorEng, key)
		fmt.Println("Для данного товара Аттрибуты цвета будут:", tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug)
	}
	// Создаём аттрибуты товара для Размера
	for _, valSize := range product.Size {
		tecalAttrColorId, tecalAttrColorName, tecalAttrColorSlug := AddAttr(woo.WooClient, woo.IdAttrSize, valSize, bases.FormingColorEng(valSize))
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
		Description:      product.Description.Rus,
		Tags:             []entity.ProductTag{{Name: idGender, Slug: product.GenderLabel}},
		ShortDescription: product.FullName,
		RegularPrice:     228.0,
		Slug:             bases.FormingColorEng(product.Name),

		Images: imageInput,

		Categories: []entity.ProductCategory{{ID: idCat}},

		Attributes: []entity.ProductAttribute{
			{
				ID:      woo.IdManuf,
				Options: []string{product.Manufacturer},
				Visible: true,
			},
			{
				ID:        woo.IdAttrColor,
				Variation: true,
				Visible:   true,
				Options:   colors,
			},
			{
				ID:        woo.IdAttrSize,
				Variation: true,
				Visible:   true,
				Options:   product.Size,
			},
		},
	}

	//asd := entity.ProductVariation{}

	item, errorItem := woo.WooClient.Services.Product.Create(paramVariableProduct)
	if errorItem != nil {
		log.Fatal(errorItem)
	}
	itemID := item.ID
	fmt.Println("Done itemID", itemID)

	// Вариационные товары
	for colorKey, colorItemValue := range product.Item {
		itemVar, errvar := woo.WooClient.Services.ProductVariation.Create(itemID, wc.CreateProductVariationRequest{
			SKU:          product.Article + colorKey,
			RegularPrice: colorItemValue.Price,
			Description:  "Цвет: " + colorItemValue.ColorEng + "\n" + product.Description.Rus,
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

	PostSmartImageErr := woo.UserWC.PostSmartImage(itemID)
	if PostSmartImageErr != nil {
		fmt.Println(PostSmartImageErr)
	}

	return nil
}
