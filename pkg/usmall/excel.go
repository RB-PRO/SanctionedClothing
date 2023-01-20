package usmall

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Создать книгу
func makeWorkBook() (*excelize.File, error) {
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
func closeXlsx(f *excelize.File, filename string) error {
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
func (variety Variety) SaveXlsx(filename string) error {
	// Создать книгу
	book, makeBookError := makeWorkBook()
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
	setHead(book, wotkSheet, 11, "Описание товара")       // Description
	setHead(book, wotkSheet, 12, "Ссылки на картинки")    // ImageLink
	setHead(book, wotkSheet, 13, "Цвета")                 // Colors
	setHead(book, wotkSheet, 14, "Размеры")               // Size
	startIndexCollumn := 15

	// Создаём мапу, которая будет содержать значения номеров колонок
	colName := make(map[string]int)
	for indexItem, valItem := range variety.Product {
		setCell(book, wotkSheet, indexItem, 1, valItem.Catalog)                   // Каталог
		setCell(book, wotkSheet, indexItem, 2, valItem.PodCatalog)                // ПодКаталог
		setCell(book, wotkSheet, indexItem, 3, valItem.Section)                   // Секция
		setCell(book, wotkSheet, indexItem, 4, valItem.PodSection)                // Подсекция
		setCell(book, wotkSheet, indexItem, 5, valItem.Name)                      // Название товара
		setCell(book, wotkSheet, indexItem, 6, valItem.FullName)                  // Полное название товара
		setCell(book, wotkSheet, indexItem, 7, URL+valItem.Link)                  // Ссылка на товар
		setCell(book, wotkSheet, indexItem, 8, valItem.Article)                   // Артикул
		setCell(book, wotkSheet, indexItem, 9, valItem.Manufacturer)              // Производитель
		setCell(book, wotkSheet, indexItem, 10, valItem.Price)                    // Цена
		setCell(book, wotkSheet, indexItem, 11, valItem.Description.rus)          // Описание товара
		setCell(book, wotkSheet, indexItem, 12, addURL_toLink(valItem.ImageLink)) // Ссылки на картинки
		setCell(book, wotkSheet, indexItem, 13, valItem.Colors)                   // Цвета
		setCell(book, wotkSheet, indexItem, 14, valItem.Size)                     // Размеры

		for key, val := range valItem.Specifications {
			if _, ok := colName[key]; ok { // Если такое значение существует(т.е. существует колонка)
				//do something here
				setCell(book, wotkSheet, indexItem, colName[key], val)
			} else {
				colName[key] = startIndexCollumn
				setHead(book, wotkSheet, colName[key], key)
				setCell(book, wotkSheet, indexItem, colName[key], val)
				startIndexCollumn++
			}

		}
	}

	// Закрыть книгу
	closeBookError := closeXlsx(book, filename)
	if closeBookError != nil {
		return closeBookError
	}
	return nil
}

// Добавить ссылку в массив строк
func addURL_toLink(links []string) []string {
	for index := range links {
		links[index] = URL + links[index]
	}
	return links
}

// Вписать значение в ячейку
func setCell(file *excelize.File, wotkSheet string, y, x int, value interface{}) {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	file.SetCellValue(wotkSheet, collumnStr+strconv.Itoa(y+2), value)
}

// Вписать значение в название колонки
func setHead(file *excelize.File, wotkSheet string, x int, value interface{}) {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	file.SetCellValue(wotkSheet, collumnStr+"1", value)
}
