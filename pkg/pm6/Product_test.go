package pm6

import (
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
		t.Error("Неправильный артикул. Article. Должно быть \"" + answerArcticle + "\", а получено " + "\"" + prod.Article + "\"")
	}

	// Название товара
	if prod.Name != "Balloon Sleeve Crew Neck Sweater" {
		t.Error("Неправильное поленое название товара. Name. Должно быть \"Balloon Sleeve Crew Neck Sweater\", а получено " + "\"" + prod.Name + "\"")
	}
	// Полное название товара
	answerFullName := "Complete your cool-weather look with the soft and cozy 1.STATE™ Balloon Sleeve Crew Neck Sweater."
	if prod.FullName != answerFullName {
		t.Error("Неправильное полное название товара. FullName. Должно быть \"" + answerFullName + "\", а получено " + "\"" + prod.FullName + "\"")
	}

	// Категории
	answerCat := bases.Cat{{"Женщины", "woman"}, {"Clothing", "clothing"}, {"Sweaters", "sweaters"}, {"1.STATE", "1-state"}}
	if prod.Cat != answerCat {
		t.Error("Неправильно получены категории товаров. Cat. Должно быть \"", answerCat, "\", а получено "+"\"", prod.Cat, "\"")
	}

	// Прочие аттрибуты
	if prod.Specifications["Length"] != "23 in" {
		t.Error("Неправильно получены аттрибуты товаров. Specifications. Должно быть [\"Length\"]!=\"23 in\", а получено", prod.Specifications)
	}

	//Balloon Sleeve Crew Neck Sweater

}
