package excel

import (
	"fmt"

	"github.com/Guizzs26/go-fyne-lib-testing/internal/models"
	"github.com/xuri/excelize/v2"
)

const (
	defaultSheetName = "Planilha Teste"
	defaultFileName  = "test.xlsx"
)

func GenerateExcel(data *models.FormData, filename string) {
	if filename == "" {
		filename = defaultFileName
	}

	file := excelize.NewFile()

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	if err := file.SetSheetName("Sheet1", defaultSheetName); err != nil {
		fmt.Println("Error setting sheet name:", err)
		return
	}

	if err := writeHeaders(file, defaultSheetName); err != nil {
		fmt.Println("Error writing headers:", err)
		return
	}

	if err := writeData(file, defaultSheetName, data); err != nil {
		fmt.Println("Error writing data:", err)
		return
	}

	if err := file.SaveAs("test.xlsx"); err != nil {
		fmt.Println("Error saving the spreadsheet :", err)
		return
	}
}

func writeHeaders(file *excelize.File, sheet string) error {
	headers := []interface{}{"First Name", "Last Name", "Age", "Birthday", "Random Float"}

	for i, header := range headers {
		cell, _ := excelize.JoinCellName(string(rune(65+i)), 1)

		if err := file.SetCellValue(sheet, cell, header); err != nil {
			return err
		}
	}

	return nil
}

func writeData(file *excelize.File, sheet string, data *models.FormData) error {
	row := []interface{}{data.FirstName, data.LastName, data.Age, data.Birthday, data.RandomFloat}

	for i, value := range row {
		cell, _ := excelize.JoinCellName(string(rune(65+i)), 2)

		if err := file.SetCellValue(sheet, cell, value); err != nil {
			return err
		}
	}

	return nil
}
