package pm6

import (
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

func TestParseProduct(t *testing.T) {
	var prod bases.Product2
	prod.Item = make(map[string]bases.ProdParam)
	prod.Specifications = make(map[string]string)
	ParseProduct(&prod, "/p/1-state-balloon-sleeve-crew-neck-sweater-wild-oak/product/9621708/color/836781")
	t.Log(PrintProduct2(prod))
	if prod.Article != "9621708" {
		t.Error("Неправильный артикул. Должно быть \"9621708\", а получено " + "\"" + prod.Article + "\"")
	}
}
