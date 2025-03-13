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

func GenerateExcel(data *models.FormData, filename string) error {
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
		return fmt.Errorf("error setting sheet name: %w", err)
	}

	if err := writeHeaders(file, defaultSheetName); err != nil {
		return fmt.Errorf("error writing headers: %w", err)
	}

	if err := writeData(file, defaultSheetName, data); err != nil {
		return fmt.Errorf("error writing data: %w", err)
	}

	if err := file.SaveAs(filename); err != nil {
		return fmt.Errorf("error saving the spreadsheet: %w", err)
	}

	return nil
}

func writeHeaders(file *excelize.File, sheet string) error {
	headers := []interface{}{"First Name", "Last Name", "Age", "Birthday", "Random Float"}

	for i, header := range headers {
		cell, err := excelize.JoinCellName(string(rune(65+i)), 1)

		if err != nil {
			return fmt.Errorf("error generating cell name: %w", err)
		}

		if err := file.SetCellValue(sheet, cell, header); err != nil {
			return err
		}
	}

	return nil
}

func writeData(file *excelize.File, sheet string, data *models.FormData) error {
	row := []interface{}{data.FirstName, data.LastName, data.Age, data.Birthday, data.RandomFloat}

	for i, value := range row {
		cell, err := excelize.JoinCellName(string(rune(65+i)), 2)

		if err != nil {
			return fmt.Errorf("error generating cell name: %w", err)
		}

		if err := file.SetCellValue(sheet, cell, value); err != nil {
			return err
		}
	}

	return nil
}
