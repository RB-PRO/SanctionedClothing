// Программа, которая парсит сайт PM6 и загружает данные на Wordpress
package pm6wp

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Старт программы, который:
// - считывает входные данные из терминала
// - проверяет данные
// - запускает парсинг и загрузку товаров.

// Взодные данные:
// - Стартовая страница, с которону необходимо начинать парсинг, integer
// - Коэффициент наценки, она же моржа, float
// - стоймость доставки, int
func Start() {
	var PageStart int  // Стартовая страница
	var walrus float64 // Моржа
	var delivery int   // Стоймость доставки

	// Проверка на 3 входных параметра
	fmt.Println(len(os.Args))
	if len(os.Args) != 4 {
		Work(0, 1.5, 150)
		log.Fatalln("Запросов должно быть 3 штуки. Стартовая страница, моржа, доставка. Пример: ./pmwp 0 1.5 100")
	}

	var ErrorConv error

	// Стартовая страница
	if PageStart, ErrorConv = strconv.Atoi(os.Args[1]); ErrorConv != nil {
		log.Fatalln("[pmwp]: Неправильный параметр стартовой страницы. Пример: ./pmwp 0 1.5 100")
	}
	// Моржа
	if walrus, ErrorConv = strconv.ParseFloat(os.Args[2], 64); ErrorConv != nil {
		log.Fatalln("[pmwp]: Неправильный параметр моржа. Пример: ./pmwp 0 1.5 100")
	}
	// Стоймость доставки
	if delivery, ErrorConv = strconv.Atoi(os.Args[3]); ErrorConv != nil {
		log.Fatalln("[pmwp]: Неправильный параметр стоймости доставки. Пример: ./pmwp 0 1.5 100")
	}

	log.Printf("[pmwp]: Начинаю работать по параметрам: Стартовая страница %d, Моржа %f, Стоймость доставки %d.", PageStart, walrus, delivery)

	// Запускаю работу парсера
	Work(PageStart, walrus, delivery)
}
