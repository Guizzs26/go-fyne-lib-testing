package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func GenerateExcel(firstName, lastName string, age int, birthday string, randomFloat float64) {
	file := excelize.NewFile()
	sheetName := "Planilha Teste"
	file.SetSheetName("Sheet1", sheetName)

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	headers := []interface{}{"First Name", "Last Name", "Age", "Birthday", "Random Float"}
	for i, header := range headers {
		cell := fmt.Sprintf("%s%d", string(rune(65+i)), 1)
		file.SetCellValue(sheetName, cell, header)
	}

	data := [][]interface{}{{firstName, lastName, age, birthday, randomFloat}}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	if err := file.SaveAs("test.xlsx"); err != nil {
		fmt.Println("Error :", err)
		return
	}
}
