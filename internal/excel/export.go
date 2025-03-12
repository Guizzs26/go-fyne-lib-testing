package excel

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func GenerateExcel(firstName, lastName, age, birthday, randomFloat string) {
	file := excelize.NewFile()

	sheetName := "Planilha Teste"
	file.SetSheetName("Sheet1", sheetName)

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	headers := []interface{}{firstName, lastName, birthday, randomFloat}
	for i, header := range headers {
		// Rune converte o valor numérico para um caractece (rune é um tipo em Go que representa um caractere unicode)
		// string(rune()) -> Converte o caractere para string
		// O 1 fixo é porque a coluna fica na primeira linha da planilha apenas
		cell := fmt.Sprintf("%s%d", string(rune(65+i)), 1)
		file.SetCellValue("Sheet1", cell, header)
	}

	// Preciso implementar uma forma de validar e transformar os dados antes de chamar a função para gerar o excel
	data := [][]interface{}{firstName, lastName, age, birthday, randomFloat}

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
