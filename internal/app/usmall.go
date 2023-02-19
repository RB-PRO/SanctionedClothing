package app

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/usmall"
	"github.com/cheggaaa/pb"
)

// func Run(startStr string) {
func Run() {

	podSections := usmall.ParsePodSection()

	/*
		startInt, startError := strconv.Atoi(startStr)
		if startError != nil {
			log.Fatalln("\""+startStr+"\"", "is not a number")
		}*/

	// *************************************************
	// Спасить вообще всё

	//podSections.Link = podSections.Link[:10]
	for indexPodSectioen, valPodSection := range podSections.Link {
		//if indexPodSectioen >= startInt {
		strPodSection := strconv.Itoa(indexPodSectioen) + " > " + strings.ReplaceAll(valPodSection, "/", "-") // Название файла текущей подсекции

		log.Println(" -> ", indexPodSectioen, "/", len(podSections.Link), " ", strPodSection)

		var variety bases.Variety2
		fmt.Println("Спарсить все pages, чтобы получить все ссылки:")
		//bar.Increment() // Прибавляем 1 к отображению
		fmt.Println("->", usmall.URL+valPodSection)
		usmall.ParsePage(&variety, valPodSection)

		// Пропарсить всё
		fmt.Println("Пропарсить всё", len(variety.Product))

		bar2 := pb.StartNew(len(variety.Product))
		for i := 0; i < len(variety.Product); i++ {
			bar2.Increment() // Прибавляем 1 к отображению
			//fmt.Println(i, usmall.URL+variety.Product[i].Link)
			MyCode := variety.Product[i].Link                // Код товара
			MyCode, _ = usmall.CodeOfLink(MyCode)            // Вычленить код товараw
			ware, _ := usmall.Ware(MyCode)                   // Получить запрос с API
			usmall.WareInProduct2(&variety.Product[i], ware) // Преобразовать в домашнюю структуру
			time.Sleep(20 * time.Microsecond)
		}
		bar2.Finish()

		// *************************************************
		//variety.SaveXlsx(strPodSection)
		variety.SaveXlsxCsvs(strPodSection) // Cохранить в формате из ТЗ

		//}
	}
}
