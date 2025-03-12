package excel

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func GenerateExcel() {
	file := excelize.NewFile()

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	headers := []string{"First Name", "Last Name", "Age", "Birthday", "Decimal"}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	data := [][]interface{}{
		{"Guilherme", "Teixeira", 21, "26/05/2003", 67.5},
		{"Cristian", "Santos", 28, "17/03/2005", 0.12381},
		{"Izabel", "Araújo", 18, "19/05/2003", 44.5},
	}

	for i, row := range data {
		dataRow := i + 2
		for j, col := range row {
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+j)), dataRow), col)
		}
	}

	if err := file.SaveAs("test.xlsx"); err != nil {
		log.Fatal(err)
	}
}
