package app

import (
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/usmall"
)

func Run() {
	items := new(usmall.Variety)

	items.ParsePodSection("https://usmall.ru/products/women/clothes/faux-fur-shearling-coats")

	fmt.Println(items)
}
