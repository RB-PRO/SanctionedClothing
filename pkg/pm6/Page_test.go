package pm6

import (
	"strconv"
	"testing"
)

func TestAllPages(t *testing.T) {
	pagesInt := AllPages("/null/.zso?s=brandNameFacetLC/asc/productName/asc/")
	if pagesInt != 1131 {
		t.Error("Неправильное к-во товаров. По ссылке https://www.6pm.com/null/.zso?s=brandNameFacetLC/asc/productName/asc/ Должно быть \"1131\", а получено " + "\"" + strconv.Itoa(pagesInt) + "\"")
	}
}
