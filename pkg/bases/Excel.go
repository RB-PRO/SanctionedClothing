package bases

import (
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Создать книгу
func MakeWorkBook() (*excelize.File, error) {
	// Создать книгу Excel
	f := excelize.NewFile()
	// Create a new sheet.
	_, err := f.NewSheet("main")
	if err != nil {
		return f, err
	}
	f.DeleteSheet("Sheet1")
	return f, nil
}

// Сохранить и закрыть файл
func CloseXlsx(f *excelize.File, filename string) error {
	if err := f.SaveAs(filename + ".xlsx"); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

/*
func WriteOneLine(f *excelize.File, ssheet string, row int, SearchBasicRes SearchBasicResponse, SearchBasicIndex int, GetPartsRemainsByCodeRes GetPartsRemainsByCodeResponse, GetPartsRemainsByCodeIndex int) {
	// SearchBasic
	writeHeadOne(f, ssheet, 1, row, SearchBasicRes.Data.Items[SearchBasicIndex].Code, "")
}
*/

// Сохранить данные в Xlsx
func (variety Variety2) SaveXlsx(filename string) error {
	// Создать книгу
	book, makeBookError := MakeWorkBook()
	if makeBookError != nil {
		return makeBookError
	}

	wotkSheet := "main"
	setHead(book, wotkSheet, 1, "Каталог")                // Catalog
	setHead(book, wotkSheet, 2, "ПодКаталог")             // PodCatalog
	setHead(book, wotkSheet, 3, "Секция")                 // Section
	setHead(book, wotkSheet, 4, "Подсекция")              // PodSection
	setHead(book, wotkSheet, 5, "Название товара")        // Name
	setHead(book, wotkSheet, 6, "Полное название товара") // FullName
	setHead(book, wotkSheet, 7, "Ссылка на товар")        // Link
	setHead(book, wotkSheet, 8, "Артикул")                // Article
	setHead(book, wotkSheet, 9, "Производитель")          // Manufacturer
	setHead(book, wotkSheet, 10, "Цена")                  // Price
	setHead(book, wotkSheet, 11, "Описание товара Rus")   // Description eng
	setHead(book, wotkSheet, 12, "Описание товара Eng")   // Description rus
	setHead(book, wotkSheet, 13, "Цвета")                 // Colors
	setHead(book, wotkSheet, 14, "Картинки")              // Colors
	setHead(book, wotkSheet, 15, "Размеры")               // Size
	startIndexCollumn := 16

	// Создаём мапу, которая будет содержать значения номеров колонок
	colName := make(map[string]int)
	for indexItem, valItem := range variety.Product {
		setCell(book, wotkSheet, indexItem+2, 1, valItem.Cat[0].Name)  // Каталог
		setCell(book, wotkSheet, indexItem+2, 2, valItem.Cat[1].Name)  // ПодКаталог
		setCell(book, wotkSheet, indexItem+2, 3, valItem.Cat[2].Name)  // Секция
		setCell(book, wotkSheet, indexItem+2, 4, valItem.Cat[3].Name)  // Подсекция
		setCell(book, wotkSheet, indexItem+2, 5, valItem.Name)         // Название товара
		setCell(book, wotkSheet, indexItem+2, 6, valItem.FullName)     // Полное название товара
		setCell(book, wotkSheet, indexItem+2, 7, valItem.Link)         // Ссылка на товар
		setCell(book, wotkSheet, indexItem+2, 8, valItem.Article)      // Артикул
		setCell(book, wotkSheet, indexItem+2, 9, valItem.Manufacturer) // Производитель
		//setCell(book, wotkSheet, indexItem+2, 10, valItem.Price)           // Цена
		setCell(book, wotkSheet, indexItem+2, 11, valItem.Description.Rus) // Описание товара Rus
		setCell(book, wotkSheet, indexItem+2, 12, valItem.Description.Eng) // Описание товара Eng
		//setCell(book, wotkSheet, indexItem+2, 13, valItem.Colors)          // Цвета
		//setCell(book, wotkSheet, indexItem+2, 14, valItem.Size)            // Размеры

		// Обработка мапы картинок
		for key, val := range valItem.Item {

			setCell(book, wotkSheet, indexItem+2, 7, val.Link) // ссылка на товар с картинкой

			setCell(book, wotkSheet, indexItem+2, 10, val.Price) // Цена
			setCell(book, wotkSheet, indexItem+2, 13, key)       // Цвет
			setCell(book, wotkSheet, indexItem+2, 14, val.Image) // Картинка
			setCell(book, wotkSheet, indexItem+2, 15, val.Size)  // Size
			//setCell(book, wotkSheet, indexItem+2, 7, URL+val.Link) // Ссылка на товар

			if _, ok := colName[key]; ok { // Если такое значение существует(т.е. существует колонка)
				setCell(book, wotkSheet, indexItem, colName[key], val)
			} else {
				colName[key] = startIndexCollumn
				setHead(book, wotkSheet, colName[key], key)
				setCell(book, wotkSheet, indexItem, colName[key], val)
				startIndexCollumn++
			}
		}
		/*
			// Обработка мапы доп полей
			for key, val := range valItem.Specifications {
				if _, ok := colName[key]; ok { // Если такое значение существует(т.е. существует колонка)
					setCell(book, wotkSheet, indexItem, colName[key], val)
				} else {
					colName[key] = startIndexCollumn
					setHead(book, wotkSheet, colName[key], key)
					setCell(book, wotkSheet, indexItem, colName[key], val)
					startIndexCollumn++
				}
			}
		*/
	}

	// Закрыть книгу
	closeBookError := CloseXlsx(book, filename)
	if closeBookError != nil {
		return closeBookError
	}
	return nil
}

/*
// Добавить ссылку в массив строк
func addURL_toLink(links []string) []string {
	for index := range links {
		links[index] = URL + links[index]
	}
	return links
}
*/

// Вписать значение в ячейку
func setCell(file *excelize.File, wotkSheet string, y, x int, value interface{}) {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	file.SetCellValue(wotkSheet, collumnStr+strconv.Itoa(y), value)
}

// Вписать значение в название колонки
func setHead(file *excelize.File, wotkSheet string, x int, value interface{}) {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	file.SetCellValue(wotkSheet, collumnStr+"1", value)
}

// *** for csv file ***
func incriment(value *int) int {
	*value++
	return *value
}

// Сохранить данные в Xlsx
func (variety Variety2) SaveXlsxCsvs(filename string) error {
	// Создать книгу
	book, makeBookError := MakeWorkBook()
	if makeBookError != nil {
		return makeBookError
	}

	//var position int
	wotkSheet := "main"
	setHead(book, wotkSheet, 1, "Номер")      // Catalog
	setHead(book, wotkSheet, 2, "Путь")       // Catalog > PodCatalog > Section
	setHead(book, wotkSheet, 3, "Каталог")    // Catalog
	setHead(book, wotkSheet, 4, "ПодКаталог") // PodCatalog
	setHead(book, wotkSheet, 5, "Секция")     // Section

	setHead(book, wotkSheet, 6, "Название товара")        // Name
	setHead(book, wotkSheet, 7, "Полное название товара") // FullName
	setHead(book, wotkSheet, 8, "Ссылка на товар")        // Link
	setHead(book, wotkSheet, 9, "Артикул")                // Article+color[0]
	setHead(book, wotkSheet, 10, "Производитель")         // Manufacturer
	setHead(book, wotkSheet, 11, "Цена")                  // Price
	setHead(book, wotkSheet, 12, "Цвет")                  // Colors
	setHead(book, wotkSheet, 13, "Ссылки на картинки")    // Colors
	setHead(book, wotkSheet, 14, "Размеры")               // Size
	setHead(book, wotkSheet, 15, "Описание товара")       // Description rus
	setHead(book, wotkSheet, 16, "Описание товара eng")   // Description eng

	var row int = 2
	for indexItem, valItem := range variety.Product {

		setCell(book, wotkSheet, row, 1, indexItem+1) // Номер

		setCell(book, wotkSheet, row, 2, valItem.Cat[0].Name+" > "+valItem.Cat[1].Name+" > "+valItem.Cat[2].Name) // Путь

		setCell(book, wotkSheet, row, 3, valItem.Cat[0].Name) // Каталог
		setCell(book, wotkSheet, row, 4, valItem.Cat[1].Name) // ПодКаталог
		setCell(book, wotkSheet, row, 5, valItem.Cat[2].Name) // Секция

		setCell(book, wotkSheet, row, 6, valItem.Name)     // Название товара
		setCell(book, wotkSheet, row, 7, valItem.FullName) // Полное название товара

		setCell(book, wotkSheet, row, 8, valItem.Link)          // Ссылка на товар
		setCell(book, wotkSheet, row, 9, valItem.Article)       // Артикул
		setCell(book, wotkSheet, row, 10, valItem.Manufacturer) // Производитель

		setCell(book, wotkSheet, row, 14, strings.Join(valItem.Size, ",")) // Все возможные размеры
		setCell(book, wotkSheet, row, 15, valItem.Description.Rus)         // Описание товара Rus
		setCell(book, wotkSheet, row, 16, valItem.Description.Eng)         // Описание товара eng
		row++

		// Обработка мапы картинок
		for keyImage, valImage := range valItem.Item {
			setCell(book, wotkSheet, row, 1, indexItem+1)                                                             // Номер
			setCell(book, wotkSheet, row, 2, valItem.Cat[0].Name+" > "+valItem.Cat[1].Name+" > "+valItem.Cat[2].Name) // Путь
			setCell(book, wotkSheet, row, 3, valItem.Cat[0].Name)                                                     // Каталог
			setCell(book, wotkSheet, row, 4, valItem.Cat[1].Name)                                                     // ПодКаталог
			setCell(book, wotkSheet, row, 5, valItem.Cat[2].Name)                                                     // Секция
			setCell(book, wotkSheet, row, 6, valItem.Name)                                                            // Название товара
			setCell(book, wotkSheet, row, 7, valItem.FullName)                                                        // Полное название товара
			setCell(book, wotkSheet, row, 8, valItem.Item[keyImage].Link)                                             // Ссылка на товар
			setCell(book, wotkSheet, row, 9, valItem.Article+"-"+valImage.ColorEng)                                   // Артикул
			setCell(book, wotkSheet, row, 10, valItem.Manufacturer)                                                   // Производитель
			setCell(book, wotkSheet, row, 11, valItem.Item[keyImage].Price)                                           // Цена
			setCell(book, wotkSheet, row, 12, keyImage)                                                               // Цвет // Буду ориентироваться на мапу картинок
			setCell(book, wotkSheet, row, 13, strings.Join(valItem.Item[keyImage].Image, ","))                        // Картинка
			setCell(book, wotkSheet, row, 14, strings.Join(valItem.Item[keyImage].Size, ","))                         // Размеры

			// Обработка мапы доп полей
			//var SpecificationsString string
			//for key, val := range valItem.Specifications {
			//	SpecificationsString += key + " - " + val + "\n"
			//}
			//setCell(book, wotkSheet, row, 16, valItem.Description.rus+"\n"+SpecificationsString) // Описание товара Rus

			row++
		}

	}

	// Закрыть книгу
	closeBookError := CloseXlsx(book, filename)
	if closeBookError != nil {
		return closeBookError
	}
	return nil
}
