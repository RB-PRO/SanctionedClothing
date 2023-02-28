package pm6

import (
	"strings"
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
	if prod.Article != "9621708" {
		t.Error("Неправильный артикул. Должно быть \"9621708\", а получено " + "\"" + prod.Article + "\"")
	}

	// Тест размера майки ребёнка
	var prod2 bases.Product2
	prod2.Item = make(map[string]bases.ProdParam)
	prod2.Specifications = make(map[string]string)
	link = "/p/4kids-essential-high-low-tank-top-little-kids-big-kids-navy/product/9450063/color/9"
	ParseProduct(&prod2, link)
	t.Log(PrintProduct2(prod2))
	if itemColors, ok := prod2.Item["navy"]; ok {
		if itemColors.Size[0] != "XS" || itemColors.Size[1] != "SM" {
			t.Error("Неправильные размеры для " + URL + link + " получено " + strings.Join(itemColors.Size, ",") + ", а должно быть XS,SM")
		}
	}
}
