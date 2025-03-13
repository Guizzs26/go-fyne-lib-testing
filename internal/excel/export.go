package excel

import (
	"fmt"

	"github.com/Guizzs26/go-fyne-lib-testing/internal/models"
	"github.com/xuri/excelize/v2"
)

func GenerateExcel(data models.FormData) {
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

	dataRow := [][]interface{}{{data.FirstName, data.LastName, data.Age, data.Birthday, data.RandomFloat}}

	for i, row := range dataRow {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue(sheetName, fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	if err := file.SaveAs("test.xlsx"); err != nil {
		fmt.Println("Error saving the spreadsheet :", err)
		return
	}
}
