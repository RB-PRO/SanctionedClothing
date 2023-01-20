package app

import (
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/usmall"
)

func Run() {
	dataUsMall, errorUsMall := usmall.Parse()
	if errorUsMall != nil {
		fmt.Println(errorUsMall)
	}
	fmt.Println(dataUsMall)
}
