package app

import (
	"fmt"
	"time"

	"github.com/RB-PRO/SanctionedClothing/pkg/usmall"
)

func Run() {

	podSections := usmall.ParsePodSection()

	fmt.Println(len(podSections.Link))
	fmt.Println(podSections.Link[0])

	var variety usmall.Variety
	variety.ParsePage("/products/boy/clothes/kids-robes")

	fmt.Printf("Всего %d товаров. Ссылка на товар %#v\n", len(variety.Product), variety.Product[0].Link)

	// Тестовый парсинг карточки товара с шубой
	//variety.Product[0].Link = "/product/477964-cropped-faux-fur-jacket-avec-les-filles"
	//variety.Product[0].ParseProduct()
	//fmt.Printf("Товар: %#v\n", variety.Product[0])

	time.Sleep(time.Second)
	for i := 0; i < len(variety.Product); i++ {
		fmt.Println(i, usmall.URL+variety.Product[i].Link)
		variety.Product[i].ParseProduct()
		time.Sleep(time.Second)
	}

	variety.SaveXlsx("usmoll")
}
