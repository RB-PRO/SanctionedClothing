package pm6

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

func TestParseProduct(t *testing.T) {

	// Обычный тест артикла товара
	var prod bases.Product2
	prod.Item = make(map[string]bases.ProdParam)
	prod.Specifications = make(map[string]string)
	link := "/p/1-state-balloon-sleeve-crew-neck-sweater-wild-oak/product/9621708/color/836781"
	ParseProduct(&prod, link)

	// Артикул
	answerArcticle := "9621708"
	if prod.Article != answerArcticle {
		t.Error("Неправильный артикул. Article. Должно быть \"" + answerArcticle + "\", а получено " + "\n>" + prod.Article)
	}
	// Название товара
	answerName := "Balloon Sleeve Crew Neck Sweater"
	if prod.Name != answerName {
		t.Error("Неправильное название товара. Name. Должно быть \"" + answerName + "\", а получено " + "\n>" + prod.Name)
	}
	// Полное название товара
	answerFullName := "Complete your cool-weather look with the soft and cozy 1.STATE™ Balloon Sleeve Crew Neck Sweater."
	if prod.FullName != answerFullName {
		t.Error("Неправильное полное название товара. FullName. Должно быть \"" + answerFullName + "\", а получено " + "\n>" + prod.FullName)
	}
	// Категории
	answerCat := bases.Cat{{"Женщины", "women"}, {"Clothing", "clothing"}, {"Sweaters", "sweaters"}, {"1.STATE", "1-state"}}
	if prod.Cat != answerCat {
		t.Error("Неправильно получены категории товаров. Cat. Должно быть \"", answerCat, "\", а получено\n>", prod.Cat)
	}
	// Прочие аттрибуты
	if prod.Specifications["Length"] != "23 in" {
		t.Error("Неправильно получены аттрибуты товаров. Specifications. Должно быть [\"Length\"]!=\"23 in\", а получено\n>", prod.Specifications)
	}
	// Ссылка на товар
	answerLink := link
	if prod.Link != answerLink {
		t.Error("Неправильная ссылка на товар. Link. Должно быть \"" + answerLink + "\", а получено " + "\n>" + prod.Link)
	}
	// Гендер
	answerGender := "women"
	if prod.GenderLabel != answerGender {
		t.Error("Неправильный гендер. GenderLabel. Должно быть \"" + answerGender + "\", а получено " + "\n>" + prod.GenderLabel)
	}

	color := "wild-oak"
	if entityColor, ok := prod.Item[color]; !ok {
		keys := ""
		for key := range prod.Item {
			keys += ">" + key + "< "
		}
		t.Error("Не добавлен цвет \""+color+"\", однако есть цвета:", keys)
	} else {
		fmt.Printf("%+v\n", entityColor)

		// Ссылка на товар
		answerLink := "/p/1-state-balloon-sleeve-crew-neck-sweater-wild-oak/product/9621708/color/836781"
		if entityColor.Link != answerLink {
			t.Error("Для цвета \""+color+"\" должна быть ссылка на товар", answerLink, "а получена\n>", prod.Item[color].Link)
		}
	}
}
