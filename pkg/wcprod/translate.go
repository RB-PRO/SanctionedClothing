package wcprod

import (
	"strings"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	gt "github.com/bas24/googletranslatefree"
)

func ProductTranslate(prod bases.Product2) bases.Product2 {

	prod.Description.Eng = strings.ReplaceAll(prod.Description.Eng, "\t", "")
	prod.Description.Eng = strings.ReplaceAll(prod.Description.Eng, "#", "")
	prod.Description.Rus, _ = gt.Translate(prod.Description.Eng, "en", "ru")
	prod.Name, _ = gt.Translate(prod.Name, "en", "ru")
	prod.FullName, _ = gt.Translate(prod.FullName, "en", "ru")
	prod.FullName = strings.ReplaceAll(prod.FullName, "Артикул:", "")

	//tr := translate.New("trnsl.1.1.20170505T201046Z.765061fd7d327f2f.c80d8b95dd956de79d7f9537011fcd3cc802e6e2")
	//tr := translate.New("trnsl.1.1.20191023T124920Z.63524b1f3817bdc2.1719c9be2a2e95a9ce652519943ee104fb9e0a56")
	//tr := translate.New("trnsl.1.1.20190120T184305Z.c3a652a65ff5dac8.3a47d3f48cf9619b3a0d89ad5296f28c220f85ad")

	/*
		response, err := tr.GetLangs("en")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(response.Langs)
			fmt.Println(response.Dirs)
		}

		translation, err := tr.Translate("ru", prod.Description.Eng)
		if err != nil {
			fmt.Println(err)
		} else {
			prod.Description.Rus = translation.Result()
		}
		translation, err = tr.Translate("ru", prod.Name)
		if err != nil {
			fmt.Println(err)
		} else {
			prod.Name = translation.Result()
		}
		translation, err = tr.Translate("ru", prod.FullName)
		if err != nil {
			fmt.Println(err)
		} else {
			prod.FullName = translation.Result()
		}
	*/

	return prod
}
