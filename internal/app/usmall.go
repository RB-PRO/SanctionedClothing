package app

import (
	"fmt"
	"time"

	"github.com/RB-PRO/SanctionedClothing/pkg/usmall"
	"github.com/cheggaaa/pb"
)

func Run() {

	podSections := usmall.ParsePodSection()

	fmt.Println(len(podSections.Link))
	fmt.Println(podSections.Link[0])

	var variety usmall.Variety
	//variety.ParsePage("/products/boy/clothes/kids-robes")

	//fmt.Printf("Всего %d товаров. Ссылка на товар %#v\n", len(variety.Product), variety.Product[0].Link)

	/*
		// Тестовый парсинг карточки товара с шубой с выгрузкой из фронта
		variety.Product[0].Link = "/product/477964-cropped-faux-fur-jacket-avec-les-filles"
		variety.Product[0].ParseProduct()
		fmt.Printf("Товар: %#v\n", variety.Product[0])
	*/

	/*
		// Тестовый парсинг карточки товара с шубой с выгрузкой из API
		variety.Product[0].Link = "/product/477964-cropped-faux-fur-jacket-avec-les-filles"
		MyCode := variety.Product[0].Link     // Код товара
		MyCode, _ = usmall.CodeOfLink(MyCode) // Вычленить код товара
		fmt.Println("MyCode", MyCode)
		ware, _ := usmall.Ware(MyCode) // Получить запрос с API
		variety.Product[0].WareInProduct(ware)
		fmt.Printf("Товар: %#v\n", variety.Product[0])
	*/

	/*
		// Пропарсить одну категорию
		for i := 0; i < len(variety.Product); i++ {
			fmt.Println(i, usmall.URL+variety.Product[i].Link)
			MyCode := variety.Product[i].Link      // Код товара
			MyCode, _ = usmall.CodeOfLink(MyCode)  // Вычленить код товара
			ware, _ := usmall.Ware(MyCode)         // Получить запрос с API
			variety.Product[i].WareInProduct(ware) // Преобразовать в домашнюю структуру
			time.Sleep(100 * time.Microsecond)
		}
	*/

	/*
		time.Sleep(time.Second)
		for i := 0; i < len(variety.Product); i++ {
			fmt.Println(i, usmall.URL+variety.Product[i].Link)
			variety.Product[i].ParseProduct()
			time.Sleep(time.Second)
		}
	*/

	// *************************************************
	// Спасить вообще всё
	fmt.Println("Спарсить все pages, чтобы получить все ссылки")
	bar := pb.StartNew(len(podSections.Link))
	for _, valuePodSection := range podSections.Link {
		bar.Increment() // Прибавляем 1 к отображению
		variety.ParsePage(valuePodSection)
	}
	bar.Finish()

	// Пропарсить всё
	bar2 := pb.StartNew(len(podSections.Link))
	for i := 0; i < len(variety.Product); i++ {
		bar2.Increment() // Прибавляем 1 к отображению
		//fmt.Println(i, usmall.URL+variety.Product[i].Link)
		MyCode := variety.Product[i].Link      // Код товара
		MyCode, _ = usmall.CodeOfLink(MyCode)  // Вычленить код товара
		ware, _ := usmall.Ware(MyCode)         // Получить запрос с API
		variety.Product[i].WareInProduct(ware) // Преобразовать в домашнюю структуру
		time.Sleep(20 * time.Microsecond)
	}
	bar.Finish()
	// *************************************************
	variety.SaveXlsx("usmoll")
}
