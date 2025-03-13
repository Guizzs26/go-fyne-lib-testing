package ui

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Guizzs26/go-fyne-lib-testing/internal/excel"
	"github.com/Guizzs26/go-fyne-lib-testing/internal/models"
)

// Estou tratando datas como string (ainda não sei como lidar com datas)
func validateAndTransformInputs(firstName, lastName, ageStr, birthday, randomFloatStr string) (*models.FormData, error) {
	if firstName == "" || lastName == "" || ageStr == "" || birthday == "" || randomFloatStr == "" {
		return nil, errors.New("all fields must be filled")
	}

	if !isValidName(firstName) || !isValidName(lastName) {
		return nil, errors.New("first name and last name can only contain letters, spaces, hyphens, or apostrophes")
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, errors.New("invalid age format, must be an integer")
	}

	if age < 0 || age > 120 {
		return nil, errors.New("age must be a positive value and less than or equal to 120")
	}

	randomFloat, err := strconv.ParseFloat(randomFloatStr, 64)
	if err != nil {
		return nil, errors.New("invalid decimal value format")
	}

	if !isValidRealDate(birthday) {
		return nil, errors.New("invalid birthday date, must be a real date in the format dd/mm/yyyy")
	}

	return &models.FormData{
		FirstName:   firstName,
		LastName:    lastName,
		Age:         age,
		Birthday:    birthday,
		RandomFloat: randomFloat,
	}, nil
}

// Valida a data no formato dd/mm/yyyy usando o regex
func isValidRealDate(dateStr string) bool {
	layout := "02/01/2006" // Formato dd/mm/yyyy da lib time do Golang (a doc explica o significa dos dígitos)
	_, err := time.Parse(layout, dateStr)
	return err == nil
}

// Permite apenas letras
func isValidName(input string) bool {
	return regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s'-]+$`).MatchString(input)
}

/*
Reponsável por:
- Criar e retornar um campo de entrada (Entry) do fyne e definir o placeholder do input
*/
func newInputField(placeholder string) *widget.Entry {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholder)
	return input
}

/*
Reponsável por:
- Receber os dados validados e chamar a função generate excel para gerar a planilha
*/
func onGenerateExcel(firstName, lastName, age, birthday, randomFloat *widget.Entry) {
	formData, err := validateAndTransformInputs(firstName.Text, lastName.Text, age.Text, birthday.Text, randomFloat.Text)
	if err != nil {
		fmt.Println("Validation error", err)
		return
	}

	excelErr := excel.GenerateExcel(formData, "Arquivo.xlsx")
	if excelErr != nil {
		fmt.Println("Error generating excel", excelErr)
	}
}

/*
Responsável por:

- Criar os campos de input (com o widget do Fyne)
- Adicionar o botão "Gerar excel" que chama a função responsável por gerar o Excel
*/
func createInputFields() (*widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry, *widget.Button) {
	firstNameInput := newInputField("Enter your first name")
	lastNameInput := newInputField("Enter your last name")
	ageInput := newInputField("Enter your age")
	birthdayInput := newInputField("Enter your birthday date (dd/mm/yyyy)")
	randomFloatInput := newInputField("Enter any decimal value")

	generateExcelButton := widget.NewButton("Generate Excel Spreadsheet", func() {
		onGenerateExcel(firstNameInput, lastNameInput, ageInput, birthdayInput, randomFloatInput)
	})

	return firstNameInput, lastNameInput, ageInput, birthdayInput, randomFloatInput, generateExcelButton
}

/*
Responsável por:

  - Organizar esses componentes o container do Fyne
*/
func createLayout(firstName, lastName, age, birthday, randomFloat *widget.Entry, button *widget.Button) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Fill in the data to generate the Excel"),
		firstName,
		lastName,
		age,
		birthday,
		randomFloat,
		button,
	)
}

/*
Responsável por:

  - Pela manipulação da janela principal (criação, definição da janela como principal, centralizar, modificar tamanho)
  - Setar o conteúdo
  - Exibir a janela e inicar o loop do Fyne
*/
func CreateWindow() {
	app := app.New()
	mainWindow := app.NewWindow("Excel generator")
	mainWindow.SetMaster()
	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(400, 200))

	firstNameInput, lastNameInput, ageInput, birthdayInput, randomFloatInput, generateExcelButton := createInputFields()
	content := createLayout(firstNameInput, lastNameInput, ageInput, birthdayInput, randomFloatInput, generateExcelButton)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
