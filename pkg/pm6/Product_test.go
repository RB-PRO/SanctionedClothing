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
		// Цвет
		answerColor := "wild-oak"
		if entityColor.ColorEng != answerColor {
			t.Error("Для цвета \""+color+"\" цвет должен быть:", answerColor, "а получен\n>", prod.Item[color].ColorEng)
		}
		// Ссылка на товар
		answerLink := "/product/9621708/color/836781"
		if entityColor.Link != answerLink {
			t.Error("Для цвета \""+color+"\" должна быть ссылка на товар", answerLink, "а получена\n>", prod.Item[color].Link)
		}
		// Цена
		answerPrice := 42.0
		if entityColor.Price != answerPrice {
			t.Error("Для цвета \""+color+"\" цена должна быть:", answerPrice, "а получена\n>", prod.Item[color].Price)
		}
		// Размеры
		answerSize := []string{"SM", "LG", "XL"}
		if !Equal(entityColor.Size, answerSize) {
			t.Error("Для цвета \""+color+"\" должны быть размеры:", answerSize, "а получены\n>", prod.Item[color].Size)
		}
		// Картинки
		answerPicture := []string{"https://m.media-amazon.com/images/I/91GJ2hRcTeL.jpg", "https://m.media-amazon.com/images/I/91WQzGVObeL.jpg", "https://m.media-amazon.com/images/I/913KXCLH1lL.jpg", "https://m.media-amazon.com/images/I/71a8c4Fw+uL.jpg"}
		if !Equal(entityColor.Image, answerPicture) {
			t.Error("Для цвета \""+color+"\" должны быть картинки:", answerPicture, "а получены\n>", prod.Item[color].Image)
		}

	}
}

// Equal проверяет, что a и b содержат одинаковые элементы.
// nil аргумент эквивалентен пустому срезу.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestPictureCode(t *testing.T) {
	var input, output, result string
	var err error

	input = "https://m.media-amazon.com/images/I/91GJ2hRcTeL._AC_SR58.88,73.60000000000001_.jpg"
	output = "91GJ2hRcTeL"
	if result, err = PictureCode(input); err != nil {
		t.Error("Преобразователь ссылка в код картинки: из входного", input, "должно было получиться:", output, "\nОднако получена ошибка:", err)
	} else {
		if result != output {
			t.Error("Преобразователь ссылка в код картинки: из входного", input, "должно было получиться:", output, "\nОднако получено:", result)
		}
	}
}
