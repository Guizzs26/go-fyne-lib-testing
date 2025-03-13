package excel

import (
	"fmt"
	"os"

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

	var file *excelize.File

	// Verifica se o arquivo já existe através da lib os (operational system)
	if _, err := os.Stat(filename); err == nil {
		file, err = excelize.OpenFile(filename)

		if err != nil {
			return fmt.Errorf("error opening existing file: %w", err)
		}
	} else if os.IsNotExist(err) {
		file = excelize.NewFile()

		if err := file.SetSheetName("Sheet1", defaultSheetName); err != nil {
			return fmt.Errorf("error setting sheet name: %w", err)
		}

		if err := writeHeaders(file, defaultSheetName); err != nil {
			return fmt.Errorf("error writing headers: %w", err)
		}
	} else {
		return fmt.Errorf("error checking file existence: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	if err := file.SetSheetName("Sheet1", defaultSheetName); err != nil {
		return fmt.Errorf("error setting sheet name: %w", err)
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
		cell, err := excelize.CoordinatesToCellName(i+1, 1) // Basicamente é coluna i+1, linha 1

		if err != nil {
			return fmt.Errorf("error generating cell name: %w", err)
		}

		if err := file.SetCellValue(sheet, cell, header); err != nil {
			return fmt.Errorf("error setting cell value: %w", err)
		}
	}

	return nil
}

func writeData(file *excelize.File, sheet string, data *models.FormData) error {
	rows, err := file.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("error getting rows: %w", err)
	}

	lastRow := len(rows) + 1 // É a próxima linha a partir da última linha preenchida

	row := []interface{}{data.FirstName, data.LastName, data.Age, data.Birthday, data.RandomFloat}

	for i, value := range row {
		cell, err := excelize.CoordinatesToCellName(i+1, lastRow)

		if err != nil {
			return fmt.Errorf("error generating cell name: %w", err)
		}

		if err := file.SetCellValue(sheet, cell, value); err != nil {
			return fmt.Errorf("error setting cell value: %w", err)
		}
	}

	return nil
}
